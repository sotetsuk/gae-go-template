package main

import (
	"fmt"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/", handle)
}

type Users struct {
	Name     string
	TImestamp time.Time
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	u1 := Users{
		Name:     "sotetsuk",
		TImestamp: time.Now(),
	}

	key := datastore.NewKey(ctx, "user", "0001", 0, nil)
	_, err := datastore.Put(ctx, key, &u1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var u2 Users
	if err = datastore.Get(ctx, key, &u2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "id: %s\n", key.String())
	fmt.Fprintf(w, "stringID: %s\n", key.StringID())
	fmt.Fprintf(w, "Stored and retrieved data: %q", u2.Name, u2.TImestamp.String())
}
