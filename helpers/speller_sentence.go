package helpers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "regexp"
    "strings"
)

type respSpeller struct {
    Word string `json:"word"`
    Pos int     `json:"pos"`
    Len int     `json:"len"`
    S []string `json:"s"`
}

func SpellerSentence(sentence string) string {
    //remove extra spaces
    //TODO move to helper
    sentence = strings.TrimSpace(sentence)
    space := regexp.MustCompile(`\s+`)
    sentence = space.ReplaceAllString(sentence, " ")

    s := strings.Split(sentence, "")
    fmt.Println(s[1:3])
    resp, _ := http.Get(fmt.Sprintf("https://speller.yandex.net/services/spellservice.json/checkText?text=%s", url.QueryEscape(sentence)))
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    var sp = []respSpeller{}
    _ = json.Unmarshal(body, &sp)

    if len(sp) == 0 {
        return sentence
    }

    //swap words in sentence
    newSentence := []string{}
    arrayOfSentenceChars := strings.Split(sentence, "")
    newSentence = append(newSentence, arrayOfSentenceChars[:sp[0].Pos]...)
    for idx, word := range sp {
        newSentence = append(newSentence, strings.Split(word.S[0], "")...)
        nextPos := 0
        if idx == len(sp) - 1 {
            nextPos = len(arrayOfSentenceChars)
        } else {
            nextPos = sp[idx+1].Pos
        }
        newSentence = append(newSentence, arrayOfSentenceChars[word.Pos+word.Len:nextPos]...)
    }

    return strings.Join(newSentence, "")
}