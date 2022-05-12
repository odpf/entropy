package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/odpf/entropy/core/resource"
	"github.com/odpf/entropy/pkg/errors"
)

const resourceRepoName = "resources"

func NewResourceRepository(db *mongo.Database) *ResourceRepository {
	return &ResourceRepository{
		coll: db.Collection(resourceRepoName),
	}
}

type ResourceRepository struct{ coll *mongo.Collection }

func (rc *ResourceRepository) GetByURN(ctx context.Context, urn string) (*resource.Resource, error) {
	var rm resourceModel

	filter := map[string]interface{}{"urn": urn}
	if err := rc.coll.FindOne(ctx, filter).Decode(&rm); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return modelToResource(rm), nil
}

func (rc *ResourceRepository) List(ctx context.Context, filter map[string]string) ([]*resource.Resource, error) {
	cur, err := rc.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var records []resourceModel
	if err = cur.All(ctx, &records); err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	var res []*resource.Resource
	for _, rec := range records {
		res = append(res, modelToResource(rec))
	}
	return res, nil
}

func (rc *ResourceRepository) Create(ctx context.Context, res resource.Resource) error {
	res.CreatedAt = time.Now()
	res.UpdatedAt = time.Now()

	_, err := rc.coll.InsertOne(ctx, modelFromResource(res))
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.ErrConflict
		}
		return err
	}

	return nil
}

func (rc *ResourceRepository) Update(ctx context.Context, res resource.Resource) error {
	res.UpdatedAt = time.Now()

	filter := map[string]interface{}{"urn": res.URN}
	updates := map[string]interface{}{"$set": modelFromResource(res)}

	_, err := rc.coll.UpdateOne(ctx, filter, updates)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.ErrNotFound
		}
		return err
	}

	return nil
}

func (rc *ResourceRepository) Delete(ctx context.Context, urn string) error {
	_, err := rc.coll.DeleteOne(ctx, map[string]interface{}{"urn": urn})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.ErrNotFound
		}
		return err
	}
	return nil
}

func (rc *ResourceRepository) DoPending(ctx context.Context, fn resource.PendingHandler) error {
	filter := map[string]interface{}{"state.status": resource.StatusPending}

	var rec resourceModel
	if err := rc.coll.FindOne(ctx, filter).Decode(&rec); err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.ErrNotFound // no pending item is available.
		}
		return err
	}

	res := modelToResource(rec)
	modified, shouldDelete, err := fn(ctx, *res)
	if err != nil {
		return err
	}

	if shouldDelete {
		_, err := rc.coll.DeleteOne(ctx, map[string]interface{}{"urn": res.URN})
		if err != nil && err != mongo.ErrNoDocuments {
			return err
		}
		return nil
	}

	return rc.Update(ctx, *modified)
}

func (rc *ResourceRepository) Migrate(ctx context.Context) error {
	return createUniqueIndex(ctx, rc.coll, "urn", 1)
}

func createUniqueIndex(ctx context.Context, collection *mongo.Collection, key string, order int) error {
	_, err := collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: key, Value: order}},
			Options: options.Index().SetUnique(true),
		},
	)
	return err
}
