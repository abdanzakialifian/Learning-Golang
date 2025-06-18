package test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:@tcp(localhost:3306)/learning_golang_orm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
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

func TestTransactionSuccess(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{Id: "10", Password: "secret", Name: Name{FirstName: "User 10"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{Id: "11", Password: "secret", Name: Name{FirstName: "User 11"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{Id: "12", Password: "secret", Name: Name{FirstName: "User 12"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func TestTransactionError(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{Id: "13", Password: "secret", Name: Name{FirstName: "User 13"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{Id: "11", Password: "secret", Name: Name{FirstName: "User 11"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.NotNil(t, err)
}

func TestManualTransactionSuccess(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{Id: "14", Password: "secret", Name: Name{FirstName: "User 14"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{Id: "15", Password: "secret", Name: Name{FirstName: "User 15"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestManualTransactionError(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{Id: "16", Password: "secret", Name: Name{FirstName: "User 16"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{Id: "15", Password: "secret", Name: Name{FirstName: "User 17"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestQuerySingleObject(t *testing.T) {
	user := User{}
	err := db.First(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.Id)

	user = User{}
	err = db.Last(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "9", user.Id)
}

func TestQuerySingleObjectInlineCondition(t *testing.T) {
	user := User{}
	err := db.Take(&user, "id = ?", "5").Error
	assert.Nil(t, err)
	assert.Equal(t, "5", user.Id)
	assert.Equal(t, "User 5", user.Name.FirstName)
}

func TestQueryAllObjects(t *testing.T) {
	var users []User
	err := db.Find(&users, "id in ?", []string{"1", "2", "3", "4", "5"}).Error
	assert.Nil(t, err)
	assert.Equal(t, 5, len(users))
}

func TestQueryCondition(t *testing.T) {
	var users []User
	err := db.Where("first_name like ?", "%User%").Where("password = ?", "secret").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 14, len(users))
}

func TestOrOperator(t *testing.T) {
	var users []User
	err := db.Where("first_name like ?", "%User%").Or("password = ?", "secret").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 15, len(users))
}

func TestNotOperator(t *testing.T) {
	var users []User
	err := db.Not("first_name like ?", "%User%").Where("password = ?", "secret").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestSelectFields(t *testing.T) {
	var users []User
	err := db.Select("id", "first_name").Find(&users).Error
	assert.Nil(t, err)

	for _, user := range users {
		assert.NotNil(t, user.Id)
		assert.Equal(t, "", user.Password)
	}

	assert.Equal(t, 15, len(users))
}

func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "User 5",
		},
		Password: "secret",
	}

	user := new(User)
	err := db.Where(userCondition).Take(user).Error
	assert.Nil(t, err)
	assert.Equal(t, "5", user.Id)
}

func TestMapCondition(t *testing.T) {
	mapCondition := map[string]any{
		"middle_name": "",
		"last_name":   "",
	}

	var users []User
	err := db.Where(mapCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 14, len(users))
}

func TestOrderLimitOffset(t *testing.T) {
	var users []User
	err := db.Order("id asc, first_name desc").Limit(5).Offset(5).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 5, len(users))
}

func TestQueryNonModel(t *testing.T) {
	type UserResponse struct {
		Id        string `gorm:"column:id"`
		FirstName string `gorm:"column:first_name"`
		LastName  string `gorm:"column:last_name"`
	}
	var users []UserResponse
	err := db.Model(new(User)).Select("id", "first_name", "last_name").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 15, len(users))
}

func TestUpdate(t *testing.T) {
	user := new(User)
	err := db.Take(user, "id = ?", "1").Error
	assert.Nil(t, err)

	user.Name.FirstName = "Agustin"
	user.Name.MiddleName = "Dini"
	user.Name.LastName = "Anugraeni"
	user.Password = "secret123"

	err = db.Save(user).Error
	assert.Nil(t, err)
}

func TestUpdateSelectedColumns(t *testing.T) {
	user := new(User)
	err := db.Model(user).Where("id = ?", "1").Updates(map[string]any{
		"middle_name": "Zaki",
		"last_name":   "",
	}).Error
	assert.Nil(t, err)

	err = db.Model(user).Where("id = ?", "1").Update("password", "secret12321").Error
	assert.Nil(t, err)

	err = db.Where("id = ?", "1").Updates(User{
		Name: Name{
			FirstName: "Abdan",
			LastName:  "Alifian",
		},
	}).Error
	assert.Nil(t, err)
}

func TestAutoIncrement(t *testing.T) {
	for range 10 {
		userLog := UserLog{
			UserId: "1",
			Action: "Test Action",
		}

		err := db.Create(&userLog).Error
		assert.Nil(t, err)

		assert.NotEqual(t, 0, userLog.Id)
		fmt.Println(userLog.Id)
	}
}

func TestSaveUpdate(t *testing.T) {
	userLog := UserLog{
		UserId: "1",
		Action: "Test Action",
	}

	err := db.Save(&userLog).Error // insert
	assert.Nil(t, err)

	userLog.UserId = "2"
	err = db.Save(&userLog).Error // update
	assert.Nil(t, err)
}

func TestSaveOrUpdateNonAutoIncrement(t *testing.T) {
	user := User{
		Id: "99",
		Name: Name{
			FirstName: "User 99",
		},
	}

	err := db.Save(&user).Error // update
	assert.Nil(t, err)

	user.Name.FirstName = "User 99 Updated"
	err = db.Save(&user).Error // insert
	assert.Nil(t, err)
}

func TestConflict(t *testing.T) {
	user := User{
		Id: "77",
		Name: Name{
			FirstName: "User 77",
		},
	}

	err := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&user).Error
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	user := new(User)
	err := db.Take(user, "id = ?", "88").Error
	assert.Nil(t, err)

	err = db.Delete(user).Error
	assert.Nil(t, err)

	err = db.Delete(user, "id = ?", "99").Error
	assert.Nil(t, err)

	err = db.Delete(user).Where("id = ?", "77").Error
	assert.Nil(t, err)
}

func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserId:      "1",
		Title:       "Todo 1",
		Description: "Description 1",
	}
	err := db.Create(&todo).Error
	assert.Nil(t, err)

	err = db.Delete(&todo).Error
	assert.Nil(t, err)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	err = db.Find(&todos).Error
	assert.Nil(t, err)
	assert.Equal(t, 0, len(todos))
}

func TestUnscope(t *testing.T) {
	todo := new(Todo)
	err := db.Unscoped().Take(todo).Where("id = ?", 6).Error
	assert.Nil(t, err)
	fmt.Println(todo)

	err = db.Unscoped().Delete(todo).Error
	assert.Nil(t, err)

	var todos []Todo
	err = db.Unscoped().Find(&todos).Error
	assert.Nil(t, err)
}
