package models

import (
    "github.com/mongodb/mongo-go-driver/bson/primitive"
    "time"
)

type Feedback struct {
    Id          primitive.ObjectID  `bson:"_id,omitempty"`
    CreatedAt   time.Time           `bson:"createdAt"`
    Rating      int                 `bson:"rating"`
    Image       string              `bson:"image"`
    Geo         string              `bson:"geo"`
    User        string              `bson:"user"`
    Lat         *float64            `bson:"lat"`
    Lng         *float64            `bson:"lng"`
    Comments    []string            `bson:"comments"`
    Tags        []string            `bson:"tags"`
}
