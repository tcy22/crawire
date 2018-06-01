package save

import (
	"testing"
	"crawier/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"crawier/engine"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url : "http://album.zhenai.com/u/109596518",
		Type: "zhenai",
		Id:   "2222222",
		Payload: model.Profile{
			Name: "tcy222222",
			Age:  "22222222222222",
			Marriage: "未婚2222222222222",
		},
	}

	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.163.128:9200"),elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	err = Save(client,expected)
	if err != nil {
		panic(err)
	}
	resp,err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s",resp.Source)

}