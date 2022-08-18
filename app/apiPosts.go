package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Post struct {
	ID       int
	Post     string
	Unixtime int
}

func init() {
	fmt.Println("Initialized API - posts")
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data, error := app.sql.Query("SELECT * FROM imageboard_db ORDER BY ID DESC")
		if error != nil {
			fmt.Println(error)
		}
		postsArray := []Post{}
		for data.Next() {
			var (
				id       int
				post     string
				unixtime int
			)
			if err := data.Scan(&id, &post, &unixtime); err != nil {
				panic(err)
			}
			InputPost := Post{ID: id, Post: post, Unixtime: unixtime}
			postsArray = append(postsArray, InputPost)
			if err := data.Err(); err != nil {
				panic(err)
			}
		}
		type dataOutput struct {
			Message string
			Error   int
			Posts   []Post
		}
		output := &dataOutput{Message: "SUCCES", Error: 0, Posts: postsArray}
		jsonData, err := json.Marshal(output)
		if err != nil {
			fmt.Printf("could not marshal json: %s\n", err)
			return
		}
		w.Write([]byte(jsonData))
		return
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		URLQuery := `INSERT INTO imageboard_db (Post, UnixTime) VALUES($1,$2)`
		unixtime := strconv.FormatInt(time.Now().Unix(), 10)
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields() // catch unwanted fields

		// anonymous struct type: handy for one-time use
		p := struct {
			PostData *string `json:"PostData"` // pointer so we can test for field absence
		}{}
		err := d.Decode(&p)
		if err != nil {
			// bad JSON or unrecognized json field
			panic(err)
			fmt.Println(err)
			//http.Error(r, err.Error(), http.StatusBadRequest)
			return
		}
		_, error := app.sql.Exec(URLQuery, *p.PostData, unixtime)
		if error != nil {
			fmt.Println(error)
			w.Write([]byte("{'message': 'Error'}"))
			return
		}
		w.Write([]byte("{'message': 'OK'}"))
		return
		// }
		w.Write([]byte("{'message': 'Error'}"))
		return
	}

	http.Error(w, "Method not allowed", 405)
}
