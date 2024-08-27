package repo

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"

)

type CertificateRepo struct {
	db *sql.DB
}

func NewCertificateRepo(db *sql.DB) *CertificateRepo {
	return &CertificateRepo{db: db}
}
func (r *CertificateRepo) CreateCertificate(id *pb.ById) (*pb.Void,error) {
	query := `INSERT INTO certificates (
		id, 
		user_course_id, 
		status, 
		created_at) 
	VALUES ($1, $2, $3, $4)`
    time := time.Now()
    _, err := r.db.Exec(query, 
		uuid.NewString(), 
		id.Id, 
		"pending", 
        time)
    if err!= nil {
        return nil, err
    }
    return &pb.Void{}, nil
}
func (r *CertificateRepo) GetCertificate(id *pb.ById) (*pb.CertificateGet,error) {
	query := `SELECT cr.id, 
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
					cat.id, 
                    cat.name,
					c.created_at,
					c.updated_at, 
                    cr.status, 
                    cr.file_url
    	FROM certificates cr
		    LEFT JOIN user_course uc 
		        on uc.id = cr.user_course_id
			LEFT JOIN courses c 
				on c.id = uc.course_id
			LEFT JOIN users u 
			    on u.id = uc.user_id
			LEFT JOIN categories cat 
				on cat.id = c.category_id			
    WHERE id = $1`
    row := r.db.QueryRow(query, id.Id)

    var certificate pb.CertificateGet
    err := row.Scan(
		&certificate.Id,
		&certificate.UserCourse.Id,
		&certificate.UserCourse.StartDate,
		&certificate.UserCourse.EndDate,
		&certificate.UserCourse.Status,
		&certificate.UserCourse.User.Id,
		&certificate.UserCourse.User.FirstName,
		&certificate.UserCourse.User.LastName,
		&certificate.UserCourse.User.Dob,
		&certificate.UserCourse.User.PhoneNumber,
		&certificate.UserCourse.User.Email,
		&certificate.UserCourse.User.Occupation,
		&certificate.UserCourse.User.Address,
		&certificate.UserCourse.User.Gender,
		&certificate.UserCourse.User.PhotoUrl,
		&certificate.UserCourse.Course.Id,
		&certificate.UserCourse.Course.Name,
		&certificate.UserCourse.Course.Description,
		&certificate.UserCourse.Course.Price,
		&certificate.UserCourse.Course.ImageUrl,
		&certificate.UserCourse.Course.Category.Id,
		&certificate.UserCourse.Course.Category.Name,
		&certificate.UserCourse.Course.CreatedAt,
		&certificate.UserCourse.Course.UpdatedAt,
		&certificate.Status,
		&certificate.FileUrl,
	)
    if err!= nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("certificate with ID %s not found", id.Id)
        }
        return nil, err
    }
    return &certificate, nil
}
func (r *CertificateRepo) UpdateCertificate(req *pb.CertificateUpdate) (*pb.Void, error) {
	query := `UPDATE certificates SET `
    var cons []string
    var args []interface{}

    if req.Body.Status!= "" && req.Body.Status!= "string" {
        cons = append(cons, "status = $"+fmt.Sprint(len(cons)+1))
        args = append(args, req.Body.Status)
    }
    if req.Body.FileUrl!= "" && req.Body.FileUrl!= "string" {
        cons = append(cons, "file_url = $"+fmt.Sprint(len(cons)+1))
        args = append(args, req.Body.FileUrl)
    }

    query += strings.Join(cons, ", ")
    query += " WHERE id = $"+fmt.Sprint(len(cons)+1)
    args = append(args, req.Id)

    _, err := r.db.Exec(query, args...)
    return &pb.Void{},err
}

func (r *CertificateRepo) DeleteCertificate(byId *pb.ById) (*pb.Void, error) {
	query := `DELETE FROM certificates WHERE id = $1`
    _, err := r.db.Exec(query, byId.Id)
    return &pb.Void{}, err
}
func (r *CertificateRepo) ListCertificates(req *pb.CertificateFilter) (*pb.CertificateList, error) {
	query := `SELECT cr.id, 
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
					cat.id, 
                    cat.name,
					c.created_at,
					c.updated_at, 
                    cr.status, 
                    cr.file_url
    	FROM certificates cr
		    LEFT JOIN user_course uc 
		        on uc.id = cr.user_course_id
			LEFT JOIN courses c 
				on c.id = uc.course_id
			LEFT JOIN users u 
			    on u.id = uc.user_id
			LEFT JOIN categories cat 
				on cat.id = c.category_id			
            `
	var cons []string
	var args []interface{}
	if req.UserId!= "" && req.UserId!= "string"{
        cons = append(cons, fmt.Sprintf("u.id = $%d",len(args)+1))
        args = append(args, req.UserId)
    }
	if req.CourseId!= "" && req.CourseId!= "string"{
        cons = append(cons, fmt.Sprintf("c.id = $%d",len(args)+1))
        args = append(args, req.CourseId)
    }
	if req.Status != "" && req.Status != "string" {
		cons = append(cons, fmt.Sprintf("cr.status = $%d",len(args)+1))
        args = append(args, req.Status)
	}
	if len(cons) > 0 {
        query += " WHERE " + strings.Join(cons, " AND ")
    }
	query += " ORDER BY cr.created_at DESC"
	if req.Filter.Limit!= 0 {
        query += fmt.Sprintf(" LIMIT $%d", len(args)+1)
        args = append(args, req.Filter.Limit)
    }
	if req.Filter.Offset!= 0 {
        query += fmt.Sprintf(" OFFSET $%d", len(args)+1)
        args = append(args, req.Filter.Offset)
    }
	rows, err := r.db.Query(query, args...)
	if err!= nil {
        return nil, err
    }
	defer rows.Close()
	var certificates pb.CertificateList
	for rows.Next() {
		var certificate pb.CertificateGet
		err := rows.Scan(
			&certificate.Id,
			&certificate.UserCourse.Id,
			&certificate.UserCourse.StartDate,
			&certificate.UserCourse.EndDate,
			&certificate.UserCourse.Status,
			&certificate.UserCourse.User.Id,
			&certificate.UserCourse.User.FirstName,
			&certificate.UserCourse.User.LastName,
			&certificate.UserCourse.User.Dob,
			&certificate.UserCourse.User.PhoneNumber,
			&certificate.UserCourse.User.Email,
			&certificate.UserCourse.User.Occupation,
			&certificate.UserCourse.User.Address,
			&certificate.UserCourse.User.Gender,
			&certificate.UserCourse.User.PhotoUrl,
			&certificate.UserCourse.Course.Id,
			&certificate.UserCourse.Course.Name,
			&certificate.UserCourse.Course.Description,
			&certificate.UserCourse.Course.Price,
			&certificate.UserCourse.Course.ImageUrl,
			&certificate.UserCourse.Course.Category.Id,
			&certificate.UserCourse.Course.Category.Name,
			&certificate.UserCourse.Course.CreatedAt,
			&certificate.UserCourse.Course.UpdatedAt,
			&certificate.Status,
			&certificate.FileUrl,
		)
		if err!= nil {
            return nil, err
        }
		certificates.Items = append(certificates.Items, &certificate)
	}	
	return &certificates, nil
}
