package main

import (
	"encoding/json"
	"reflect"
)

// ShoutcastAdapter adapts into icecast data, to get the source info.
type ShoutcastAdapter struct {}

// ShoutcastStat contains an slice of streams.
type ShoutcastStat struct {
	Streams []ShoutcastStream `json:"streams"`
}

// IcecastStream is source data for an icecast stream.
type ShoutcastStream struct {
	Listeners int `json:"currentlisteners"`
	CurrentlyPlaying string `json:"songtitle"`
}

// Compare compares prev against body.
func (ia ShoutcastAdapter) Compare(prev, body []byte) bool {
	var p ShoutcastStat
	var b ShoutcastStat

	json.Unmarshal(prev, &p)
	json.Unmarshal(body, &b)

	return reflect.DeepEqual(p, b)
}