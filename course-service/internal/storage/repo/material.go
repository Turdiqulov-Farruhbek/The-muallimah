package repo

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"

)

type MaterialRepo struct {
	db *sql.DB
}

func NewMaterialRepo(db *sql.DB) *MaterialRepo {
	return &MaterialRepo{db: db}
}


func (r *MaterialRepo) CreateMaterial(req *pb.MaterialCreateReq) (*pb.Void,error) {
	query := `INSERT INTO additional_materials (
                                id,    
                                lesson_id, 
                                title, 
                                description, 
                                resource_url)
              VALUES ($1, $2, $3, $4, $5)`
    _, err := r.db.Exec(query,
        uuid.NewString(),
        req.LessonId,
        req.Title,
        req.Description,
        req.ResourceUrl,
    )
    return &pb.Void{}, err
}

func (r *MaterialRepo) GetMaterial(byId *pb.ById) (*pb.MaterialGetRes, error) {
	query := `SELECT id, 
                    lesson_id, 
                    title, 
                    description, 
                    resource_url
              FROM additional_materials 
              WHERE id = $1`
    row := r.db.QueryRow(query, byId.Id)

    var material pb.MaterialGetRes
    err := row.Scan(
        &material.Id,
        &material.LessonId,
        &material.Title,
        &material.Description,
        &material.ResourceUrl,
    )
    if err!= nil {
        return nil, err
    }

    return &material, nil
}

func (r *MaterialRepo) UpdateMaterial(req *pb.MaterialUpdateReq) (*pb.Void,error) {
	query := `UPDATE additional_materials SET `
    var cons []string
    var args []interface{}

    if req.Body.Title!= "" && req.Body.Title!= "string" {
        cons = append(cons, fmt.Sprintf("title = $%d", len(args)+1))
        args = append(args, req.Body.Title)
    }
    if req.Body.Description!= "" && req.Body.Description!= "string" {
        cons = append(cons, fmt.Sprintf("description = $%d", len(args)+1))
        args = append(args, req.Body.Description)
    }
    if req.Body.ResourceUrl!= "" && req.Body.ResourceUrl!= "string" {
        cons = append(cons, fmt.Sprintf("resource_url = $%d", len(args)+1))
        args = append(args, req.Body.ResourceUrl)
	}
	if len(cons) == 0 {
        return nil, fmt.Errorf("at least one field should be updated")
    }
	query += strings.Join(cons, ", ")
	query += fmt.Sprintf(" WHERE id = $%d", len(args)+1)
	args = append(args, req.Id)
	_, err := r.db.Exec(query, args...)
	return &pb.Void{}, err
}

func (r *MaterialRepo) DeleteMaterial(byId *pb.ById) (*pb.Void, error) {
	query := `DELETE from additional_materials WHERE id = $1`
    _, err := r.db.Exec(query, byId.Id)
    return &pb.Void{}, err
}

func (r *MaterialRepo) ListMaterials(req *pb.MaterialListReq) (*pb.MaterialListRes, error) {
	query := `SELECT id, 
                    lesson_id, 
                    title, 
                    description, 
                    resource_url
              FROM additional_materials 
              WHERE lesson_id = $1`
    rows, err := r.db.Query(query, req.LessonId)
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    var materials pb.MaterialListRes
    for rows.Next() {
        var material pb.MaterialGetRes
        err := rows.Scan(
            &material.Id,
            &material.LessonId,
            &material.Title,
            &material.Description,
            &material.ResourceUrl,
        )
        if err!= nil {
            return nil, err
        }
        materials.Materials = append(materials.Materials, &material)
    }

    return &materials, nil
}