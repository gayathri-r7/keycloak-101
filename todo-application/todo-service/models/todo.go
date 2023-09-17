package models

type Todo struct {
	ID          string `bson:"_id" json:"id"`
	Task        string `bson:"task" json:"task"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
	Completed   bool   `bson:"completed,omitempty" json:"completed,omitempty"`
}

type Success struct {
	Message string `bson:"message" json:"message"`
	TodoID  string `bson:"todoId" json:"todoId"`
}
