package main

type Adapter interface {
	Compare(prev, body []byte) bool
}