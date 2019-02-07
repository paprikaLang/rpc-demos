package model

import (
	"encoding/json"
)

// Profile data
type Profile struct {
	Title  string
	Domain string
	Topic  string
}

// MapToProfile type the map
func MapToProfile(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o) //o map
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile) // 重新指定给profile
	return profile, err

}
