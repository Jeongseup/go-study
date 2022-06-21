package main

import "log"

// =============== Database =================

type Database interface {
	GetUser() string
	GetAllUsers() []string
}

type DefaultDatabase struct{}

func (db DefaultDatabase) GetUser() string {
	return "seup"
}

type FamouseDatabase struct{}

func (db FamouseDatabase) GetUser() string {
	return "Will smith"
}

func (db FamouseDatabase) GetAllUsers() []string {
	return []string{}
}

// =============== Greeter =================

type Greeter interface {
	Greet(userName string)
}

type NiceGreeter struct{}

func (ng NiceGreeter) Greet(userName string) {
	log.Printf("Hello %s!", userName)
}

// =============== Program =================

type Program struct {
	db      Database
	greeter Greeter
}

func (p Program) Execute() {
	user := p.db.GetUser()
	p.greeter.Greet(user)
}

func main() {

	myDB := FamouseDatabase{}
	myGreeter := NiceGreeter{}

	p := Program{
		db:      myDB,
		greeter: myGreeter,
	}

	p.Execute()
}
