package controller

import (
	"context"
	"log"
	model "subscription/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Pranjal:Pranjal%40123@cluster0.sc7ucqz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "subscriptionBucket"
const colName = "UserEmail"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(colName)
}
func CheckIfAlreadyExit(email string) bool {
	var filter = bson.M{"email": email}
	var result bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	return err != mongo.ErrNoDocuments
}
func InsertIntoDatabase(DataToInsert model.DataPaylod) *mongo.InsertOneResult {
	result, err := collection.InsertOne(context.Background(), DataToInsert)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
func GetAllData() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var bsonData []primitive.M
	for cursor.Next(context.Background()) {
		var insideData bson.M
		err := cursor.Decode(&insideData)
		if err != nil {
			log.Fatal(err)
		}
		bsonData = append(bsonData, insideData)
	}
	defer cursor.Close(context.Background())
	return bsonData
}
