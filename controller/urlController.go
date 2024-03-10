package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ARMAAN199/practiceURL/modal"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

const connectionString = "mongodb+srv://testing:qwerty1239@cluster0.z0vkxn2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "test"
const collectionName = "urls"

var urlcollection *mongo.Collection

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}

	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()

	urlcollection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Connection with DB established")
}

func insertUrl(url modal.Url) (*mongo.InsertOneResult, error) {
	res, err := urlcollection.InsertOne(context.Background(), url)
	fmt.Println("here 4")
	if err != nil {
		return nil, err
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
	return res, nil
}

func getOldUrl(shortVersion string) (modal.Url, error) {
	filter := bson.D{{"shortened", shortVersion}}
	var result modal.Url

	fmt.Println("here 3")
	err := urlcollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Printf("Not Found Url mapped to %s", shortVersion)
		}
		panic(err)
	}
	return result, nil
}

func GetActualUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "GET")

	Qparams := mux.Vars(r)
	oldUrl, err := getOldUrl(Qparams["short"])
	fmt.Println("here 2")
	if err != nil {
		fmt.Printf("Didn't find anything")
	}

	_ = json.NewEncoder(w).Encode(oldUrl)
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "POST")

	fmt.Println("here 1")
	var newUrl modal.Url
	err := json.NewDecoder(r.Body).Decode(&newUrl)

	newUrl.Shortened = uuid.New().String()

	if err != nil {
		fmt.Printf("Cannot Understand This Url")
	}

	fmt.Printf("Request %+v", newUrl)

	res, err := insertUrl(newUrl)
	if err != nil {
		fmt.Printf("Cannot insert URl")
	}
	fmt.Printf("inserted Url %+v", res)
	_ = json.NewEncoder(w).Encode("Inserted Movie")
}
