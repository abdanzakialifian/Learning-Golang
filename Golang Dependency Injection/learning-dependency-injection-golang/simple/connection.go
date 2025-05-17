package simple

import "fmt"

type Connection struct {
	File *File
}

func NewConnection(file *File) (*Connection, func()) {
	connection := Connection{File: file}
	return &connection, func() {
		connection.Close()
	}
}

func (connection *Connection) Close() {
	fmt.Println("Close Connection", connection.File.Name)
}
