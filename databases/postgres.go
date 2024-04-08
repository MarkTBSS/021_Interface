package databases

import "fmt"

type Postgres struct{}

func (p *Postgres) Insert() error {
	fmt.Println("Postgres Insert!!!")
	return nil
}

func (p *Postgres) Update() error {
	fmt.Println("Postgres Update!!!")
	return nil
}
