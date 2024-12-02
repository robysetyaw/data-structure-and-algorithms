package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sda-learn/rest/handlers"
	"sda-learn/rest/repositories"
	"sda-learn/rest/services"
)

func Rest() {
	repo, err := repositories.NewUserRepository("sda-learn/rest/data")
	if err != nil {
		log.Fatal("Cannot load file")
	}
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)
	// Set up routes

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetUsers(w, r)
		case http.MethodPost:
			handler.InsertUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetUserByID(w, r)
		case http.MethodPut:
			handler.UpdateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	ID    int `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func RestV2() {

	user, err := restV2LoadData("./rest/data.json")
	if err != nil {
		log.Print(err)
	}

	http.HandleFunc("/users",func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		}
	})

	err = http.ListenAndServe(":8080", nil);
	if err != nil {
		log.Printf("Failed to start server %v",err)
	}
}


func restV2LoadData(filePath string) ([]User, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	var users []User
	err = json.Unmarshal(data, &users)

	if err != nil {
		log.Print(err)
		return nil, err
	}
	return users, nil
}
