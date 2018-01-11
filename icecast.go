package main

import (
	"encoding/xml"
	"reflect"
)

// IcecastAdapter adapts into icecast data, to get the source info.
type IcecastAdapter struct {}

// IcecastStat contains an slice of sources.
type IcecastStat struct {
	Stat struct {
		Sources []IcecastSource `xml:"source"`
	} `xml:"icestats"`
}

// IcecastSource is source data for an icecast stream.
type IcecastSource struct {
	Listeners string `xml:"listeners"`
	CurrentlyPlaying string `xml:"yp_currently_playing"`
}

// Compare compares prev against body.
func (ia IcecastAdapter) Compare(prev, body []byte) bool {
	var p IcecastStat
	var b IcecastStat

	xml.Unmarshal(prev, &p)
	xml.Unmarshal(body, &b)

	return reflect.DeepEqual(p, b)
}