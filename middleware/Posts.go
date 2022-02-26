package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/pavanilla/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func Init() {
	createDBInstance()
}

func createDBInstance() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database("appointy").Collection("posts")
}

// inserting a single post in golang

func CreatePost(w http.ResponseWriter, r *http.Request) {
	p := models.Post{}

	// decoding the request body into the struct.If there is a error
	// respond to the client with a error message and a 400 status code
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = Collection.InsertOne(context.TODO(), p)
	if err != nil {
		panic(err)
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// edit the posts

func EditPost(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("User_id")

	UserID, err := primitive.ObjectIDFromHex(user_id)

	if err != nil {
		panic(err)
	}

	// Declare an _id filter to get a specific MongoDB document
	filter := bson.M{"_id": bson.M{"$eq": UserID}}
	// Declare a filter that will change a Post Body   to `updated post `
	update := bson.M{"$set": bson.M{"post": "update post "}}
	_, err = Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

//list all posts
func Allposts(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{}} //bson.D{{}} specifies 'all documents'
	posts := []models.Post{}
	//we are going to get options frm mongo/options
	findOptions := options.Find()
	// Sort by `posted` field chronologically
	findOptions.SetSort(bson.D{{"posted", 1}})
	result, err := Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		panic(err)
	}
	for result.Next(context.TODO()) {
		var ele models.Post
		err = result.Decode(&ele)
		if err != nil {
			panic(err)
		}
		posts = append(posts, ele)
	}
	result.Close(context.TODO())
	if len(posts) == 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// delete all posts

func DeletePosts(w http.ResponseWriter, r *http.Request) {
	delete_id := r.URL.Query().Get("delete_id")
	DeleteID, err := primitive.ObjectIDFromHex(delete_id)
	if err != nil {
		panic(err)
	}
	// Declare an _id filter to get a specific MongoDB document
	filter := bson.M{"_id": bson.M{"$eq": DeleteID}}
	_, err = Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Please check the userId ")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}

// method to print error output for http respon
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// method to print output for http respon
// parameter  [w (Http.RestponWriter), http.statuscode, payload/data/msg]
// payload is data credential which will be trans to other part
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
