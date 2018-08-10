package parser

import (
	"io/ioutil"
	"testing"

	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler/model"
)

func TestProfileParser(t *testing.T) {

	contents, err := ioutil.ReadFile("profile_contents.html")
	if err != nil {
		panic(err)
	}
	expectedProfile := model.Profile{
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
		Payload: expectedProfile,
	}
	expectedUrl := []string{
		"http://album.zhenai.com/u/1479932680",
		"http://album.zhenai.com/u/1772003674",
		"http://album.zhenai.com/u/1834530365",
	}

	result := ParseProfile(contents, expectedItem.Url, expectedProfile.Name)
	for _, v := range result.Items {
		if v.Url != expectedItem.Url {
			t.Errorf("Expected url %s, but actual url %s", expectedItem.Url, v.Url)
		}
		if v.Type != expectedItem.Type {
			t.Errorf("Expected type %s, but actual type %s", expectedItem.Type, v.Type)
		}

		if v.Id != expectedItem.Id {
			t.Errorf("Expected Id %s, but actual Id %s", expectedItem.Id, v.Id)
		}
		if v, ok := v.Payload.(model.Profile); ok {
			if v != expectedProfile {
				t.Errorf("Expected profile %v, but actul %v", expectedProfile, v)
			}
		} else {
			t.Errorf("Paylod is not type of model.Profile")
		}
	}

	if len(expectedUrl) != len(result.Requests) {
		t.Errorf("Expected guess url len %d, but acutal %d", len(expectedUrl), len(result.Requests))
	}

	for i, v := range result.Requests {
		if v.Url != expectedUrl[i] {
			t.Errorf("Expected Url %s, but actual %s", expectedUrl[i], v.Url)
		}
	}

}
