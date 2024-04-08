package databases

import "fmt"

type MySQL struct{}

func (m *MySQL) Insert() error {
	fmt.Println("MySQL Insert!!!")
	return nil
}

func (m *MySQL) Update() error {
	fmt.Println("MySQL Update!!!")
	return nil
}
