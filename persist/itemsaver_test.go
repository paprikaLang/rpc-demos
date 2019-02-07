package persist

import (
	"context"
	"encoding/json"
	"go-spider/model"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Url:    "https://laravelcollections.com/go/532",
		Domain: "github.com",
		Title:  "Laravel v5.7.22 Released &nbsp;",
	}
	id, err := save(expected)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	res, err := client.Get().Index("laravel_collection").Type("posts").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", res)
	t.Logf("%s", string(*res.Source)) //type *json.RawMessage
	var profileModel model.Profile
	err = json.Unmarshal(*res.Source, &profileModel)
	if err != nil {
		panic(err)
	}
	if profileModel != expected {
		t.Errorf("got %v; expected %v", profileModel, expected)
	}

}
