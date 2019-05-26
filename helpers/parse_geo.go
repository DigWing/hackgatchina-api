package helpers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type geoResp struct {
    Response struct{
        GeoObjectCollection struct {
            FeatureMember []struct {
                GeoObject struct {
                    MetaDataProperty struct {
                        GeocoderMetaData struct {
                            Text string `json:"text"`
                        }
                    } `json:"metaDataProperty"`
                }
            } `json:"featureMember"`
        }
    } `json:"response"`
}

func ParseGeo(lat, lng float64) string {
    resp, _ := http.Get(fmt.Sprintf("https://geocode-maps.yandex.ru/1.x/?apikey=6dd24157-ee18-4702-87df-f50c73f40e0d&format=json&geocode=%f,%f", lng, lat))
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    var geo = geoResp{}
    _ = json.Unmarshal(body, &geo)
    return geo.Response.GeoObjectCollection.FeatureMember[0].GeoObject.MetaDataProperty.GeocoderMetaData.Text
}
