package main

import (
	"fmt"
)

type (
	Database interface {
		Insert() error
		Update() error
	}

	PogrestDB struct{}
	MockDB    struct{}
)


func NewPostgresDB() Database{
	return &PogrestDB{}
}

func NewMockDB() Database {
	return &MockDB{}
}

func (p *PogrestDB) Insert() error {
	fmt.Println("Insert from Real DB")
	return nil
}

func (p *PogrestDB)Update() error {
	fmt.Println("Update from Real DB")
	return nil
}

func (m *MockDB)Insert() error {
	fmt.Println("Insert from Mock DB")
	return nil
}

func (m * MockDB)Update() error {
	fmt.Println("Update from Mock DB")
	return nil
}


func InsertItemShop(db Database) error {
	return db.Insert()
}

func UpdateItemShop(db Database) error {
	return  db.Update()
}

func main(){
	postgres := &PogrestDB{}
	mock := &MockDB{}

	InsertItemShop(postgres)
	UpdateItemShop(postgres)

	InsertItemShop(mock)
	UpdateItemShop(mock)
}