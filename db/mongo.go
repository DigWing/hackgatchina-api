package db

import (
    "context"
    "github.com/joho/godotenv"
    "github.com/mongodb/mongo-go-driver/mongo"
    "os"
    "time"
)

var MongoDB *mongo.Database

func init() {
    _ = godotenv.Load()

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    client, _ := mongo.Connect(ctx, os.Getenv("MONGO_HOST"))

    MongoDB = client.Database(os.Getenv("MONGO_DB"))
}