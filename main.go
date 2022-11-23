package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Asdasd123*"
	dbname   = "postgres"
)

var db *sql.DB

func init() {
	var err error
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}
}

type Movie struct {
	ID, Isbn, Title string
}

func InsertMovie(data Movie) {
	fmt.Println(data.Title)
	result, err := db.Exec("Insert into golang.movies(id,isbn,title) values($1,$2,$3)", data.ID, data.Isbn, data.Title)
	if err != nil {
		log.Fatal(err)
	} else {
		rowsAffected, _ := result.RowsAffected()
		fmt.Printf("etkilenen kayıt sayısı: %d", rowsAffected)
	}
}

func UpdateMovie(data Movie) {
	result, err := db.Exec("Update golang.movies set isbn=$2,title=$3 where id=$1", data.ID, data.Isbn, data.Title)
	if err != nil {
		log.Fatal(err)
	} else {
		rowsAffected, _ := result.RowsAffected()
		fmt.Printf("etkilenen kayıt sayısı: %d", rowsAffected)
	}
}

func GetMovies() {
	rows, err := db.Query("Select * from golang.movies")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No records found!  ")
			return
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var movies []*Movie
	for rows.Next() {
		mov := &Movie{}
		err := rows.Scan(&mov.ID, &mov.Isbn, &mov.Title)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, mov)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, value := range movies {
		fmt.Printf("%s-  %s, %s", value.ID, value.Isbn, value.Title)
	}
}

func GetMovieById(id string) {
	var movie string

	err := db.QueryRow("Select title from golang.movies where id=$1", id).Scan(&movie)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No records found!")
	case err != nil:
		log.Fatal((err))
	default:
		fmt.Printf("Movieis %s", movie)

	}
}

func main() {
	movie := Movie{
		ID:    "3",
		Isbn:  "456",
		Title: "Not defteri",
	}

	InsertMovie(movie)
	x := Movie{
		ID:    "3",
		Isbn:  "456",
		Title: "Eot defteri",
	}
	UpdateMovie(x)

	GetMovies()
	GetMovieById("1")

}
