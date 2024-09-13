package mongodb

import (
	"context"
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
	postCollectionName = "posts"
)

type postRepo struct {
	col *mongo.Collection
}

func NewPostManager(db *mongo.Database) repository.Posts {
	return &postRepo{
		col: db.Collection(postCollectionName),
	}
}
func (r *postRepo) CreatePost(ctx context.Context, post *entity.PostCreate) error {
	now := time.Now().Format(time.RFC3339)

	posts := bson.M{
		"Title":       post.Title,
		"Description": post.Content,
		"PictureUrls": post.PictureUrls,
		"CreatedAt":   now,
		"UpdatedAt":   now,
		"DeletedAt":   0,
	}

	_, err := r.col.InsertOne(ctx, posts)
	if err != nil {
		return err
	}

	return nil

}

func (r *postRepo) UpdatePost(ctx context.Context, post *entity.PostUpdate) error {
	obj_id, err := primitive.ObjectIDFromHex(post.ID)
	if err != nil {
		return err
	}

	updateFields := bson.M{}
	if post.Body.Title != "" {
		updateFields["Title"] = post.Body.Title
	}
	if post.Body.Content != "" {
		updateFields["Content"] = post.Body.Content
	}

	updateFields["UpdatedAt"] = time.Now().Format(time.RFC3339)

	filter := bson.M{"_id": obj_id, "DeletedAt": 0}
	update := bson.M{"$set": updateFields}

	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	log.Println("updated post")

	return nil

}

func (r *postRepo) DeletePost(ctx context.Context, postID string) error {
	deletedAt := time.Now().Unix()
	obj_id, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": obj_id, "DeletedAt": 0}
	update := bson.M{"$set": bson.M{"DeletedAt": deletedAt}}

	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	log.Println("deleted post")

	return nil

}

func (r *postRepo) GetPosts(ctx context.Context, post *entity.Pagination) (*entity.PostList, error) {
	if r.col == nil {
		return nil, fmt.Errorf("col is not initialized")
	}

	if post == nil {
		return nil, fmt.Errorf("request is nil")
	}


	filter := bson.M{"DeletedAt": 0}

	options := options.Find()
	options.SetLimit(int64(post.Limit))
	options.SetSkip(int64(post.Offset))

	projection := bson.M{
		"_id":         1,
		"Title":       1,
		"Content":     1,
		"PictureUrls": 1,
		"CreatedAt":   1,
	}
	options.SetProjection(projection)


	cursor, err := r.col.Find(ctx, filter, options)
	if err != nil {
		log.Println("Error finding posts:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []*entity.PostGet
	for cursor.Next(ctx) {
		var post entity.PostGet
		if err := cursor.Decode(&post); err != nil {
			log.Println("Error decoding post:", err)
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error after processing:", err)
		return nil, err
	}

	totalCount, err := r.col.CountDocuments(ctx, filter)
	if err != nil {
		log.Println("Error counting documents:", err)
		return nil, err
	}

	log.Println("Total count of posts:", totalCount)

	return &entity.PostList{
		Posts:      posts,
		TotalCount: int32(totalCount),
		Pagination: entity.Pagination{
			Limit:   post.Limit,
			Offset:  post.Offset},
	}, nil
}


func (r *postRepo) GetPost(ctx context.Context, postID string) (*entity.PostGet, error) {
	var posts entity.PostGet
	obj_id, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, err
	}
	log.Println(obj_id)
	filter := bson.M{"_id": obj_id, "DeletedAt": 0}
	projection := bson.M{
		"_id":       1,
		"Title":     1,
		"Content":   1,
		"PriceUrls": 1,
		"CreatedAt": 1,
		"UpdatedAt": 1,
	}
	var acc Posts

	err = r.col.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&acc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	posts.ID = postID
	posts.Title = acc.Title
	posts.Content = acc.Content
	posts.PictureUrls = acc.PictureUrls
	posts.CreatedAt = acc.CreatedAt

	return &posts, nil

}

func (r *postRepo) AddPostPicture(ctx context.Context, postpic *entity.PostPicture) error {

	pictures := bson.M{
		"PostID":      postpic.PostID,
		"PictureUrls": postpic.PictureUrl,
	}

	_, err := r.col.InsertOne(ctx, pictures)
	if err != nil {
		return err
	}

	return nil
}
func (r *postRepo) DeletePostPicture(ctx context.Context, postpic *entity.PostPicture) error {
	objID, err := primitive.ObjectIDFromHex(postpic.PostID)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id":       objID,
		"DeletedAt": 0,
	}

	update := bson.M{
		"$pull": bson.M{"PictureUrls": postpic.PictureUrl},
	}

	res, err := r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	log.Println("Picture deleted for post ID:", postpic.PostID)

	return nil
}

type Posts struct {
	Title       string   `bson:"Title"`
	Content     string   `bson:"Content"`
	PictureUrls []string `bson:"PriceUrls"`
	CreatedAt   string   `bson:"CreatedAt"`
	UpdatedAt   string   `bson:"UpdatedAt"`
}
