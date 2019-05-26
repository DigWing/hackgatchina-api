package helpers

import (
    "github.com/IBM/go-sdk-core/core"
    "github.com/watson-developer-cloud/go-sdk/visualrecognitionv3"
)

var visualRecognition, _ = visualrecognitionv3.NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
    Version:   "2018-03-19",
    IAMApiKey: "ogZZOO6MP4MbBEpvJMC9wApgpUNWhcdg76DxV5CedGuz",
})

var imageTags = map[string][]string{
    "garbage": []string{"garbage heap","garbage","bottle green color","receptacle","trash can","litter basket","litter"},
    "graffiti": []string{"graffiti"},
    "road": []string{"road","curved road","bypath","Highway","arterial road","autostrada (highway)","highway","autobahn (highway)","carriageway","street","'pothole (road)", "hole", "sinkhole"},
    "building": []string{"building","government building","balcony (of building)","tenement house","apartment building","cornice (building)","palace","shouldered arch","arch","guildhall (ornate building)","facade"},
}

var wordsTags = map[string][]string{
    "garbage": []string{"мусор","пакет","мусорный","урна","ведро","бычок","бутылка","окурок","семечко","отходы","грязь","грязные","говно","фекалии","cрань","хлам","плесень","отбросы","пыль","труха","помойка"},
    "graffiti": []string{"граффити","вандал","разрисовывать"},
    "road": []string{"дорога","яма","грунтовка","светофор","пробка","перекресток","шоссе","трасса","мост","разметка","переход","затор","улица","грунт"},
    "building": []string{"здание","стена","забросить","постройка","стройка","жилище","жкх","ресторан","магазин","церковь","храм","отопление","башня","подъезд"},
}

func ParseImageTags(url string) (tags map[string]bool) {
    tags = map[string]bool{}
    response, responseErr := visualRecognition.Classify(
        &visualrecognitionv3.ClassifyOptions{
            URL: core.StringPtr(url),
        },
    )

    if responseErr != nil {
        panic(responseErr)
    }
    result := visualRecognition.GetClassifyResult(response)

    for _, class := range result.Images[0].Classifiers[0].Classes {
        for tag, anchors := range imageTags {
            for _, anchor := range anchors {
                if anchor == *class.ClassName {
                    tags[tag] = true
                }
            }
        }
    }

    return
}

func ParseWordTags(word string) (tags map[string]bool) {
    tags = map[string]bool{}
    for tag, anchors := range wordsTags {
        for _, anchor := range anchors {
            if anchor == word {
                tags[tag] = true
                break
            }
        }
    }
    return
}