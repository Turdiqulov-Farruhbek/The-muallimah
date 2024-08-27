package repo

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"

)

type UserCourseRepo struct {
	db *sql.DB
}

func NewUserCourseRepo(db *sql.DB) *UserCourseRepo {
	return &UserCourseRepo{db: db}
}

func (r *UserCourseRepo) EnrollUserInCourse(req *pb.UserCourseCreateReq) error {
	query := `INSERT INTO user_courses (
								id, 
								user_id, 
								course_id, 
								start_date, 
								end_date, 
								status)
	          VALUES ($1, $2, $3, $4, $5, $6)`
	y,m,d := time.Now().Date()
	startDate := fmt.Sprintf("%d-%02d-%02d", y, m, d)		  
	_, err := r.db.Exec(query, 
		uuid.NewString(), 
		req.UserId, 
		req.CourseId, 
		startDate, 
		startDate, 
		"started",
	)
	return err
}

func (r *UserCourseRepo) GetUserCourse(byId *pb.ById) (*pb.UserCourse, error) {
	query := `
	SELECT 
		uc.id, 
		uc.start_date, 
		uc.end_date, 
		uc.status, 
		u.id, 
		u.first_name, 
		u.last_name, 
		u.dob, 
		u.phone_number, 
		u.email, 
		u.occupation, 
		u.address, 
		u.gender, 
		u.photo_url, 
		c.id, 
		c.name, 
		c.description, 
		c.price, 
		c.image_url, 
		ct.id,
		ct.name,
		c.created_at, 
		c.updated_at 
	FROM 
		user_courses uc
	INNER JOIN 
		users u ON uc.user_id = u.id
	INNER JOIN 
		courses c ON uc.course_id = c.id
	INNER JOIN 
		categories ct on ct.id = c.category_id	
	WHERE 
		uc.id = $1 AND uc.deleted_at = 0`

	row := r.db.QueryRow(query, byId.Id)

	var userCourse pb.UserCourse
	var user pb.UserGetRes
	var course pb.Course
	err := row.Scan(
		&userCourse.Id, 
		&userCourse.StartDate, 
		&userCourse.EndDate, 
		&userCourse.Status,
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Dob,
		&user.PhoneNumber,
		&user.Email,
		&user.Occupation,
		&user.Address,
		&user.Gender,
		&user.PhotoUrl,
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
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user course with ID %s not found", byId.Id)
		}
		return nil, err
	}
	userCourse.User = &user
	userCourse.Course = &course

	return &userCourse, nil
}

func (r *UserCourseRepo) UpdateUserCourse(req *pb.UserCourseUpdateReq) error {
	query := `UPDATE user_courses SET `
	var cons []string
	var args []interface{}
	if req.Body.StartDate!= "" && req.Body.EndDate!= "string" {
        cons = append(cons, fmt.Sprintf("start_date = $%d", len(args)+1))
        args = append(args, req.Body.StartDate)
    }
	if req.Body.EndDate!= "" && req.Body.EndDate!= "string" {
        cons = append(cons, fmt.Sprintf("end_date = $%d", len(args)+1))
        args = append(args, req.Body.EndDate)
    }
	if req.Body.Status != "" && req.Body.Status != "string" {
		cons = append(cons, fmt.Sprintf("status = $%d", len(args)+1))
        args = append(args, req.Body.Status)
		if req.Body.Status == "completed" {
			y,m,d := time.Now().Date()
			endDate := fmt.Sprintf("%d-%02d-%02d", y, m, d)
			cons = append(cons, fmt.Sprintf("end_date = $%d", len(args)+1))
			args = append(args, endDate)
		}
	}
	if len(cons) == 0 {
        return fmt.Errorf("at least one field should be updated")
    }
	query += strings.Join(cons, ", ")
	query += " WHERE id = $%d"
	args = append(args, req.Id)
	_, err := r.db.Exec(query, args...)
	return err

}
func (r *UserCourseRepo) DeleteUserCourse(byId *pb.ById) error {
	query := `update user_courses set deleted_at = $1 WHERE id = $2 and deleted_at = 0`
	_, err := r.db.Exec(query,time.Now().Unix(), byId.Id)
	return err
}

func (r *UserCourseRepo) ListUserCourses(req *pb.UserCourseListsReq) (*pb.UserCourseListsRes, error) {
	query := `
	SELECT 
		uc.id, 
		uc.start_date, 
		uc.end_date, 
		uc.status, 
		u.id, 
		u.first_name, 
		u.last_name, 
		u.dob, 
		u.phone_number, 
		u.email, 
		u.occupation, 
		u.address, 
		u.gender, 
		u.photo_url, 
		c.id, 
		c.name, 
		c.description, 
		c.price, 
		c.image_url, 
		ct.id,
		ct.name,
		c.created_at, 
		c.updated_at 
	FROM 
		user_courses uc
	INNER JOIN 
		users u ON uc.user_id = u.id
	INNER JOIN 
		courses c ON uc.course_id = c.id
	INNER JOIN 
		categories ct on ct.id = c.category_id	
	`

	var conditions []string
	var args []interface{}

	if req.UserId != "" {
		conditions = append(conditions, fmt.Sprintf("uc.user_id = $%d", len(args)+1))
		args = append(args, req.UserId)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	
	query += " AND uc.deleted_at = 0 ORDER BY uc.created_at DESC"

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

	var userCourses []*pb.UserCourse
	for rows.Next() {
		var userCourse pb.UserCourse
		err := rows.Scan(
			&userCourse.Id, 
			&userCourse.StartDate, 
			&userCourse.EndDate, 
			&userCourse.Status,
			&userCourse.User.Id,
			&userCourse.User.FirstName,
			&userCourse.User.LastName,
			&userCourse.User.Dob,
			&userCourse.User.PhoneNumber,
			&userCourse.User.Email,
			&userCourse.User.Occupation,
			&userCourse.User.Address,
			&userCourse.User.Gender,
			&userCourse.User.PhotoUrl,
			&userCourse.Course.Id,
			&userCourse.Course.Name,
			&userCourse.Course.Description,
			&userCourse.Course.Price,
			&userCourse.Course.ImageUrl,
			&userCourse.Course.Category.Id,
			&userCourse.Course.Category.Name,
			&userCourse.Course.CreatedAt,
			&userCourse.Course.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		userCourses = append(userCourses, &userCourse)
	}

	count := len(userCourses)

	return &pb.UserCourseListsRes{
		UserCourses: userCourses,
		Pagination:  req.Filter,
		TotalCount:  int32(count),
	}, nil
}
