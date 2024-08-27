package repo

import (
	"database/sql"

	"gitlab.com/muallimah/course_service/internal/storage"
)

type Storage struct {
	CategoryS    storage.CategoryI
	CertificateS storage.CertificateI
	CourseS      storage.CourseI
	LessonS      storage.LessonI
	MaterialS    storage.MaterialI
	TransactionS storage.TransactionI
	UserCourseS  storage.UserCourseI
	UserLessonS  storage.UserLessonI
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
        CategoryS:    NewCategoryRepo(db),
        CertificateS: NewCertificateRepo(db),
        CourseS:      NewCourseRepo(db),
        LessonS:      NewLessonRepo(db),
        MaterialS:    NewMaterialRepo(db),
        TransactionS: NewTransactionRepo(db),
        UserCourseS:  NewUserCourseRepo(db),
        UserLessonS:  NewUserLessonRepo(db),
    }
}

func (s *Storage) Category() storage.CategoryI {
    return s.CategoryS
}

func (s *Storage) Certificate() storage.CertificateI {
    return s.CertificateS
}

func (s *Storage) Course() storage.CourseI {
    return s.CourseS
}

func (s *Storage) Lesson() storage.LessonI {
    return s.LessonS
}

func (s *Storage) Material() storage.MaterialI {
    return s.MaterialS
}

func (s *Storage) Transaction() storage.TransactionI {
    return s.TransactionS
}

func (s *Storage) UserCourse() storage.UserCourseI {
    return s.UserCourseS
}

func (s *Storage) UserLesson() storage.UserLessonI {
    return s.UserLessonS
}