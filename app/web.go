package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("Initialized WEB")
}

type application struct {
	sql *sql.DB
}

func startWeb() {
	db, err := openDB()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	app := &application{
		sql: db,
	}
	mux := http.NewServeMux()
	api_ver := "1"
	fileServer := http.FileServer(http.Dir("/app/app/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.Handle("/static", http.NotFoundHandler())

	mux.HandleFunc("/api/v"+api_ver+"/posts", app.posts)

	fileServerPage := http.FileServer(http.Dir("/app/app/web_pages"))
	mux.Handle("/", http.StripPrefix("/", fileServerPage))

	ServerAddress := ":" + os.Getenv("PORT")
	//ServerAddress := ":9999"
	fmt.Println("Server starts on :" + os.Getenv("PORT"))
	error := http.ListenAndServe(ServerAddress, mux)
	//mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	log.Fatal(error)
}
