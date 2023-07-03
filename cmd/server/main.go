package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	greetv1 "github.com/danny-personal/go-connect-sample/gen/greet/v1"

	"github.com/bufbuild/connect-go"
	"github.com/bufbuild/protovalidate-go"
	"github.com/danny-personal/go-connect-sample/gen/greet/v1/greetv1connect"
	_ "github.com/lib/pq"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type GreetServer struct {
	db *sql.DB
}

const (
	// .devcontainerの.env参照
	HOST     = "localhost"
	DATABASE = "postgres"
	USER     = "postgres"
	PASSWORD = "postgres"
	PORT     = "5432"
)

type User struct {
	ID   int    `json:"id"`
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	v, err := protovalidate.New()
	if err != nil {
		log.Println("failed to initialize validator:", err)
	}
	if err = v.Validate(req.Msg); err != nil {
		log.Println("validation failed:", err)
	} else {
		log.Println("validation succeeded")
	}
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	rows, err := s.db.Query("SELECT id, age, name FROM users")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Age, &u.Name)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(u)
		users = append(users, u)
	}
	fmt.Println(users)
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	var connStr string = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", HOST, PORT, USER, DATABASE, PASSWORD)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	greeter := &GreetServer{db: db}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
