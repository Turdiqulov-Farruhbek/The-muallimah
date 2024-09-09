package repo

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	pb "gitlab.com/muallimah/course_service/internal/pkg/genproto"
)

type TransactionRepo struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (r *TransactionRepo) CreateTransaction(req *pb.TransactionCreateReq) (*pb.Void, error) {
	if req.UserCourseId != "" && req.UserCourseId != "string" {
		price_query := `select c.price, 
								uc.user_id 
						FROM courses c 
						LEFT JOIN user_courses uc 
						on uc.course_id = c.id
						where uc.id = $1
		`
		var price float32
		var user_id string
		err := r.db.QueryRow(price_query, req.UserCourseId).Scan(&price, &user_id)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user course not found")
		}
		if price == 0 {
			return nil, fmt.Errorf("course price not found")
		}
		req.Amount = price
		req.Type = "credit"
		req.UserId = user_id

		id := uuid.New().String()
		query := `INSERT INTO transactions (id, 
											user_course_id, 
											amount, 
											type,
											user_id,
											created_at) 
					VALUES ($1, $2, $3, $4, $5, $6)`

		_, err = r.db.Exec(query,
			id,
			req.UserCourseId,
			req.Amount,
			req.Type,
			req.UserId,
			time.Now())

		if err != nil {
			return nil, err
		}
		return &pb.Void{}, nil
	}
	id := uuid.New().String()
	query := `INSERT INTO transactions (id,  
										amount, 
										type,
										user_id,
										created_at) 
				VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(query,
		id,
		req.Amount,
		req.Type,
		req.UserId,
		time.Now())

	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *TransactionRepo) GetTransaction(id *pb.ById) (*pb.TransactionGetRes, error) {
	query := `SELECT id, 
                    user_course_id, 
                    amount, 
                    type, 
					user_id,
					FROM transactions
					WHERE id = $1
    `
	var user_course_id sql.NullString
	var transaction pb.TransactionGetRes
	err := r.db.QueryRow(query, id.Id).Scan(
		&transaction.Id,
		&user_course_id,
		&transaction.Amount,
		&transaction.Type,
		&transaction.User.Id,
	)
	if err != nil {
		return nil, err
	}
	if user_course_id.Valid {
		transaction.UserCourse.Id = user_course_id.String
	}
	return &transaction, nil
}
func (r *TransactionRepo) ListTransactions(req *pb.TransactionListReq) (*pb.TransactionListRes, error) {
	query := `SELECT t.id, 
                    t.user_course_id, 
                    t.amount, 
                    t.type,
					t.user_id
                    FROM transactions t 
					LEFT JOIN user_courses uc on t.user_course_id = uc.id
					LEFT JOIN courses c ON uc.
                    `
	var cons []string
	var args []interface{}
	if req.UserId != "" && req.UserId != "string" {
		cons = append(cons, fmt.Sprintf("u.id = $%d", len(args)+1))
		args = append(args, req.UserId)
	}
	if req.CourseId != "" && req.CourseId != "string" {
		cons = append(cons, fmt.Sprintf("c.id = $%d", len(args)+1))
		args = append(args, req.CourseId)
	}
	if req.UserCourseId != "" && req.UserCourseId != "string" {
		cons = append(cons, fmt.Sprintf("t.user_course_id = $%d", len(args)+1))
		args = append(args, req.UserCourseId)
	}
	if req.AmountFrom != 0.0 && req.AmountFrom != 0 {
		cons = append(cons, fmt.Sprintf("t.amount >= $%d", len(args)+1))
		args = append(args, req.AmountFrom)
	}
	if req.AmountTo != 0.0 && req.AmountTo != 0 {
		cons = append(cons, fmt.Sprintf("t.amount <= $%d", len(args)+1))
		args = append(args, req.AmountTo)
	}
	if req.Type != "" && req.Type != "string" {
		cons = append(cons, fmt.Sprintf("t.type = $%d", len(args)+1))
		args = append(args, req.Type)
	}
	if len(cons) > 0 {
		query += " WHERE " + strings.Join(cons, " AND ")
	}
	query += " ORDER BY t.created_at DESC"
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
	var transactions pb.TransactionListRes
	for rows.Next() {
		var transaction pb.TransactionGetRes
		var user_course_id sql.NullString
		err := rows.Scan(
			&transaction.Id,
			&user_course_id,
			&transaction.Amount,
			&transaction.Type,
		)
		if err != nil {
			return nil, err
		}
		if user_course_id.Valid {
			transaction.UserCourse.Id = user_course_id.String
		}
		transactions.Transactions = append(transactions.Transactions, &transaction)
	}
	return &transactions, nil
}

func (r *TransactionRepo) GetBalance(req *pb.ById) (*pb.BalanceGetRes, error) {
	res := pb.BalanceGetRes{}
	query := `
	WITH user_credit AS (
		SELECT 
			COALESCE(SUM(amount), 0) AS total_credit
		FROM 
			transactions 
		WHERE 
			user_id = $1 
			AND type = 'credit'
	),
	user_debit AS (
		SELECT 
			COALESCE(SUM(CASE 
				WHEN t.user_course_id IS NOT NULL THEN c.price
				WHEN t.order_id IS NOT NULL THEN o.total_price
				ELSE t.amount
			END), 0) AS total_debit
		FROM 
			transactions t
		LEFT JOIN user_courses uc ON t.user_course_id = uc.id
		LEFT JOIN courses c ON uc.course_id = c.id
		LEFT JOIN orders o ON t.order_id = o.id
		WHERE 
			t.user_id = $1 
			AND t.type = 'debit'
	)
	SELECT 
		(SELECT total_credit FROM user_credit) - (SELECT total_debit FROM user_debit) AS balance;
	`
	err := r.db.QueryRow(query, req.Id).Scan(
		&res.Balance,
	)
	if err != nil {
		return nil, err
	}
	fmt.Println(&res)
	return &res, nil
}
