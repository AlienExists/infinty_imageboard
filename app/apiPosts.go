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
	Unixtime string
}

func init() {
	fmt.Println("Initialized API - posts")
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, error := app.sql.Query("SELECT * FROM imageboard_db")
		if error != nil {
			fmt.Println(error)
		}
		postsArray := []Post{}
		for data.Next() {
			var (
				id       int
				post     string
				unixtime string
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
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		for key, value := range r.Form {
			if key == "post" {
				URLQuery := `INSERT INTO imageboard_db (post, unixtime) VALUES($1,$2)`
				unixtime := strconv.FormatInt(time.Now().Unix(), 10)
				_, error := app.sql.Exec(URLQuery, value[0], unixtime)
				if error != nil {
					fmt.Println(error)
					w.Write([]byte("{'message': 'Error'}"))
					return
				}
				w.Write([]byte("{'message': 'OK'}"))
				return
			}
		}
		w.Write([]byte(r.URL.Query().Get("post")))
		return
	}

	http.Error(w, "Method not allowed", 405)
}
