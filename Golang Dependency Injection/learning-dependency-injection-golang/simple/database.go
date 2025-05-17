package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL *Database
type DatabaseMongoDB *Database

func NewDatabasePostgreSQL() DatabasePostgreSQL {
	return &Database{Name: "PostgreSQL"}
}

func NewDatabaseMongoDB() DatabaseMongoDB {
	return &Database{Name: "MongoDB"}
}

type DatabaseRepository struct {
	DatabasePostgreSQL DatabasePostgreSQL
	DatabaseMongoDB    DatabaseMongoDB
}

func NewDatabaseRepository(databasePostgreSQL DatabasePostgreSQL, databaseMongoDB DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: databasePostgreSQL, DatabaseMongoDB: databaseMongoDB}
}
