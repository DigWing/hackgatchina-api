package service

import (
    "context"
    "github.com/DigWing/hackaton/db"
    "github.com/DigWing/hackaton/models"
    "github.com/DigWing/hackaton/schemas"
    "github.com/mongodb/mongo-go-driver/bson"
    "github.com/mongodb/mongo-go-driver/bson/primitive"
    "time"
)

func SaveFeedback(f schemas.Feedback, parsedGeo string, tags []string) (*string, error) {
    coll := db.MongoDB.Collection("feedback")

    newFeedback := models.Feedback{
        CreatedAt: time.Now(),
        Image: f.Image,
        Geo: parsedGeo,
        Lat: f.Ltd,
        Lng: f.Lng,
        User: f.User,
        Tags: tags,
        Rating: 0,
        Comments: []string{},
    }

    resFeedback, err := coll.InsertOne(
        context.Background(),
        newFeedback,
    )

    if err != nil {
        return nil, err
    }

    id := resFeedback.InsertedID.(primitive.ObjectID).Hex()
    return &id, nil
}

func UpdateFeedbackTags(feedbackId string, tags []string) error {
    coll := db.MongoDB.Collection("feedback")
    fid, _ := primitive.ObjectIDFromHex(feedbackId)

    _, err := coll.UpdateOne(
        context.Background(),
        bson.D{
            {"_id", fid},
        },
        bson.D{
            {"$addToSet", bson.D{
                {"tags", bson.D{
                    { "$each", tags },
                }},
            }},
        },
    )

    return err
}

func SaveComment(feedbackId, message string) error {
    coll := db.MongoDB.Collection("feedback")
    fid, _ := primitive.ObjectIDFromHex(feedbackId)

    _, err := coll.UpdateOne(
        context.Background(),
        bson.D{
            {"_id", fid},
        },
        bson.D{
            {"$push", bson.D{
                {"comments", message},
            }},
        },
    )

    return err
}

func GetFeedbacks() []schemas.RespFeedback {
    coll := db.MongoDB.Collection("feedback")

    var feedbacks = []schemas.RespFeedback{}

    cur, _ := coll.Find(context.Background(), bson.D{})
    for cur.Next(context.Background()) {
        var feedback models.Feedback
        _ = cur.Decode(&feedback)
        feedbacks = append(feedbacks, schemas.RespFeedback{
            Id: feedback.Id.Hex(),
            CreatedAt: feedback.CreatedAt,
            Rating: feedback.Rating,
            Image: feedback.Image,
            Geo: feedback.Geo,
            User: feedback.User,
            Lat: feedback.Lat,
            Lng: feedback.Lng,
            Comments: feedback.Comments,
            Tags: feedback.Tags,
        })
    }

    return feedbacks
}
