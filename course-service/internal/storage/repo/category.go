package repo

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"

)

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *CategoryRepo {
    return &CategoryRepo{db: db}
}

func (r *CategoryRepo) CreateCategory(req *pb.CategoryCreateReq) (*pb.Void,error) {
	id := uuid.New().String()
    query := `INSERT INTO categories (id,  
                                    name) 
                VALUES ($1, $2)`
                
    _, err := r.db.Exec(query, 
            id, 
            req.Name)    

    if err!= nil {
        return nil, err
    }            
    return &pb.Void{}, nil
}
func (r *CategoryRepo) GetCategory(id *pb.ById) (*pb.Category,error) {
	query := `SELECT id, 
                    name
                    FROM categories
                    WHERE id = $1
    `
    var category pb.Category
    err := r.db.QueryRow(query, id.Id).Scan(
        &category.Id, 
        &category.Name,
    )
    if err!= nil {
        return nil, err
    }
    return &category, nil
}
func (r *CategoryRepo) ListCategories(req *pb.Pagination) (*pb.CategoryListRes, error) {
	query := `SELECT id, 
                    name
                    FROM categories
    `
	var args []interface{}
	if req.Limit != 0 {
		query += fmt.Sprintf("LIMIT %d ", req.Limit)
		args = append(args, req.Limit)
	}
	if req.Offset!= 0 {
        query += fmt.Sprintf("OFFSET %d ", req.Offset)
        args = append(args, req.Offset)
    }

    rows, err := r.db.Query(query, args...)
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    var categories []*pb.Category
    for rows.Next() {
        var category pb.Category
        err := rows.Scan(&category.Id, &category.Name)
        if err!= nil {
            return nil, err
        }
        categories = append(categories, &category)
    }

    totalCount := len(categories) 

    return &pb.CategoryListRes{
        Categories: categories,
        TotalCount: int32(totalCount),
    }, nil
}
func (r *CategoryRepo) UpdateCategory(req *pb.CategoryUpdateReq) (*pb.Void,error) {
	query := `UPDATE categories SET name = $1 where id = $2`

    _, err := r.db.Exec(query, req.Body.Name, req.Id)
    if err!= nil {
        return nil, err
    }

    return &pb.Void{}, nil
}
func (r *CategoryRepo) DeleteCategory(byId *pb.ById) (*pb.Void,error) {
    query := `DELETE FROM categories WHERE id = $1`
    _, err := r.db.Exec(query, byId.Id)
    return &pb.Void{},err
}