package main

import (
	"html/template"
	"log"
	"net/http"
"fmt"
 "github.com/dgraph-io/badger/v2"
)

type NewsAggPage struct {
	Fname    string
	News     string
	Lastname string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Fname: "gagan", News: "some news", Lastname: "singh"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}
func data() {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("answer"), []byte("42"))
		return err
	})
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("answer"))
		handle(err)
}
err := item.Value(func(val []byte) error {
    // This func with val would only be called if item.Value encounters no error.

    // Accessing val here is valid.
	fmt.Printf("The answer is: %s\n", val)
}
}

func main() {

	http.HandleFunc("/", newsAggHandler)
	http.HandleFunc("/d", data)

	http.ListenAndServe(":8080", nil)
}
