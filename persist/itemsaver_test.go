package persist

import (
	"context"
	"encoding/json"
	"go-spider/engine"
	"go-spider/model"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		URL:  "https://laravelcollections.com/go/612",
		Type: "collections",
		Id:   "612",
		Payload: model.Profile{
			Topic:  "News",
			Domain: "github.com",
			Title:  "Laravel v5.7.25 Released &nbsp;",
		},
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	index := "laravel_collections"
	err = Save(client, expected, index)
	if err != nil {
		panic(err)
	}

	res, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", res)
	t.Logf("%s", string(*res.Source)) //type *json.RawMessage
	var profileModel engine.Item
	json.Unmarshal(*res.Source, &profileModel)
	//got {map[Title:Laravel v5.7.22 Released &nbsp; Domain:github.com Topic:News]};
	//expected {{Laravel v5.7.22 Released &nbsp; github.com News}}
	profile, _ := model.MapToProfile(
		profileModel.Payload)
	profileModel.Payload = profile

	if profileModel != expected {
		t.Errorf("got %v; expected %v", profileModel, expected)
	}
}
