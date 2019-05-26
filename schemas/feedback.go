package schemas

import (
    "time"
)

type Feedback struct {
    Image string `json:"image"`
    Ltd   *float64 `json:"ltd"`
    Lng   *float64 `json:"lng"`
    User  string `json:"user"`
}

type Comment struct {
    Message    string   `json:"message"`
}

type RespFeedback struct {
    Id          string              `json:"id"`
    CreatedAt   time.Time           `json:"createdAt"`
    Rating      int                 `json:"rating"`
    Image       string              `json:"image"`
    Geo         string              `json:"geo"`
    User        string              `json:"user"`
    Lat         *float64            `json:"lat"`
    Lng         *float64            `json:"lng"`
    Comments    []string            `json:"comments"`
    Tags        []string            `json:"tags"`
}
