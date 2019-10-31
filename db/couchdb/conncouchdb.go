package main

import (
	"fmt"
	"github.com/joelnb/sofa"
	"time"
)



func main() {
	conn, err := sofa.NewConnection("http://localhost:5984", 10*time.Second, sofa.NullAuthenticator())
	if err != nil {
		panic(err)
	}
	db := conn.Database("test")
	doc := &struct {
		sofa.DocumentMetadata
		Name string `json:"name"`
		Type string `json:"type"`
	}{
		DocumentMetadata: sofa.DocumentMetadata{
			ID: "test",
		},
		Name: "apple",
		Type: "fruit",
	}
	rev, err := db.Put(doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(rev)


}
