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
		panic(err)
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

func GetTodoById(todoId string) (Todo, error) {
	todoCollection := dbSession.DB("todo-go").C("todo")

	var todo Todo

	err := todoCollection.Find(bson.M{"_id": bson.ObjectIdHex(todoId)}).One(&todo)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func CreateTodo(todo Todo) error {
	todoCollection := dbSession.DB("todo-go").C("todo")

	err := todoCollection.Insert(todo)

	return err
}

//func DestroyTodo(id int) error {
	
//}