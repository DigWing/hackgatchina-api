package controller

import (
    "bytes"
    "github.com/DigWing/hackaton/helpers"
    "github.com/DigWing/hackaton/schemas"
    "github.com/DigWing/hackaton/service"
    "github.com/labstack/echo"
    "net/http"
    "os/exec"
    "regexp"
    "strings"
)

type PostFeedbackResp struct {
    Id string `json:"id"`
}

func PostFeedback(c echo.Context) error {
    f := new(schemas.Feedback)
    if err := c.Bind(f); err != nil {
        return c.JSON(http.StatusBadRequest, nil)
    }

    //parsing image tags
    t := helpers.ParseImageTags(f.Image)
    tags := []string{}
    for tag := range t {
        tags = append(tags, tag)
    }

    //parsing geo by coords
    geo := ""
    if f.Lng != nil && f.Ltd != nil {
        geo = helpers.ParseGeo(*f.Ltd, *f.Lng)
    }
    id, _ := service.SaveFeedback(*f, geo, tags)

    return c.JSON(http.StatusOK, PostFeedbackResp{
        Id: *id,
    })
}

func PostFeedbackComment(c echo.Context) error {
    feedbackId := c.Param("feedbackId")
    comment := new(schemas.Comment)
    if err := c.Bind(comment); err != nil {
        return c.JSON(http.StatusBadRequest, nil)
    }

    //words processing
    spelledMessage := helpers.SpellerSentence(comment.Message)
    cmd := exec.Command("./mystem")
    cmd.Stdin = strings.NewReader(spelledMessage)

    var out bytes.Buffer

    cmd.Stdout = &out
    _ = cmd.Run()

    r, _ := regexp.Compile("{.+?}")
    words := r.FindAllString(string(out.Bytes()), -1)

    reg := regexp.MustCompile("{|}")
    for i, word := range words {
        words[i] = strings.Split(reg.ReplaceAllString(word, ""), "|")[0]
    }
    //words - token & lemma

    var tags = map[string]bool{}
    for _, word := range words {
        t := helpers.ParseWordTags(word)
        for tag := range t {
            tags[tag] = true
        }
    }

    arrayTags := []string{}
    for tag := range tags {
        arrayTags = append(arrayTags, tag)
    }

    _ = service.UpdateFeedbackTags(feedbackId, arrayTags)
    _ = service.SaveComment(feedbackId, spelledMessage)

    return c.JSON(http.StatusOK, PostFeedbackResp{
        Id: feedbackId,
    })
}

func GetFeedbacks(c echo.Context) error {
    return c.JSON(http.StatusOK, service.GetFeedbacks())
}