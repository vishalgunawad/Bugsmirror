package main

import (
	"context"

	"net/http"

	"encoding/json"

	_ "log"

	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUser = "mongo_db_admin"

	dbPass = "EXAMPLE_PASSWORD"

	dbName = "shop_db"
)

func main() {

	http.HandleFunc("/api/v1/products", requestHandler)

	http.ListenAndServe(":8080", nil)

}

func requestHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbUser+":"+dbPass+"@localhost:27017"))

	if err != nil {

		fmt.Println(err.Error())

	}

	collection := client.Database(dbName).Collection("products")

	data := map[string]interface{}{}

	err = json.NewDecoder(req.Body).Decode(&data)

	if err != nil {

		fmt.Println(err.Error())

	}

	switch req.Method {

	case "POST":

		response, err = createRecord(collection, ctx, data)

	case "GET":

		response, err = getRecords(collection, ctx)

	case "PUT":

		response, err = updateRecord(collection, ctx, data)

	case "DELETE":

		response, err = deleteRecord(collection, ctx, data)

	}

	if err != nil {

		response = map[string]interface{}{"error": err.Error()}

	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {

		fmt.Println(err.Error())

	}

}
func deleteRecord(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {

	_, err := collection.DeleteOne(ctx, bson.M{"product_id": data["product_id"]})

	if err != nil {

		return nil, err

	}

	res := map[string]interface{}{

		"data": "Document deleted successfully.",
	}

	return res, nil

}

func getRecords(collection *mongo.Collection, ctx context.Context) (map[string]interface{}, error) {

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {

		return nil, err

	}

	defer cur.Close(ctx)

	var products []bson.M

	for cur.Next(ctx) {

		var product bson.M

		if err = cur.Decode(&product); err != nil {

			return nil, err

		}

		products = append(products, product)

	}

	res := map[string]interface{}{}

	res = map[string]interface{}{

		"data": products,
	}

	return res, nil

}

func createRecord(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {

	req, err := collection.InsertOne(ctx, data)

	if err != nil {

		return nil, err

	}

	insertedId := req.InsertedID

	res := map[string]interface{}{

		"data": map[string]interface{}{

			"insertedId": insertedId,
		},
	}

	return res, nil

}

func updateRecord(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {

	filter := bson.M{"product_id": data["product_id"]}

	fields := bson.M{"$set": data}

	_, err := collection.UpdateOne(ctx, filter, fields)

	if err != nil {

		return nil, err

	}

	res := map[string]interface{}{

		"data": "Document updated successfully.",
	}

	return res, nil

}
