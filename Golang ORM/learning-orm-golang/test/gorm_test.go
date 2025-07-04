package test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

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

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

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

func TestLock(t *testing.T) {
	db.Transaction(func(tx *gorm.DB) error {
		user := new(User)
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(user, "id = ?", "1").Error

		if err != nil {
			return err
		}

		user.Name.FirstName = "Abdan"
		user.Name.MiddleName = "Alifian"

		err = tx.Save(user).Error

		return err
	})
}

func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		Id:      "1",
		UserId:  "1",
		Balance: 1000000,
	}

	err := db.Create(&wallet).Error
	assert.Nil(t, err)
}

func TestRetrieveRelation(t *testing.T) {
	var user = new(User)
	err := db.Preload("Wallet").Take(user, "id = ?", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.Id)
	assert.Equal(t, "1", user.Wallet.UserId)
}

func TestRetrieveRelationJoin(t *testing.T) {
	var user = new(User)
	err := db.Joins("Wallet").Take(user, "user.id = ?", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.Id)
	assert.Equal(t, "1", user.Wallet.UserId)
}

func TestAutoCreateUpdate(t *testing.T) {
	user := User{
		Id:       "20",
		Password: "secret",
		Name: Name{
			FirstName: "User 20",
		},
		Wallet: Wallet{
			Id:      "20",
			UserId:  "20",
			Balance: 1000000,
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
}

func TestSkipAutoCreateUpdate(t *testing.T) {
	user := User{
		Id:       "21",
		Password: "secret",
		Name: Name{
			FirstName: "User 21",
		},
		Wallet: Wallet{
			Id:      "21",
			UserId:  "21",
			Balance: 1000000,
		},
	}

	err := db.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
}

func TestUserAndAddress(t *testing.T) {
	user := User{
		Id:       "2",
		Password: "secret",
		Name: Name{
			FirstName: "User 2",
		},
		Wallet: Wallet{
			Id:      "2",
			UserId:  "2",
			Balance: 1000000,
		},
		Addresses: []Address{
			{
				UserId:  "2",
				Address: "Street A",
			},
			{
				UserId:  "2",
				Address: "Street B",
			},
		},
	}

	err := db.Save(&user).Error
	assert.Nil(t, err)
}

func TestPreloadJoinsOneToMany(t *testing.T) {
	var users []User
	err := db.Preload("Addresses").Joins("Wallet").Find(&users).Error
	assert.Nil(t, err)
}

func TestTakePreloadJoinsOneToMany(t *testing.T) {
	user := new(User)
	err := db.Preload("Addresses").Joins("Wallet").Take(user, "user.id = ?", "50").Error
	assert.Nil(t, err)
}

func TestBelongsToOneToMany(t *testing.T) {
	fmt.Println("============Preload============")
	var addresses []Address
	err := db.Preload("User").Find(&addresses).Error
	assert.Nil(t, err)

	fmt.Println("============Joins============")

	addresses = []Address{}
	err = db.Joins("User").Find(&addresses).Error
	assert.Nil(t, err)
}

func TestBelongToOneToOne(t *testing.T) {
	fmt.Println("============Preload============")
	var wallets []Wallet
	err := db.Preload("User").Find(&wallets).Error
	assert.Nil(t, err)

	fmt.Println("============Joins============")
	wallets = []Wallet{}
	err = db.Joins("User").Find(&wallets).Error
	assert.Nil(t, err)
}

func TestCreateManyToMany(t *testing.T) {
	product := Product{
		Id:    "P001",
		Name:  "Example Product",
		Price: 1000000,
	}

	err := db.Save(&product).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]any{
		"user_id":    "1",
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]any{
		"user_id":    "2",
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)
}

func TestPreloadManyToManyProduct(t *testing.T) {
	product := new(Product)
	err := db.Preload("LikedByUsers").Take(product, "id = ?", "P001").Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(product.LikedByUsers))
}

func TestManyToManyUser(t *testing.T) {
	user := new(User)
	err := db.Preload("LikeProducts").Take(user, "id = ?", "1").Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(user.LikeProducts))
}

