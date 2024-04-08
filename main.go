package main

import "github.com/MarkTBSS/021_Interface/databases"

type IDatabase interface {
	Insert() error
	Update() error
}

func InsertAnyDatabase(iDatabaseInstance IDatabase) error {
	return iDatabaseInstance.Insert()
}

func UpdateAnyDatabase(iDatabaseInstance IDatabase) error {
	return iDatabaseInstance.Update()
}

func main() {
	postgresInstance := &databases.Postgres{}
	InsertAnyDatabase(postgresInstance)
	UpdateAnyDatabase(postgresInstance)

	mysqlInstance := &databases.MySQL{}
	InsertAnyDatabase(mysqlInstance)
	UpdateAnyDatabase(mysqlInstance)

	mockDatabaseInstance := &databases.MockDatabase{}
	InsertAnyDatabase(mockDatabaseInstance)
	UpdateAnyDatabase(mockDatabaseInstance)
}
