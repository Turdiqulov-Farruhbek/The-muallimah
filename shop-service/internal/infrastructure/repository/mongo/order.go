package mongodb

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
	pb "gitlab.com/acumen5524834/shop-service/internal/pkg/genproto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepository struct {
	db          *sql.DB
	mongoClient *mongo.Client
}

func NewOrderRepository(db *sql.DB, mongoClient *mongo.Client) *OrderRepository {

	return &OrderRepository{
		db:          db,
		mongoClient: mongoClient,
	}
}

func (o *OrderRepository) CreateOrder(c context.Context, req *pb.OrderCreateReq) (*pb.Void, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO orders
		(id, user_id, item_id,type, quantity, total_price,status)
		VALUES
		($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := o.db.Exec(query, id, req.UserId, req.ItemId, req.Type, req.Quantity, req.TotalPrice, req.Status)

	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (o *OrderRepository) GetOrder(c context.Context, id *pb.OrderGetReq) (*pb.OrderGetRes, error) {
	if err := o.mongoClient.Ping(c, nil); err != nil {
		if reconnectErr := o.reconnectMongoClient(c); reconnectErr != nil {
			return nil, fmt.Errorf("failed to reconnect to MongoDB: %v", reconnectErr)
		}
	}

	query := `
	SELECT 
		id,
		user_id,
		item_id,
		type,
		quantity,
		total_price,
		status
	FROM 
		orders
	WHERE 
		id = $1
	`
	row := o.db.QueryRow(query, id.Id)

	var order pb.OrderGetRes
	order.Order = &pb.Order{}
	if err := row.Scan(
		&order.Order.Id,
		&order.Order.UserId,
		&order.Order.ItemId,
		&order.Order.Type,
		&order.Order.Quantity,
		&order.Order.TotalPrice,
		&order.Order.Status); err != nil {
		return nil, err
	}

	var collectionName string
	switch order.Order.Type {
	case "book":
		collectionName = "books"
	case "product":
		collectionName = "products"
	default:
		return nil, fmt.Errorf("unknown type: %s", order.Order.Type)
	}

	objectID, err := primitive.ObjectIDFromHex(order.Order.ItemId)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectId format: %v", err)
	}

	var item struct {
		Price float64 `bson:"price"`
	}

	err = o.mongoClient.Database("shop").Collection(collectionName).
		FindOne(c, bson.M{"_id": objectID}).
		Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no %s found with ID %s", order.Order.Type, order.Order.ItemId)
		}
		return nil, fmt.Errorf("failed to fetch %s details: %v", order.Order.Type, err)
	}

	totalPrice := item.Price * float64(order.Order.Quantity)
	order.Order.TotalPrice = fmt.Sprintf("%.2f", totalPrice)

	totalPriceFloat, err := strconv.ParseFloat(order.Order.TotalPrice, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse total price: %v", err)
	}

	updateQuery := `
			UPDATE 
				orders
			SET 
				total_price = $1
			WHERE 
				id = $2
			`
	_, updateErr := o.db.Exec(updateQuery, totalPriceFloat, order.Order.Id)
	if updateErr != nil {
		return nil, fmt.Errorf("failed to update order in PostgreSQL: %v", updateErr)
	}

	return &order, nil
}

func (o *OrderRepository) reconnectMongoClient(c context.Context) error {
	if err := o.mongoClient.Disconnect(c); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %v", err)
	}

	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	newClient, err := mongo.Connect(c, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to reconnect to MongoDB: %v", err)
	}

	if err := newClient.Ping(c, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB after reconnecting: %v", err)
	}

	o.mongoClient = newClient
	return nil
}

func (o *OrderRepository) UpdateOrder(c context.Context, req *pb.OrderUpdateReq) (*pb.Void, error) {
	var args []interface{}
	var conditions []string

	if req.Body.ItemId != "" && req.Body.ItemId != "string" {
		args = append(args, req.Body.ItemId)
		conditions = append(conditions, fmt.Sprintf("item_id = $%d", len(args)))
	}
	if req.Body.Type != "" && req.Body.Type != "string" {
		args = append(args, req.Body.Type)
		conditions = append(conditions, fmt.Sprintf("type = $%d", len(args)))
	}
	if req.Body.Quantity != 0 {
		args = append(args, req.Body.Quantity)
		conditions = append(conditions, fmt.Sprintf("quantity = $%d", len(args)))
	}
	if req.Body.TotalPrice != "" && req.Body.TotalPrice != "string" {
		args = append(args, req.Body.TotalPrice)
		conditions = append(conditions, fmt.Sprintf("total_price = $%d", len(args)))
	}
	if req.Body.Status != "" && req.Body.Status != "string" {
		args = append(args, req.Body.Status)
		conditions = append(conditions, fmt.Sprintf("status = $%d", len(args)))
	}

	query := `UPDATE orders SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
	args = append(args, req.Id)

	_, err := o.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating orders", err)
		return nil, err
	}

	return &pb.Void{}, nil

}

func (o *OrderRepository) DeleteOrder(c context.Context, id *pb.OrderDeleteReq) (*pb.Void, error) {
	query := `
	UPDATE
		orders
	SET
		deleted_at = extract(epoch from now())
	WHERE
		id = $1
	`

	_, err := o.db.Exec(query, id.Id)
	if err != nil {
		log.Println("Error while deleting orders", err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (o *OrderRepository) ListOrders(c context.Context, req *pb.OrderListsReq) (*pb.OrderListsRes, error) {
	query := `
	SELECT
		id,
		user_id,
		item_id,
		type,
		quantity,
		total_price,
		status
	FROM
		orders
	`

	var args []interface{}
	argCount := 1
	filters := []string{}

	if req.UserId != "" && req.UserId != "string" {
		filters = append(filters, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, req.UserId)
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}
	if req.Filter != nil {
		if req.Filter.Limit > 0 {
			query += fmt.Sprintf(" LIMIT $%d", argCount)
			args = append(args, req.Filter.Limit)
			argCount++
		}
		if req.Filter.Offset > 0 {
			query += fmt.Sprintf(" OFFSET $%d", argCount)
			args = append(args, req.Filter.Offset)
			argCount++
		}
	}

	rows, err := o.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []*pb.Order{}

	for rows.Next() {
		var order pb.Order
		err := rows.Scan(
			&order.Id,
			&order.UserId,
			&order.ItemId,
			&order.Type,
			&order.Quantity,
			&order.TotalPrice,
			&order.Status,
		)
		if err != nil {
			log.Println("no rows result set")
			return nil, err
		}
		orders = append(orders, &order)
	}
	totalCount := len(orders)

	return &pb.OrderListsRes{
		Orders:     orders,
		TotalCount: int32(totalCount),
	}, nil

}
