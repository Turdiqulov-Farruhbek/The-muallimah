package storage

import pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"

type StorageI interface {
	Category() CategoryI
	Certificate() CertificateI
	Course() CourseI
	Lesson() LessonI
	Material() MaterialI
	Transaction() TransactionI
	UserCourse() UserCourseI
	UserLesson() UserLessonI
}
type CategoryI interface {
	CreateCategory(req *pb.CategoryCreateReq) (*pb.Void,error) 
	GetCategory(id *pb.ById) (*pb.Category,error) 
	ListCategories(req *pb.Pagination) (*pb.CategoryListRes, error)
	UpdateCategory(req *pb.CategoryUpdateReq) (*pb.Void,error) 
	DeleteCategory(byId *pb.ById) (*pb.Void,error) 
}
type CertificateI interface {
	CreateCertificate(id *pb.ById) (*pb.Void,error)
	GetCertificate(id *pb.ById) (*pb.CertificateGet,error) 
	UpdateCertificate(req *pb.CertificateUpdate) (*pb.Void, error) 
	DeleteCertificate(byId *pb.ById) (*pb.Void, error) 
	ListCertificates(req *pb.CertificateFilter) (*pb.CertificateList, error) 
}
type CourseI interface {
	CreateCourse(req *pb.CourseCreateReq) error
	GetCourse(byId *pb.ById) (*pb.Course, error)
	UpdateCourse(req *pb.CourseUpdateReq) error 
	DeleteCourse(byId *pb.ById) error 
	ListCourses(req *pb.CourseListsReq) (*pb.CourseListsRes, error) 
}
type LessonI interface {
	CreateLesson(req *pb.LessonCreateReq) error
	GetLesson(byId *pb.ById) (*pb.LessonGet, error)
	UpdateLesson(req *pb.LessonUpdate) error
	DeleteLesson(byId *pb.ById) error 
	ListLessons(req *pb.LessonFilter) (*pb.LessonList, error) 
}
type MaterialI interface {
	CreateMaterial(req *pb.MaterialCreateReq) (*pb.Void,error) 
	GetMaterial(byId *pb.ById) (*pb.MaterialGetRes, error)
	UpdateMaterial(req *pb.MaterialUpdateReq) (*pb.Void,error) 
	DeleteMaterial(byId *pb.ById) (*pb.Void, error)
	ListMaterials(req *pb.MaterialListReq) (*pb.MaterialListRes, error) 
}
type TransactionI interface {
	CreateTransaction(req *pb.TransactionCreateReq) (*pb.Void,error) 
	GetTransaction(id *pb.ById) (*pb.TransactionGetRes,error)
	ListTransactions(req *pb.TransactionListReq) (*pb.TransactionListRes, error)
}
type UserCourseI interface {
	EnrollUserInCourse(req *pb.UserCourseCreateReq) error
	GetUserCourse(byId *pb.ById) (*pb.UserCourse, error) 
	UpdateUserCourse(req *pb.UserCourseUpdateReq) error 
	DeleteUserCourse(byId *pb.ById) error
	ListUserCourses(req *pb.UserCourseListsReq) (*pb.UserCourseListsRes, error)
}
type UserLessonI interface {
	CreateUserLesson(req *pb.UserLessonCreateReq) error 
	GetUserLesson(byId *pb.ById) (*pb.UserLesson, error)
	UpdateUserLesson(req *pb.UserLessonUpdateReq) error 
	DeleteUserLesson(byId *pb.ById) error 
	ListUserLessons(req *pb.UserLessonListsReq) (*pb.UserLessonListsRes, error) 
}
