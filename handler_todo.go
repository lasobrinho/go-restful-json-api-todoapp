package main 

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)



func GetAllTodos() Todos {
	todoCollection := dbSession.DB("todo-go").C("todo")
	
	var todos Todos
	
	err := todoCollection.Find(bson.M{}).All(&todos)
	
	if err != nil {
		fmt.Println(err)
	}

	return todos
}

func GetTodoCount() int {
	todoCollection := dbSession.DB("todo-go").C("todo")
	count, err := todoCollection.Count()
	
	if err != nil {
		fmt.Println(err)
	}

	return count
}

//func CreateTodo(t Todo) Todo {

//}

//func DestroyTodo(id int) error {
	
//}