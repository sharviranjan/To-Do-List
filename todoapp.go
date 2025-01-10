package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)
type Employee struct {
	Name string
	Todo []string
}

type Assign struct {
	Status string
	Result string
}
var e1todo=make([]string,0,2)
var e2todo=make([]string,0,2)

var EmployeeDB=map[string][]string{
	"employee1":e1todo,
	"employee2":e2todo,
}

func handler_function(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path{
	case "/TODO":
		var fileName="index.html"
		t,err:=template.ParseFiles(fileName)
		if err !=nil{
			fmt.Println("Error when parsing the file",err)
			return
		}
		err= t.ExecuteTemplate(w, fileName, nil)
		if err !=nil{
			fmt.Println("Error when executing template",err)
			return
		}


	case "/manager":
		var fileName="manager.html"
		t,err:=template.ParseFiles(fileName)
		if err !=nil{
			fmt.Println("Error when parsing the file",err)
			return
		}
		err= t.ExecuteTemplate(w, fileName, "Manager")
		if err !=nil{
			fmt.Println("Error when executing template",err)
			return
		}
	
	case "/asignDone":
		var fileName="assignDone.html"
		var result string
		var status string
		t,err:=template.ParseFiles(fileName)
		if err !=nil{
			fmt.Println("Error when parsing the file",err)
			return
		}
		task:= r.FormValue("task")
		assignee:= r.FormValue("assignee")
		_,i := EmployeeDB[assignee]
		if i {
			EmployeeDB[assignee]=append(EmployeeDB[assignee],task)
			result=assignee+" has been assigned the task "+task
			status="Task Successful"

		}else{
			status="Task Unsuccessful"
			result="No such employee in database"
		}

		err= t.ExecuteTemplate(w, fileName, Assign{status , result})
		if err !=nil{
			fmt.Println("Error when executing template",err)
			return
		}

	case "/employee1":
		var fileName="employee.html"
		t,err:=template.ParseFiles(fileName)
		if err !=nil{
			fmt.Println("Error when parsing the file",err)
			return
		}
		err= t.ExecuteTemplate(w, fileName, Employee{"Employee1", EmployeeDB["employee1"]})
		if err !=nil{
			fmt.Println("Error when executing template",err)
			return
		}


	case "/employee2":
		var fileName="employee.html"
		t,err:=template.ParseFiles(fileName)
		if err !=nil{
			fmt.Println("Error when parsing the file",err)
			return
		}
		err= t.ExecuteTemplate(w, fileName, Employee{"Employee2", EmployeeDB["employee2"]})
		if err !=nil{
			fmt.Println("Error when executing template",err)
			return
		}


	default:
		fmt.Fprintf(w,"<h1 style=\"color: darkblue\" align=\"center\"><u>???</u></h1>")
	}

}

func main() {
    http.HandleFunc("/", handler_function)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

//http://localhost:8080/TODO