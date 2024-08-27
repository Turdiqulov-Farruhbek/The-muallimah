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
	bookCollectionName = "books"
)

type bookRepo struct {
	col *mongo.Collection
}

func NewBookManager(db *mongo.Database) repository.Books {
	return &bookRepo{
		col: db.Collection(bookCollectionName),
	}
}
func (r *bookRepo) CreateBook(ctx context.Context, book *entity.BookCreate) error {
	now := time.Now().Format(time.RFC3339)

	books := bson.M{
		"Title":       book.Title,
		"Description": book.Description,
		"Price":       book.Price,
		"PictureUrls": book.PictureUrls,
		"CreatedAt":   now,
		"UpdatedAt":   now,
		"DeletedAt":   0,
	}

	_, err := r.col.InsertOne(ctx, books)
	if err != nil {
		return err
	}

	return nil

}

func (r *bookRepo) UpdateBook(ctx context.Context, book *entity.BookUpdate) error {
	obj_id, err := primitive.ObjectIDFromHex(book.Id)
	if err != nil {
		return err
	}

	updateFields := bson.M{}
	if book.Body.Description != "" {
		updateFields["Description"] = book.Body.Description
	}
	if book.Body.Title != "" {
		updateFields["Title"] = book.Body.Title
	}
	if book.Body.Price != 0.0 {
		updateFields["Price"] = book.Body.Price
	}

	updateFields["UpdatedAt"] = time.Now().Format(time.RFC3339)

	filter := bson.M{"_id": obj_id, "DeletedAt": 0}
	update := bson.M{"$set": updateFields}

	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	log.Println("updated book")

	return nil

}

func (r *bookRepo) DeleteBook(ctx context.Context, bookID string) error {
	deletedAt := time.Now().Unix()
	obj_id, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": obj_id, "DeletedAt": 0}
	update := bson.M{"$set": bson.M{"DeletedAt": deletedAt}}

	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	log.Println("deleted book")

	return nil

}

func (r *bookRepo) ListBooks(ctx context.Context, book *entity.BookFilter) (*entity.BookList, error) {
	if r.col == nil {
		return nil, fmt.Errorf("col is not initialized")
	}

	if book == nil {
		return nil, fmt.Errorf("request is nil")
	}

	if book.Pagination.Limit <= 0 {
		return nil, fmt.Errorf("pagination limit must be greater than 0")
	}
	if book.Pagination.Offset < 0 {
		return nil, fmt.Errorf("pagination offset cannot be negative")
	}

	filter := bson.M{"DeletedAt": 0}
	if book.Title != "" {
		filter["Title"] = book.Title
	}
	if book.PriceFrom != 0.0 {
		filter["Price"] = bson.M{"$gte": book.PriceFrom}
	}
	if book.PriceTo != 0.0 {
		filter["Price"] = bson.M{"$lte": book.PriceTo}
	}

	options := options.Find()
	options.SetLimit(int64(book.Pagination.Limit))
	options.SetSkip(int64(book.Pagination.Offset))

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

	var books []*entity.BookGet
	for cursor.Next(ctx) {
		var book entity.BookGet
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	totalCount, err := r.col.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &entity.BookList{
		Books:      books,
		TotalCount: int64(totalCount),
	}, nil
}

func (r *bookRepo) GetBook(ctx context.Context, bookID string) (*entity.BookGet, error) {
	var books entity.BookGet
	obj_id, err := primitive.ObjectIDFromHex(bookID)
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
	var acc BOoks

	err = r.col.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&acc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	books.Id = bookID
	books.Title = acc.Title
	books.Description = acc.Description
	books.Price = acc.Price
	books.PictureUrls = acc.PictureUrls
	books.CreatedAt = acc.CreatedAt

	return &books, nil

}

func (r *bookRepo) AddPicture(ctx context.Context, picture *entity.BookPicture) error {

	pictures := bson.M{
		"BookID":      picture.BookId,
		"PictureUrls": picture.PictureUrl,
	}

	_, err := r.col.InsertOne(ctx, pictures)
	if err != nil {
		return err
	}

	return nil
}
func (r *bookRepo) DeletePicture(ctx context.Context, bookpic *entity.BookPicture) error {
	// Validate the BookId to ensure it's a valid MongoDB ObjectID
	if !primitive.IsValidObjectID(bookpic.BookId) {
		log.Printf("Invalid ObjectID format for BookId: %s", bookpic.BookId)
		return errors.New("invalid ObjectID format")
	}

	// Convert BookId string to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(bookpic.BookId)
	if err != nil {
		return err
	}

	// Define the filter for the document to be updated
	filter := bson.M{
		"_id":          objID,
		"picture_urls": bookpic.PictureUrl,
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
	log.Println("Picture deleted for book ID:", bookpic.BookId)

	return nil
}



type BOoks struct {
	Title       string   `bson:"Title"`
	Description string   `bson:"Description"`
	Price       float32  `bson:"Price"`
	PictureUrls []string `bson:"PriceUrls"`
	CreatedAt   string   `bson:"CreatedAt"`
	UpdatedAt   string   `bson:"UpdatedAt"`
}
