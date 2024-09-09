package test

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"gitlab.com/acumen5524834/shop-service/internal/entity"
// 	repo "gitlab.com/acumen5524834/shop-service/internal/infrastructure/repository/mongo"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func TestCreateBook(t *testing.T) {
// 	mockCol := new(MockMongoCollection)
// 	bookRepo := &repo.BookRepo{
// 		Col: mockCol,  // Inject the mock collection here
// 	}

// 	mockCol.On("InsertOne", mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{}, nil)

// 	book := &entity.BookCreate{
// 		Title:       "Test Book",
// 		Description: "Test Description",
// 		Price:       9.99,
// 		PictureUrls: []string{"https://example.com/image.jpg"},
// 	}

// 	err := bookRepo.CreateBook(context.Background(), book)
// 	assert.NoError(t, err)
// 	mockCol.AssertExpectations(t)
// }

// // Example of a test for UpdateBook
// func TestUpdateBook(t *testing.T) {
// 	mockCol := new(MockMongoCollection)
// 	bookRepo := &repo.BookRepo{
// 		Col: mockCol,  // Inject the mock collection here
// 	}

// 	book := &entity.BookUpdate{
// 		Id: "5f8d0d55b54764421b7156b0",
// 		Body: entity.BookUpt{
// 			Title:       "Updated Title",
// 			Description: "Updated Description",
// 			Price:       19.99,
// 		},
// 	}

// 	objID, _ := primitive.ObjectIDFromHex(book.Id)
// 	mockCol.On("UpdateOne", mock.Anything, bson.M{"_id": objID, "DeletedAt": 0}, mock.Anything).Return(&mongo.UpdateResult{MatchedCount: 1}, nil)

// 	err := bookRepo.UpdateBook(context.Background(), book)
// 	assert.NoError(t, err)
// 	mockCol.AssertExpectations(t)
// }

// // Example of a test for DeleteBook
// func TestDeleteBook(t *testing.T) {
// 	mockCol := new(MockMongoCollection)
// 	bookRepo := &repo.BookRepo{
// 		Col: mockCol,  // Inject the mock collection here
// 	}

// 	bookID := "5f8d0d55b54764421b7156b0"
// 	objID, _ := primitive.ObjectIDFromHex(bookID)
// 	mockCol.On("UpdateOne", mock.Anything, bson.M{"_id": objID, "DeletedAt": 0}, mock.Anything).Return(&mongo.UpdateResult{MatchedCount: 1}, nil)

// 	err := bookRepo.DeleteBook(context.Background(), bookID)
// 	assert.NoError(t, err)
// 	mockCol.AssertExpectations(t)
// }

// // Example of a test for GetBook
// func TestGetBook(t *testing.T) {
// 	mockCol := new(MockMongoCollection)
// 	bookRepo := &repo.BookRepo{

// 	}

// 	bookID := "5f8d0d55b54764421b7156b0"
// 	objID, _ := primitive.ObjectIDFromHex(bookID)

// 	mockSingleResult := new(mongo.SingleResult)
// 	mockCol.On("FindOne", mock.Anything, bson.M{"_id": objID, "DeletedAt": 0}).Return(mockSingleResult)

// 	_, err := bookRepo.GetBook(context.Background(), bookID)
// 	assert.NoError(t, err)
// 	mockCol.AssertExpectations(t)
// }
