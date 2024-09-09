package repo

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
)

type LessonRepo struct {
	db *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{db: db}
}

func (r *LessonRepo) CreateLesson(req *pb.LessonCreateReq) error {
	query := `INSERT INTO lessons (
								id, 
								course_id, 
								name, 
								description, 
								video_url, 
								created_at,
								updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query,
		uuid.NewString(),
		req.CurseId,
		req.Name,
		req.Description,
		req.VideoUrl,
		time.Now(),
		time.Now(),
	)
	return err
}

func (r *LessonRepo) GetLesson(byId *pb.ById) (*pb.LessonGet, error) {
	query := `SELECT 
					id, 
					course_id, 
					name, 
					description, 
					video_url, 
					created_at 
	          FROM lessons 
			  WHERE id = $1`
	row := r.db.QueryRow(query, byId.Id)

	var lesson pb.LessonGet
	err := row.Scan(
		&lesson.Id,
		&lesson.CurseId,
		&lesson.Name,
		&lesson.Description,
		&lesson.VideoUrl,
		&lesson.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (r *LessonRepo) UpdateLesson(req *pb.LessonUpdate) error {
	query := `UPDATE lessons set updated_at = $1 `
	var cons []string
	var args []interface{}
	args = append(args, time.Now())
	if req.Body.Name != "" && req.Body.Name != "string" {
		cons = append(cons, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, req.Body.Name)
	}
	if req.Body.Description != "" && req.Body.Description != "string" {
		cons = append(cons, fmt.Sprintf("description = $%d", len(args)+1))
		args = append(args, req.Body.Description)
	}
	if req.Body.VideoUrl != "" && req.Body.VideoUrl != "string" {
		cons = append(cons, fmt.Sprintf("video_url = $%d", len(args)+1))
		args = append(args, req.Body.VideoUrl)
	}
	if len(cons) == 0 {
		return fmt.Errorf("at least one field should be updated")
	}
	query += ", " + strings.Join(cons, ", ")
	query += fmt.Sprintf( " WHERE id = $%d", len(args)+1)
	args = append(args, req.Id)
	_, err := r.db.Exec(query, args...)
	return err

}

func (r *LessonRepo) DeleteLesson(byId *pb.ById) error {
	query := `update lessons set deleted_at = $1 WHERE id = $2 and deleted_at = 0`
	_, err := r.db.Exec(query, time.Now().Unix(), byId.Id)
	return err
}

func (r *LessonRepo) ListLessons(req *pb.LessonFilter) (*pb.LessonList, error) {
	query := `SELECT id, 
					course_id, 
					name, 
					description, 
					video_url, 
					created_at 
	          FROM lessons 
			  WHERE deleted_at = 0`
	var cons []string
	var args []interface{}
	if req.CourseId != "" && req.CourseId != "string" {
		cons = append(cons, fmt.Sprintf("course_id = $%d", len(args)+1))
		args = append(args, req.CourseId)
	}
	if len(cons) > 0 {
		query += " AND " + strings.Join(cons, " AND ")
	}
	query += " ORDER BY created_at DESC"
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
	var lessons pb.LessonList
	for rows.Next() {
		var lesson pb.LessonGet
		err := rows.Scan(
			&lesson.Id,
			&lesson.CurseId,
			&lesson.Name,
			&lesson.Description,
			&lesson.VideoUrl,
			&lesson.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		lessons.Lessons = append(lessons.Lessons, &lesson)
	}
	return &lessons, nil
}
 