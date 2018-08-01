package parser

import (
	"io/ioutil"
	"testing"

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

	result := ParseProfile(contents, expectedProfile.Name)
	for _, val := range result.Items {
		if v, ok := val.(model.Profile); ok {

			if v.Name != expectedProfile.Name {
				t.Errorf("Expected name %s, but result name %s", expectedProfile.Name, v.Name)
			}

			if v.Age != expectedProfile.Age {
				t.Errorf("Expected age %d, but result age %d", expectedProfile.Age, v.Age)
			}
			if v.Gender != expectedProfile.Gender {
				t.Errorf("Expected gender %s, but result gender %s", expectedProfile.Gender, v.Gender)
			}
			if v.Height != expectedProfile.Height {
				t.Errorf("Expected height %d, but result height %d", expectedProfile.Height, v.Height)
			}
			if v.Weight != expectedProfile.Weight {
				t.Errorf("Expected weight %d, but result weight %d", expectedProfile.Weight, v.Weight)
			}
			if v.Income != expectedProfile.Income {
				t.Errorf("Expected income %s, but result income %s", expectedProfile.Income, v.Income)
			}
			if v.Marriage != expectedProfile.Marriage {
				t.Errorf("Expected marriage %s, but result marriage %s", expectedProfile.Marriage, v.Marriage)
			}
		} else {
			t.Errorf("Error Type")
		}
	}

}
