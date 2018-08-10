package persist

import (
	"context"
	"encoding/json"
	"testing"

	"gopkg.in/olivere/elastic.v5"

	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler/model"
)

func TestSave(t *testing.T) {

	// 测试的样例尽量就写在文件不要单独统一写，便于调试查错
	profile := model.Profile{
		Name:     "如果芸许",
		Gender:   "女",
		Age:      35,
		Height:   155,
		Income:   "3000元以下",
		Marriage: "离异",
	}

	expectedItem := engine.Item{
		Url:     "http://album.zhenai.com/u/108848769",
		Id:      "108848769",
		Type:    "zhenai",
		Payload: profile,
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	err = Save(client, expectedItem, "profile_test")
	if err != nil {
		panic(err)
	}

	// TODO: try to start up elastic search
	// here using docker go client
	// index和type里的字符串常量也不用拎出去单独写，便于查找拼写错误
	resp, err := client.Get().
		Index("profile_test").
		Type(expectedItem.Type).
		Id(expectedItem.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", resp)
	t.Logf("%s", *resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual) // 无法递归unmarshal内部json形式的数据
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expectedItem {
		t.Errorf("got %v; expected %v", actual, expectedItem)
	}

}
