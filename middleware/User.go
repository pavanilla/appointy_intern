package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pavanilla/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CollectionUser *mongo.Collection

func InitUser() {
	createUserInstance()
}

func createUserInstance() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	CollectionUser = client.Database("appointy").Collection("Users")
}

//Create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	filter := bson.D{{"username", u.Username}}
	total, err := Collection.Find(context.TODO(), filter)
	if total != nil {
		respondWithError(w, http.StatusBadRequest, "User already exists with the username ")
		return
	}
	// In this way we can elaimate creating the duplicate usernames
	result, err := CollectionUser.InsertOne(context.TODO(), u)
	fmt.Println(result)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Please check the user Id ")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}

//Edit user details
// here getting the data in the form of url encoded data
// since the username is the uniue constraint here we can use it as a filter and perform the update step on the document

func EditUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	filter := bson.D{{"username", username}}
	replacement := bson.D{{"name", "pavan_illa"}, {"email", "pavanilla457@gmail.com"}, {"username", "pavan"}, {"password", "hashed"}, {"dob", "03-12-1999"}}
	_, err := CollectionUser.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		panic(err)
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

//Get User
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := models.User{}
	user_id := r.URL.Query().Get("delete_id")
	UserID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		panic(err)
	}
	// Declare an _id filter to get a specific MongoDB document
	filter := bson.M{"_id": bson.M{"$eq": UserID}}
	err = CollectionUser.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Please check the user Id ")
		return
	}
	json.NewEncoder(w).Encode(user)
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

//Search User
// in the mongoriver find function there is a third parameter
func SearchUser(w http.ResponseWriter, r *http.Request) {

	//find records
	//pass these options to the Find method
	findOptions := options.Find()
	//Define an array in which you can store the decoded documenents
	var users []models.User

	//Passing the bson.D{{}} as the filter matches  documents in the collection
	result, err := CollectionUser.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	//Finding multiple documents returns a cursor
	//Iterate through the cursor allows us to decode documents one at a time

	for result.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem models.User
		err := result.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, elem)

	}
	if len(users) == 0 {
		respondWithError(w, http.StatusBadRequest, "No users  found with the given data ")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

//Login Endpoint
// It was clearly mentioned in the user filed that use name is going to be uniue we can uery on that
func LoginEndPoint(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	user := models.User{}
	filter := bson.D{{"username", username}}
	err := CollectionUser.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(user)
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
