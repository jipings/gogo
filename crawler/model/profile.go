package model

import (
	"encoding/json"
)

type Profile struct {
	Name     string
	Gender   string
	Age      int
	Height   int
	Weight   int
	Income   string
	Marriage string
	// Education  string
	// Occupation string
	// HoKou      string
	// Xinzuo     string
	// House      string
	// Car        string
}

// 因为无法将从es拉过来的数据unmarshal成包含payload的json数据
func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)

	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
