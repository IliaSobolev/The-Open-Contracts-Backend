package mongo

import (
	"TOC/pkg/domain"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	c *mongo.Collection
}

func NewCodeblockRepo(db *mongo.Database) domain.CodeBlockRepository {
	return &repo{db.Collection("codeblocks")}
}

func (r repo) Create(ctx context.Context, codeBlock *domain.CodeBlock) error {
	_, err := r.c.InsertOne(ctx, codeBlock)
	if err != nil {
		return errors.Wrap(err, "mongo")
	}
	return nil
}

func (r repo) Get(ctx context.Context, id string) (*domain.CodeBlock, error) {
	res := r.c.FindOne(ctx, bson.M{"_id": id})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, domain.ErrCodeBlockNotFound
	} else if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "mongo")
	}
	var obj domain.CodeBlock
	err := res.Decode(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}
	return &obj, nil
}

func (r repo) List(ctx context.Context, params domain.SearchParams) ([]*domain.CodeBlock, error) {
	pipeline := bson.A{}

	matchStage := bson.M{}
	if params.AuthorName != nil {
		matchStage["author_name"] = *params.AuthorName
	}
	if params.Lang != nil {
		matchStage["lang"] = *params.Lang
	}
	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}
	sortOrderValue := -1
	if params.SortOrder == "asc" {
		sortOrderValue = 1
	}
	sortStage := bson.M{
		"$sort": bson.M{
			"rating": sortOrderValue,
		},
	}
	pipeline = append(pipeline, sortStage)
	fmt.Println(pipeline)
	cursor, err := r.c.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*domain.CodeBlock
	for cursor.Next(ctx) {
		var codeBlock domain.CodeBlock
		if err := cursor.Decode(&codeBlock); err != nil {
			return nil, fmt.Errorf("decoding failed: %v", err)
		}
		results = append(results, &codeBlock)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	//var results []*domain.CodeBlock
	//if err := cursor.All(ctx, &results); err != nil {
	//	return nil, err
	//}
	if len(results) == 0 {
		return nil, domain.ErrCodeBlockNotFound
	}
	return results, nil

	//opts := options.Find().SetSort(bson.D{{"rating", -1}})
	//
	//res, err := r.c.Find(ctx, opts)
	//if err != nil {
	//	return nil, domain.ErrCodeBlockNotFound
	//}
	//var codeBlocks []*domain.CodeBlock
	//if err := res.All(ctx, &codeBlocks); err != nil {
	//	return nil, domain.ErrCodeBlockNotFound
	//}
	//return codeBlocks, nil
}
