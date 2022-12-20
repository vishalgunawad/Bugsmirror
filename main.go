package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "log"
	"net/http"
)

const (
	dbUser = "dark"

	dbPass = "lordvishal"

	dbName = "BugsMirror"
)

func main() {

	http.HandleFunc("/api/v1/Users", requestHandler)

	http.ListenAndServe(":8080", nil)

}

func requestHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://dark:lordvishal@cluster0.pppaw.mongodb.net/BugsMirror"))

	if err != nil {

		fmt.Println(err.Error())

	}

	collection := client.Database(dbName).Collection("Users")

	data := map[string]interface{}{}

	err = json.NewDecoder(req.Body).Decode(&data)

	if err != nil {

		fmt.Println(err.Error())

	}

	switch req.Method {

	case "POST":

		response, err = Create_a_new_User(collection, ctx, data)

	case "GET":

		response, err = Get_All_Users(collection, ctx)

	case "PUT":

		response, err = Edit_User(collection, ctx, data)

	case "DELETE":

		response, err = Delete_a_User(collection, ctx, data)

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

func Create_a_new_User(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {

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

func Delete_a_User(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {

	_, err := collection.DeleteOne(ctx, bson.M{"id": data["id"]})

	if err != nil {

		return nil, err

	}

	res := map[string]interface{}{

		"data": "User deleted successfully.",
	}

	return res, nil

}

func Get_All_Users(collection *mongo.Collection, ctx context.Context) (map[string]interface{}, error) {

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {

		return nil, err

	}

	defer cur.Close(ctx)

	var Users []bson.M

	for cur.Next(ctx) {

		var product bson.M

		if err = cur.Decode(&product); err != nil {

			return nil, err

		}

		Users = append(Users, product)

	}

	res := map[string]interface{}{}

	res = map[string]interface{}{

		"data": Users,
	}

	return res, nil

}

func Edit_User(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {

	filter := bson.M{"id": data["id"]}

	fields := bson.M{"$set": data}

	_, err := collection.UpdateOne(ctx, filter, fields)

	if err != nil {

		return nil, err

	}

	res := map[string]interface{}{

		"data": "Users Edited successfully.",
	}

	return res, nil

}
