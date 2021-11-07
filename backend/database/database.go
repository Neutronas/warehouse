package database

import (
	"context"
	"log"
	"time"

	"github.com/neutronas/warehouse/backend/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LocationType struct {
	name string
	lat float32
	lng float32
}
type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		client: client,
	}
}

func (db* DB) CreateWarehouse(input *model.NewWarehouse) *model.Warehouse {
	collection := db.client.Database("db").Collection("warehouses")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Warehouse {
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
		Name: input.Name,
		Description: input.Description,
		Location: &model.Location {
			Name: input.LocationName,
			Lat: input.Lat,
			Lng: input.Lng,
		},
	}
}

func (db *DB) FindWarehouseByID(ID string) *model.Warehouse {
	ObjectID, err := primitive.ObjectIDFromHex(ID)

	collection := db.client.Database("db").Collection("warehouses")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		log.Fatal(err)
	}
	warehouse := model.Warehouse {}
	res.Decode(&warehouse)
	return &warehouse	
}

func (db *DB) FindAllWarehouses() []*model.Warehouse {
	collection := db.client.Database("db").Collection("warehouses")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var warehouses []*model.Warehouse
	for cur.Next(ctx) {
		var warehouse *model.Warehouse
		err := cur.Decode(&warehouse)
		if err != nil {
			log.Fatal(err)
		}
		warehouses = append(warehouses, warehouse)
	}
	return warehouses
}
