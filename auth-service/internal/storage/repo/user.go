package repo

import (
	pb "auth-service/internal/pkg/genproto"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type UserManager struct {
	PgClient *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{PgClient: db}
}

func (u *UserManager) CreateUser(ctx context.Context, req *pb.UserCreateReq) (*pb.Void, error) {
	query := `
		INSERT INTO users (first_name, last_name, dob, phone_number, email, occupation, address, password, gender, photo_url, role)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := u.PgClient.ExecContext(ctx, query, req.FirstName, req.LastName, req.Dob, req.PhoneNumber, req.Email, req.Occupation, req.Address, req.Password, req.Gender, req.PhotoUrl, req.Role)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *UserManager) GetUserByID(ctx context.Context, req *pb.ById) (*pb.UserGetRes, error) {
	query := `
		SELECT id, first_name, last_name, dob, phone_number, email, occupation, address, gender, photo_url, is_confirmed
		FROM users WHERE id = $1 and deleted_at = 0
	`
	var user pb.UserGetRes

	err := u.PgClient.QueryRowContext(ctx, query, req.Id).Scan(
		&user.Id, &user.FirstName, &user.LastName, &user.Dob, &user.PhoneNumber, &user.Email, &user.Occupation, &user.Address, &user.Gender, &user.PhotoUrl, &user.IsConfirmed,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (u *UserManager) GetUserByEmail(ctx context.Context, req *pb.ByEmail) (*pb.UserGetRes, error) {
	query := `
		SELECT id, first_name, last_name, dob, phone_number, email, occupation, address, gender, photo_url, is_confirmed
		FROM users WHERE email = $1 AND deleted_at=0
	`
	var user pb.UserGetRes

	err := u.PgClient.QueryRowContext(ctx, query, req.Email).Scan(
		&user.Id, &user.FirstName, &user.LastName, &user.Dob, &user.PhoneNumber, &user.Email, &user.Occupation, &user.Address, &user.Gender, &user.PhotoUrl, &user.IsConfirmed,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserManager) UpdateUser(ctx context.Context, req *pb.UserUpdateReq) (*pb.Void, error) {
	query := `
		UPDATE users SET first_name = $1, last_name = $2, dob = $3, phone_number = $4, occupation = $5, 
		                  address = $6, gender = $7, photo_url = $8
		WHERE id = $9 AND deleted_at=0
	`

	_, err := u.PgClient.ExecContext(ctx, query, req.Body.FirstName, req.Body.LastName, req.Body.Dob, req.Body.PhoneNumber,
		req.Body.Occupation, req.Body.Address, req.Body.Gender, req.Body.PhotoUrl, req.Id,
	)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *UserManager) ChangeUserPassword(ctx context.Context, req *pb.UserChangePasswordReq) (*pb.Void, error) {
	query := `
		UPDATE users SET password = $1
		WHERE email = $2 AND deleted_at=0
	`

	_, err := u.PgClient.ExecContext(ctx, query, req.NewPassword, req.Email)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *UserManager) DeleteUser(ctx context.Context, req *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE users SET deleted_at = 1 WHERE id = $1
	`
	_, err := u.PgClient.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *UserManager) ListUsers(ctx context.Context, req *pb.UsersGetAllReq) (*pb.UsersGetAllRes, error) {
	var (
		params       []interface{}
		paramCounter = 1
	)

	query := `SELECT id, first_name, last_name, dob, phone_number, email, occupation, address, gender, photo_url, is_confirmed FROM users WHERE deleted_at=0`

	if req.Body.FirstName != "" {
		query += fmt.Sprintf(" AND first_name LIKE $%d", paramCounter)
		params = append(params, "%"+req.Body.FirstName+"%")
		paramCounter++
	}

	if req.Body.LastName != "" {
		query += fmt.Sprintf(" AND last_name LIKE $%d", paramCounter)
		params = append(params, "%"+req.Body.LastName+"%")
		paramCounter++
	}
	if req.Body.PhoneNumber != "" {
		query += fmt.Sprintf(" AND phone_number LIKE $%d", paramCounter)
		params = append(params, "%"+req.Body.PhoneNumber+"%")
		paramCounter++
	}
	if req.Body.Email != "" {
		query += fmt.Sprintf(" AND email LIKE $%d", paramCounter)
		params = append(params, "%"+req.Body.Email+"%")
		paramCounter++
	}
	if req.Body.Occupation != "" {
		query += fmt.Sprintf(" AND occupation LIKE $%d", paramCounter)
		params = append(params, "%"+req.Body.Occupation+"%")
		paramCounter++
	}
	if req.Body.Address != "" {
		query += fmt.Sprintf(" AND address LIKE $%d", paramCounter)
		params = append(params, "%"+req.Body.Address+"%")
		paramCounter++
	}
	if req.Body.Gender != "" {
		query += fmt.Sprintf(" AND gender = $%d", paramCounter)
		params = append(params, req.Body.Gender)
		paramCounter++
	}
	if req.Body.DobFrom != "" {
		layout := "2006-01-02"
		dobFrom, err := time.Parse(layout, req.Body.DobFrom)
		if err == nil {
			query += fmt.Sprintf(" AND dob >= $%d", paramCounter)
			params = append(params, dobFrom)
			paramCounter++
		}
	}
	if req.Body.DobTo != "" {
		layout := "2006-01-02"
		dobTo, err := time.Parse(layout, req.Body.DobTo)
		if err == nil {
			query += fmt.Sprintf(" AND dob <= $%d", paramCounter)
			params = append(params, dobTo)
			paramCounter++
		}
	}
	if req.Pagination.Limit != 0 {
		query += fmt.Sprintf(" LIMIT $%d", paramCounter)
		params = append(params, req.Pagination.Limit)
		paramCounter++
	}
	if req.Pagination.Offset != 0 {
		query += fmt.Sprintf(" OFFSET $%d", paramCounter)
		params = append(params, req.Pagination.Offset)
		paramCounter++
	}
	query += " ORDER BY created_at DESC "

	rows, err := u.PgClient.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*pb.UserGetRes
	for rows.Next() {
		var user pb.UserGetRes
		err := rows.Scan(
			&user.Id, &user.FirstName, &user.LastName, &user.Dob, &user.PhoneNumber, &user.Email, &user.Occupation, &user.Address, &user.Gender, &user.PhotoUrl, &user.IsConfirmed,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return &pb.UsersGetAllRes{
		Users:      users,
		Pagination: req.Pagination,
	}, nil
}

func (u *UserManager) IsEmailExists(ctx context.Context, req *pb.UserEmailCheckReq) (*pb.UserEmailCheckRes, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1 AND deleted_at=0)`
	err := u.PgClient.QueryRowContext(ctx, query, req.Email).Scan(&exists)
	if err != nil {
		return nil, err
	}
	return &pb.UserEmailCheckRes{Exists: exists}, nil
}

func (u *UserManager) GetUserSecurityByEmail(ctx context.Context, req *pb.ByEmail) (*pb.UserSecurityRes, error) {
	query := `
		SELECT id, email, password, role, is_confirmed
		FROM users WHERE email = $1 AND deleted_at=0
	`

	var user pb.UserSecurityRes

	err := u.PgClient.QueryRowContext(ctx, query, req.Email).Scan(
		&user.Id, &user.Email, &user.Password, &user.Role, &user.IsConfirmed,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserManager) ConfirmUser(ctx context.Context, req *pb.ByEmail) (*pb.Void, error) {
	query := `
		UPDATE users SET is_confirmed = true WHERE email = $1 AND deleted_at=0
	`
	_, err := u.PgClient.ExecContext(ctx, query, req.Email)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *UserManager) ChangeUserPFP(ctx context.Context, req *pb.UserChangePFPReq) (*pb.Void, error) {
	query := `UPDATE users SET photo_url = $1 WHERE email = $2 AND deleted_at=0`
	_, err := u.PgClient.ExecContext(ctx, query, req.PhotoUrl, req.Email)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
