package fd

import "testing"

func TestCreateEmp(t *testing.T) {
	err := CreateEmployee(Employee{id: "1", name: "Vishal", dept: "EAIS"})

}
