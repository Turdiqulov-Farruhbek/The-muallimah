package repo

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"

)

type UserLessonRepo struct {
	db *sql.DB
}

func NewUserLessonRepo(db *sql.DB) *UserLessonRepo {
    return &UserLessonRepo{db: db}
}

func (r *UserLessonRepo) CreateUserLesson(req *pb.UserLessonCreateReq) error {
	query := `INSERT INTO user_lessons (
								id, 
								lesson_id, 
								user_course_id, 
								status)
	          VALUES ($1, $2, $3, $4)`
		  
	_, err := r.db.Exec(query, 
		uuid.NewString(), 
		req.LessonId, 
		req.UserCourseId, 
		req.Status,
	)
	return err
}

func (r *UserLessonRepo) GetUserLesson(byId *pb.ById) (*pb.UserLesson, error) {
	query := `SELECT id, 
					lesson_id, 
					user_course_id, 
					status
	          FROM user_lessons 
			  WHERE id = $1`
	row := r.db.QueryRow(query, byId.Id)

	var userLesson pb.UserLesson
	err := row.Scan(
		&userLesson.Id, 
		&userLesson.LessonId, 
		&userLesson.UserCourse.Id, 
		&userLesson.Status,
	)
	if err != nil {
		return nil, err
	}

	return &userLesson, nil
}
func (r *UserLessonRepo) UpdateUserLesson(req *pb.UserLessonUpdateReq) error {
	query := `UPDATE user_lessons SET `
	var cons []string
	var args []interface{}

	if req.Body.Status != "" && req.Body.Status != "string" {
		cons = append(cons, fmt.Sprintf("status = $%d", len(args)+1))
        args = append(args, req.Body.Status)
	}

	if len(cons) == 0 {
        return fmt.Errorf("at least one field should be updated")
    }
	query += strings.Join(cons, ", ")
	query += fmt.Sprintf(" WHERE id = $%d", len(args)+1)
	args = append(args, req.Id)
	_, err := r.db.Exec(query, args...)
	return err
}
func (r *UserLessonRepo) DeleteUserLesson(byId *pb.ById) error {
	query := `UPDATE user_lessons SET deleted_at = $1 WHERE id = $2 AND deleted_at = 0`
	_, err := r.db.Exec(query, time.Now().Unix(), byId.Id)
	return err
}


func (r *UserLessonRepo) ListUserLessons(req *pb.UserLessonListsReq) (*pb.UserLessonListsRes, error) {
	query := `SELECT id, 
					lesson_id, 
					user_course_id, 
					status
	          FROM user_lessons`

	var cons []string
	var args []interface{}		  
	if req.UserCourseId != "" && req.UserCourseId != "string" {
		cons = append(cons, fmt.Sprintf("user_course_id = $%d", len(args)+1))
        args = append(args, req.UserCourseId)
    }

	if len(cons) > 0 {
        query += " WHERE " + strings.Join(cons, " AND ")
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

	var userLessons []*pb.UserLesson
	for rows.Next() {
		var userLesson pb.UserLesson
		err := rows.Scan(
			&userLesson.Id, 
			&userLesson.LessonId, 
			&userLesson.UserCourse.Id,
			&userLesson.Status,
		)
		if err != nil {
			return nil, err
		}
		userLessons = append(userLessons, &userLesson)
	}

	count := len(userLessons)

	return &pb.UserLessonListsRes{
		UserLessons: userLessons,
		Pagination:  req.Filter,
		TotalCount:  int32(count),
	}, nil
}
