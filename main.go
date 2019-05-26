package main

import (
    "github.com/DigWing/hackaton/controller"
    _ "github.com/DigWing/hackaton/db"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "net/http"
)

func main() {
    e := echo.New()

    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
    }))

    //helpers.ParseGeo(37.611347,55.760241)
    //tags := helpers.ParseTags("https://cs.pikabu.ru/post_img/big/2013/06/19/7/1371635872_974328138.jpg")
    e.POST("/feedback", controller.PostFeedback)
    e.POST("/feedback/:feedbackId/comment", controller.PostFeedbackComment)
    e.GET("/feedbacks", controller.GetFeedbacks)

    _ = e.Start(":8080")
}