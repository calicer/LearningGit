package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Employee struct {
	id   string
	name string
	dept string
}

var a = 10

func CreateEmployee(employee Employee) error {

	db, err := sql.Open("postgres", "postgres://postgres:postgres@177.0.0.1:5432/postgres")
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO Employee(id,name,dept) values($1,$2,$3)")

	if err != nil {
		return err
	}
	result, err := stmt.Exec(employee.id, employee.name, employee.dept)

	if err != nil {
		return err
	}
	rowsEff, err := result.RowsAffected()

	fmt.Println("Rows Effected ", rowsEff)

	return err

}

func main() {
	err := CreateEmployee(Employee{id: "1", name: "Vishal", dept: "EAIS"})

	fmt.Println(err)
}
