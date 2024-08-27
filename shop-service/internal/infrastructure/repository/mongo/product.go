package mongodb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"gitlab.com/acumen5524834/shop-service/internal/entity"
	"gitlab.com/acumen5524834/shop-service/internal/infrastructure/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	productCollectionName = "products"
)

type productRepo struct {
	col *mongo.Collection
}

func NewProductManager(db *mongo.Database) repository.Products {
	return &productRepo{
		col: db.Collection(productCollectionName),
	}
}
func (r *productRepo) CreateProduct(ctx context.Context, product *entity.ProductCreate) error {
	now := time.Now().Format(time.RFC3339)

	products := bson.M{
		"Title":       product.Title,
		"Description": product.Description,
		"Price":       product.Price,
		"PictureUrls": product.PictureUrls,
		"CreatedAt":   now,
		"UpdatedAt":   now,
		"DeletedAt":   0,
	}

	_, err := r.col.InsertOne(ctx, products)
	if err != nil {
		return err
	}

	return nil

}

func (r *productRepo) UpdateProduct(ctx context.Context, product *entity.ProductUpdate) error {
	obj_id, err := primitive.ObjectIDFromHex(product.Id)
	if err != nil {
		return err
	}

	updateFields := bson.M{}
	if product.Body.Description != "" {
		updateFields["Description"] = product.Body.Description
	}
	if product.Body.Title != "" {
		updateFields["Title"] = product.Body.Title
	}
	if product.Body.Price != 0.0 {
		updateFields["Price"] = product.Body.Price
	}

	updateFields["UpdatedAt"] = time.Now().Format(time.RFC3339)

	filter := bson.M{"_id": obj_id, "DeletedAt": 0}
	update := bson.M{"$set": updateFields}

	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	log.Println("updated product")

	return nil

}

func (r *productRepo) DeleteProduct(ctx context.Context, productID string) error {
	deletedAt := time.Now().Unix()
	obj_id, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": obj_id, "DeletedAt": 0}
	update := bson.M{"$set": bson.M{"DeletedAt": deletedAt}}

	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	log.Println("deleted product")

	return nil

}
func (r *productRepo) ListProducts(ctx context.Context, products *entity.ProductFilter) (*entity.ProductList, error) {
	if r.col == nil {
		return nil, fmt.Errorf("col is not initialized")
	}

	if products == nil {
		return nil, fmt.Errorf("request is nil")
	}

	if products.Pagination.Limit <= 0 {
		return nil, fmt.Errorf("pagination limit must be greater than 0")
	}
	if products.Pagination.Offset < 0 {
		return nil, fmt.Errorf("pagination offset cannot be negative")
	}

	filter := bson.M{"DeletedAt": 0}
	if products.Title != "" {
		filter["Title"] = products.Title
	}
	if products.PriceFrom != 0.0 {
		filter["Price"] = bson.M{"$gte": products.PriceFrom}
	}
	if products.PriceTo != 0.0 {
		filter["Price"] = bson.M{"$lte": products.PriceTo}
	}

	options := options.Find()
	options.SetLimit(int64(products.Pagination.Limit))
	options.SetSkip(int64(products.Pagination.Offset))

	projection := bson.M{
		"_id":         1,
		"Title":       1,
		"Description": 1,
		"Price":       1,
		"PictureUrls": 1,
		"CreatedAt":   1,
	}
	options.SetProjection(projection)

	cursor, err := r.col.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var productss []*entity.ProductGet
	for cursor.Next(ctx) {
		var products entity.ProductGet
		if err := cursor.Decode(&products); err != nil {
			return nil, err
		}
		productss = append(productss, &products)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	totalCount, err := r.col.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &entity.ProductList{
		Products:   productss,
		TotalCount: int64(totalCount),
	}, nil
}

func (r *productRepo) GetProduct(ctx context.Context, productID string) (*entity.ProductGet, error) {
	var prod entity.ProductGet
	obj_id, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, err
	}
	log.Println(obj_id)
	filter := bson.M{"_id": obj_id, "DeletedAt": 0}
	projection := bson.M{
		"_id":         1,
		"Title":       1,
		"Description": 1,
		"Price":       1,
		"PriceUrls":   1,
		"CreatedAt":   1,
		"UpdatedAt":   1,
	}
	var acc Products

	err = r.col.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&acc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("products not found")
		}
		return nil, err
	}

	prod.Id = productID
	prod.Title = acc.Title
	prod.Description = acc.Description
	prod.Price = acc.Price
	prod.PictureUrls = acc.PictureUrls
	prod.CreatedAt = acc.CreatedAt

	return &prod, nil

}

func (r *productRepo) AddPicture(ctx context.Context, productpic *entity.ProductPicture) error {

	pictures := bson.M{
		"ProductId":   productpic.ProductId,
		"PictureUrls": productpic.PictureUrl,
	}

	_, err := r.col.InsertOne(ctx, pictures)
	if err != nil {
		return err
	}

	return nil
}
func (r *productRepo) DeletePicture(ctx context.Context, productpic *entity.ProductPicture) error {
	// Validate the ProductId to ensure it's a valid MongoDB ObjectID
	if !primitive.IsValidObjectID(productpic.ProductId) {
		return errors.New("invalid ObjectID format")
	}

	// Convert ProductId string to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(productpic.ProductId)
	if err != nil {
		return err
	}

	// Define the filter for the document to be updated
	filter := bson.M{
		"_id":          objID,
		"picture_urls": productpic.PictureUrl,
		"DeletedAt":    0,
	}

	// Update the document to set the DeletedAt timestamp
	res, err := r.col.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"DeletedAt": time.Now().Unix()}})
	if err != nil {
		return err
	}

	// Check if any document was matched and updated
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	// Log the successful deletion
	log.Println("Picture deleted for product ID:", productpic.ProductId)

	return nil
}

type Products struct {
	Title       string   `bson:"Title"`
	Description string   `bson:"Description"`
	Price       float32  `bson:"Price"`
	PictureUrls []string `bson:"PriceUrls"`
	CreatedAt   string   `bson:"CreatedAt"`
	UpdatedAt   string   `bson:"UpdatedAt"`
}