func TestAssociationFind(t *testing.T) {
	var product = new(Product)
	err := db.Take(product, "id = ?", "P001").Error
	assert.Nil(t, err)

	var users []User
	err = db.Model(product).Where("first_name LIKE ?", "User%").Association("LikedByUsers").Find(&users)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestAssociationAppend(t *testing.T) {
	user := new(User)
	err := db.Take(user, "id = ?", "3").Error
	assert.Nil(t, err)

	product := new(Product)
	err = db.Take(product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(user).Association("LikeProducts").Append(product)
	assert.Nil(t, err)
}

func TestAssociationReplace(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		user := new(User)
		err := tx.Take(user, "id = ?", "1").Error
		assert.Nil(t, err)

		wallet := Wallet{
			Id:      "01",
			UserId:  user.Id,
			Balance: 2000000,
		}

		err = tx.Model(user).Association("Wallet").Replace(&wallet)
		return err
	})
	assert.Nil(t, err)
}

func TestAssociationDelete(t *testing.T) {
	user := new(User)
	err := db.Take(user, "id = ?", "3").Error
	assert.Nil(t, err)

	product := new(Product)
	err = db.Take(product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(user).Association("LikeProducts").Delete(product)
	assert.Nil(t, err)
}

func TestAssociationClear(t *testing.T) {
	product := new(Product)
	err := db.Take(product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(product).Association("LikedByUsers").Clear()
	assert.Nil(t, err)
}

func TestPreloadingWithCondition(t *testing.T) {
	user := new(User)
	err := db.Preload("Wallet", "balance >= ?", 1000000).Take(user, "id = ?", "1").Error
	assert.Nil(t, err)
	fmt.Println(user)
}

func TestPreloadingNested(t *testing.T) {
	var wallet = new(Wallet)
	err := db.Preload("User.Addresses").Take(wallet, "id = ?", "2").Error
	assert.Nil(t, err)
	fmt.Println(wallet)
	fmt.Println(wallet.User)
	fmt.Println(wallet.User.Addresses)
}

func TestPreloadingAll(t *testing.T) {
	var user = new(User)
	err := db.Preload(clause.Associations).Take(user, "id = ?", "1").Error
	assert.Nil(t, err)
}

func TestJoinQuery(t *testing.T) {
	var users []User
	err := db.Joins("join wallet on wallet.user_id = user.id").Find(&users).Error // inner join
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	users = []User{}
	err = db.Joins("Wallet").Find(&users).Error // left join
	assert.Nil(t, err)
	assert.Equal(t, 19, len(users))
}

func TestJoinWithCondition(t *testing.T) {
	var users []User
	err := db.Joins("join wallet on wallet.user_id = user.id AND wallet.balance > ?", 500000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	users = []User{}
	err = db.Joins("Wallet").Where("Wallet.balance > ?", 500000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

func TestCount(t *testing.T) {
	var count int64
	err := db.Model(&User{}).Joins("Wallet").Where("Wallet.balance > ?", 500000).Count(&count).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(4), count)
}

type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestAggregation(t *testing.T) {
	var result = new(AggregationResult)
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").Take(result).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(4000000), result.TotalBalance)
	assert.Equal(t, int64(1000000), result.MinBalance)
	assert.Equal(t, int64(1000000), result.MaxBalance)
	assert.Equal(t, float64(1000000), result.AvgBalance)
}

func TestAggregationGroupByAndHaving(t *testing.T) {
	var results []AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance",
		"max(balance) as max_balance", "avg(balance) as avg_balance").Joins("User").Group("User.id").Having("sum(balance) > ?", 1000000).Find(&results).Error
	assert.Nil(t, err)
	assert.Equal(t, 0, len(results))
}

func TestContext(t *testing.T) {
	ctx := context.Background()

	var users []User
	err := db.WithContext(ctx).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 19, len(users))
}

func BrokeWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance = ?", 0)
}

func SultanWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance > ?", 1000000)
}

func TestScopes(t *testing.T) {
	var wallets []Wallet
	err := db.Scopes(BrokeWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)

	wallets = []Wallet{}
	err = db.Scopes(SultanWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)
}

func TestMigrator(t *testing.T) {
	err := db.Migrator().AutoMigrate(&GuestBook{})
	assert.Nil(t, err)
}

func TestHook(t *testing.T) {
	user := User{
		Password: "secret",
		Name: Name{
			FirstName: "User 100",
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
	assert.NotEqual(t, "", user.Id)

	fmt.Println(user.Id)
}
