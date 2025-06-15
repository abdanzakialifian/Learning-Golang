package test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:@tcp(localhost:3306)/learning_golang_orm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into sample(id,name) values(?,?)", "1", "Abdan").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id,name) values(?,?)", "2", "Zaki").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id,name) values(?,?)", "3", "Alifian").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   string
	Name string
}

func TestRawSQL(t *testing.T) {
	var sample Sample
	err := db.Raw("select id,name from sample where id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "Abdan", sample.Name)

	var samples []Sample
	err = db.Raw("select * from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(samples))
}

func TestRowsSQL(t *testing.T) {
	rows, err := db.Raw("select * from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample

	for rows.Next() {
		var id string
		var name string

		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			Id:   id,
			Name: name,
		})
	}

	assert.Equal(t, 3, len(samples))
}

func TestScanRowsSQL(t *testing.T) {
	rows, err := db.Raw("select * from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}
	assert.Equal(t, 3, len(samples))
}

func TestCreateUser(t *testing.T) {
	user := User{
		Id:       "1",
		Password: "secret",
		Name: Name{
			FirstName:  "Abdan",
			MiddleName: "Zaki",
			LastName:   "Alifian",
		},
		Information: "This will be ignored",
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 2; i < 10; i++ {
		users = append(users, User{
			Id:       strconv.Itoa(i),
			Password: "secret",
			Name: Name{
				FirstName: "User " + strconv.Itoa(i),
			},
		})
	}

	response := db.Create(&users)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(8), response.RowsAffected)
}
