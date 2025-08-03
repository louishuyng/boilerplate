package db

import (
	"context"
	"golang-mongo/internal/application/core/domain"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Adapter struct {
	db *mongo.Database
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dataSourceUrl))

	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
		return &Adapter{}, err

	}

	db := client.Database("dev")

	return &Adapter{
		db: db,
	}, nil
}

func (a Adapter) Get(ctx context.Context, id string) (domain.Payment, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Payment{}, err
	}

	collection := a.db.Collection("payments")
	var payment domain.Payment
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&payment)
	if err != nil {
		return domain.Payment{}, err
	}

	return payment, nil
}

func (a Adapter) Save(ctx context.Context, payment *domain.Payment) error {
	collection := a.db.Collection("payments")

	if payment.ID == 0 {
		result, err := collection.InsertOne(ctx, payment)
		if err != nil {
			return err
		}

		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			payment.ID = int64(oid.Timestamp().Unix())
		}
	} else {
		_, err := collection.ReplaceOne(ctx, bson.M{"id": payment.ID}, payment)
		if err != nil {
			return err
		}
	}

	return nil
}

