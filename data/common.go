package data

type node struct {
	Data interface{}
	Next *node
	Previous *node
}
