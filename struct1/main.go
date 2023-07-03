package main

import (
	"fmt"
	"time"
)

type Employee struct{
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

func main(){
	var dilbert Employee

	dilbert.Salary -= 5000 
	position := &dilbert.Position
	*position = "Senior"+*position

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"

	fmt.Println(dilbert.Position)
	fmt.Println(employeeOfTheMonth.Position)	
}