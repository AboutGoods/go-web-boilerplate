package database

import (
    "App/utils/log"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "net/url"
    "strings"
)

var client *mongo.Client
var database *mongo.Database
var Context context.Context

func NewConnection(connectionString string) {
    Context = context.Background()

    parsedUrl, err := url.Parse(connectionString)
    log.PanicOnError(err, "cannot parse mongodb connection url")

    databaseName := strings.Trim(parsedUrl.RequestURI(), "/")

    password, _ := parsedUrl.User.Password()
    credentials := options.Credential{
        Username: parsedUrl.User.Username(),
        Password: password,
    }

    opt := options.
        Client().
        ApplyURI(connectionString).
        SetAuth(credentials)
    client, err = mongo.Connect(Context, opt)
    log.PanicOnError(err, "cannot connect to mongodb")
    database = client.Database(databaseName)
}

func Collection(collectionName string) *mongo.Collection {
    return database.Collection(collectionName)
}


