package repo

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
)

type CourseRepo struct {
	db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db: db}
}

func (r *CourseRepo) CreateCourse(req *pb.CourseCreateReq) error {
	query := `INSERT INTO courses (
		id,	
		name, 
		description, 
		price, 
		image_url, 
		category_id, 
		created_at, 
		updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(query,
		uuid.NewString(),
		req.Name,
		req.Description,
		req.Price,
		req.ImageUrl,
		req.CategoryId,
		time.Now(),
		time.Now(),
	)
	return err
}

func (r *CourseRepo) GetCourse(byId *pb.ById) (*pb.Course, error) {
	query := `SELECT c.id, 
					c.name, 
					c.description, 
					c.price, 
					c.image_url, 
					ct.id,
					ct.name, 
					c.created_at, 
					c.updated_at 
	          FROM courses c
			  LEFT JOIN categories ct ON ct.id = c.category_id
			  WHERE c.id = $1`
	row := r.db.QueryRow(query, byId.Id)

	var course pb.Course
	course.Category = &pb.Category{}
	err := row.Scan(
		&course.Id,
		&course.Name,
		&course.Description,
		&course.Price,
		&course.ImageUrl,
		&course.Category.Id,
		&course.Category.Name,
		&course.CreatedAt,
		&course.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *CourseRepo) UpdateCourse(req *pb.CourseUpdateReq) error {
	query := `UPDATE courses SET updated_at = NOW() `
	var cons []string
	var args []interface{}

	if req.Body.Name != "" && req.Body.Name != "string" {
		cons = append(cons, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, req.Body.Name)
	}
	if req.Body.Description != "" && req.Body.Description != "string" {
		cons = append(cons, fmt.Sprintf("description = $%d", len(args)+1))
		args = append(args, req.Body.Description)
	}
	if req.Body.Price != 0 {
		cons = append(cons, fmt.Sprintf("price = $%d", len(args)+1))
		args = append(args, req.Body.Price)
	}
	if req.Body.ImageUrl != "" && req.Body.ImageUrl != "string" {
		cons = append(cons, fmt.Sprintf("image_url = $%d", len(args)+1))
		args = append(args, req.Body.ImageUrl)
	}
	if len(cons) == 0 {
		return fmt.Errorf("at least one field should be updated")
	}
	query += ", " + strings.Join(cons, ", ")
	query += " WHERE id = $%d"
	args = append(args, req.Id)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *CourseRepo) DeleteCourse(byId *pb.ById) error {
	query := `update courses set deleted_at = $1 WHERE id = $2 and deleted_at = 0`
	_, err := r.db.Exec(query, time.Now().Unix(), byId.Id)
	return err
}

func (r *CourseRepo) ListCourses(req *pb.CourseListsReq) (*pb.CourseListsRes, error) {
	query := `SELECT c.id, 
					c.name, 
					c.description, 
					c.price, 
					c.image_url, 
					ct.id,
					ct.name, 
					c.created_at, 
					c.updated_at 
	          FROM courses c
			  INNER JOIN categories ct ON ct.id = c.category_id
			  WHERE c.deleted_at = 0`
	var cons []string
	var args []interface{}
	if req.CategoryId != "" && req.CategoryId != "string" {
		cons = append(cons, fmt.Sprintf("c.category_id = $%d", len(args)+1))
		args = append(args, req.CategoryId)
	}
	if len(cons) > 0 {
		query += " AND " + strings.Join(cons, " AND ")
	}
	query += " ORDER BY c.created_at DESC"
	if req.Filter.Limit != 0 {
		query += fmt.Sprintf(" LIMIT $%d", len(args)+1)
		args = append(args, req.Filter.Limit)
	}
	if req.Filter.Offset != 0 {
		query += fmt.Sprintf(" OFFSET $%d", len(args)+1)
		args = append(args, req.Filter.Offset)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*pb.Course
	for rows.Next() {
		var course pb.Course
		course.Category = &pb.Category{}
		err := rows.Scan(
			&course.Id,
			&course.Name,
			&course.Description,
			&course.Price,
			&course.ImageUrl,
			&course.Category.Id,
			&course.Category.Name,
			&course.CreatedAt,
			&course.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}

	totalCount := len(courses)

	return &pb.CourseListsRes{
		Courses:    courses,
		Pagination: req.Filter,
		TotalCount: int32(totalCount),
	}, nil
}
