package databases

import "fmt"

type MockDatabase struct{}

func (m *MockDatabase) Insert() error {
	fmt.Println("MockDB Insert!!!")
	return nil
}

func (m *MockDatabase) Update() error {
	fmt.Println("MockDB Update!!!")
	return nil
}
