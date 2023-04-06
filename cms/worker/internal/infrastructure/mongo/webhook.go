package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/reearth/reearth-cms/worker/internal/usecase/repo"
	"github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/rerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const webhookCol = "webhook_sent"

var _ repo.Webhook = (*Webhook)(nil)

type Webhook struct {
	c *mongo.Collection
}

func NewWebhook(db *mongo.Database) *Webhook {
	return &Webhook{
		c: db.Collection(webhookCol),
	}
}

func (w *Webhook) InitIndex(ctx context.Context) error {
	indexes, err := w.c.Indexes().List(ctx)
	if err != nil {
		return rerror.ErrInternalBy(err)
	}

	for indexes.Next(ctx) {
		d := struct {
			Name string `bson:"name"`
		}{}
		if err := indexes.Decode(&d); err != nil {
			return rerror.ErrInternalBy(err)
		}
		if d.Name == "event" {
			return nil
		}
	}

	_, err = w.c.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "event", Value: 1},
		},
		Options: options.Index().SetName("event").SetExpireAfterSeconds(int32((time.Hour * 24).Seconds())),
	})
	if err != nil {
		return rerror.ErrInternalBy(err)
	}

	log.Infof("worker mongo: webhook index created")
	return nil
}

func (w *Webhook) Get(ctx context.Context, eventID string) (bool, error) {
	res := w.c.FindOne(ctx, bson.M{
		"event": eventID,
	}, nil)
	var d bson.M
	if err := res.Decode(&d); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) || errors.Is(err, mongo.ErrNilDocument) {
			return false, nil
		}
		return false, rerror.ErrInternalBy(err)
	}
	return true, nil
}

func (w *Webhook) GetAndSet(ctx context.Context, eventID string) (bool, error) {
	res := w.c.FindOneAndUpdate(ctx, bson.M{
		"event": eventID,
	}, bson.M{
		"$set": bson.M{
			"event": eventID,
		},
	}, options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.Before))
	var d bson.M
	if err := res.Decode(&d); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) || errors.Is(err, mongo.ErrNilDocument) {
			return false, nil
		}
		return false, rerror.ErrInternalBy(err)
	}
	return true, nil
}

func (w *Webhook) Delete(ctx context.Context, eventID string) error {
	_, err := w.c.DeleteOne(ctx, bson.M{
		"event": eventID,
	})
	if err != nil {
		return rerror.ErrInternalBy(err)
	}
	return nil
}
