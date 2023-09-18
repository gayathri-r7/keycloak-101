package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	. "todoservice/config"
	. "todoservice/dao"
	. "todoservice/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = TodoDAO{}

func ListTodosEndPoint(w http.ResponseWriter, r *http.Request) {
	log.Println("List all todos API hit!")
	todos, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(todos) != 0 {
		respondWithJson(w, http.StatusOK, todos)
	} else {
		todos := []Todo{}
		respondWithJson(w, http.StatusOK, todos)
	}
}

// GET a todo by its ID
func FindTodoEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todo, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Todo ID")
		return
	}
	respondWithJson(w, http.StatusOK, todo)
}

// DELETE an existing todo
func DeleteTodoEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := dao.Delete(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"message": "Todo deleted successfully"})
}

// DELETE an existing todo
func DeleteTodosEndPoint(w http.ResponseWriter, r *http.Request) {
	err := dao.DeleteAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"message": "Todos deleted successfully"})
}

// POST a new todo
func CreateTodoEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	uuid := uuid.New().String()
	UsrID := strings.Replace(uuid, "-", "", -1)

	todo.ID = "tsk" + string(UsrID)
	todo.Status = "In-Progress"
	if err := dao.Insert(todo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	var resp Success
	resp.Message = "Todo created successfully."
	resp.TodoID = todo.ID
	respondWithJson(w, http.StatusCreated, resp)
}

// PUT update an existing todo
func UpdateTodoEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usrId := params["id"]
	defer r.Body.Close()
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	todo.ID = usrId
	if err := dao.Update(todo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"message": "Todo updated successfully."})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	log.Println("MongoDB Server: ", dao.Server)
	log.Println("MongoDB Database: ", dao.Database)
	log.Println("Trying to connect with mongodb...")
	dao.Connect()
	log.Println("Connected to MongoDB!")
	log.Println("Todo-Service is running on the port 3000!")
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todo", CreateTodoEndPoint).Methods("POST")
	r.HandleFunc("/todo/{id}", FindTodoEndpoint).Methods("GET")
	r.HandleFunc("/todo/{id}", UpdateTodoEndPoint).Methods("PUT")
	r.HandleFunc("/todo/{id}", DeleteTodoEndPoint).Methods("DELETE")
	r.HandleFunc("/todos", ListTodosEndPoint).Methods("GET")
	r.HandleFunc("/todos", DeleteTodosEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
