package dao

import (
	"fmt"
	"log"

	. "todoservice/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TodoDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "todos"
)

// Establish a connection to database
func (m *TodoDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal("Error while connecting to MongoDB. Reason: ", err)
	}
	db = session.DB(m.Database)
}

// Find list of todos
func (m *TodoDAO) FindAll() ([]Todo, error) {
	log.Println("Find all the todos.")
	var todos []Todo
	err := db.C(COLLECTION).Find(bson.M{}).All(&todos)
	return todos, err
}

// Find a todo by its id
func (m *TodoDAO) FindById(id string) (Todo, error) {
	var todo Todo
	err := db.C(COLLECTION).FindId(id).One(&todo)
	return todo, err
}

// Insert a todo into database
func (m *TodoDAO) Insert(todo Todo) error {
	err := db.C(COLLECTION).Insert(&todo)
	return err
}

func (m *TodoDAO) Delete(id string) error {
	fmt.Println(id)
	err := db.C(COLLECTION).Remove(bson.M{"_id": id})
	return err
}

func (m *TodoDAO) DeleteAll() error {
	info, err := db.C(COLLECTION).RemoveAll(nil)
	log.Println(info)
	return err
}

// Update an existing todo
func (m *TodoDAO) Update(todo Todo) error {
	err := db.C(COLLECTION).UpdateId(todo.ID, &todo)
	return err
}
