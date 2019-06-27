package models

import (
    "App/database"
    "App/utils"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

const USERS = "members"

type User struct {
    ID       *primitive.ObjectID `bson:"_id,omitempty"`
    Name     string              `json:"name",bson:"name"`
    Age      int                 `json:"age",bson:"age"`
    Company  string              `json:"company",bson:"company"`
    IsActive bool                `json:"is_active",bson:"isActive"`
}

type UserPatch struct {
    Age *int `json:"age"`
}

func (User *User) ApplyPatch(patch UserPatch) {
    if patch.Age != nil {
        User.Age = *patch.Age
    }
}

func (user *User) Load(id string) error{

    userDB,err := UserRepository{}.FindById(id)
    *user = userDB
    return err
}

func (user *User) Drop() error {
    _, err := UserRepository{}.Collection().DeleteOne(database.Context, bson.M{"_id": user.ID})
    if err == nil {
        user.ID = nil
    }
    return err
}

func (user *User) Store() {
    collection := UserRepository{}.Collection()
    if user.ID == nil {
        fmt.Println("abcdef")
        response, err := collection.InsertOne(database.Context, user)
        utils.PanicOnError(err)
        id := response.InsertedID.(primitive.ObjectID)
        user.ID = &id
        return
    }
    _, err := collection.UpdateOne(database.Context, bson.M{"_id": *user.ID},
        bson.M{"$set": user})
    utils.PanicOnError(err)
}

type UserRepository struct {
}

func (UserRepository) Collection() *mongo.Collection {
    return database.Collection(USERS)
}

func (UserRepository) FindAll() []User {

    cur, _ := database.Collection(USERS).Find(database.Context, bson.D{})
    users := []User{}
    if cur != nil {
        for cur.Next(database.Context) {
            user := User{}
            cur.Decode(&user)
            users = append(users, user)
        }
        cur.Close(database.Context)
    }
    return users

}

func (UserRepository) FindById(id string) (User, error) {
    objectId, _ := primitive.ObjectIDFromHex(id)
    filter := bson.M{"_id": objectId}
    var user User
    err := UserRepository{}.Collection().FindOne(database.Context, filter).Decode(&user)
    return user, err
}
