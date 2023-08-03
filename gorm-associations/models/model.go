package models

import (
	"fmt"
	"gopractise-demo/gorm-associations/utils"
)

type UserEntity struct {
	Id   int    `json:"id" gorm:"type:int(11);autoIncrement;primaryKey;column:id;"`
	Name string `json:"name" gorm:"type:varchar(30);not null;"`
	Age  int    `json:"age" gorm:"type:int(11);default:0;"`
}

type IdCardEntity struct {
	Id int    `json:"id" gorm:"type:int(11);autoIncrement;primaryKey;column:id"`
	No string `json:"no" gorm:"type:varchar(20);not null;"`

	// Foreign key association user
	User   UserEntity `json:user`
	UserId int        `json:"user_id"`
}

func (UserEntity) TableName() string {
	return "user"
}

func (IdCardEntity) TableName() string {
	return "id_card"
}

type StudentEntity struct {
	Id      int    `json:"id" gorm:"type:int(11);autoIncrement;primaryKey;column:id"`
	Name    string `json:"name" gorm:"type:varchar(30);not null;comment:姓名"`
	ClassId int    `json:"class_id" gorm:"type:int(11);not null;comment:教室ID"`
}

type ClassEntity struct {
	Id   int    `json:"id" gorm:"type:int(11);autoIncrement;primaryKey;column:id"`
	Name string `json:"name" gorm:"type:varchar(30);not null"`
	// One class association via foreign key
	StudentList []StudentEntity `json:"student_list" gorm:"foreignKey:ClassId"`
}

func (StudentEntity) TableName() string {
	return "student"
}

func (ClassEntity) TableName() string {
	return "class"
}

func init() {
	//utils.DB.AutoMigrate(&UserEntity{}, &IdCardEntity{})
	utils.DB.AutoMigrate(&ClassEntity{}, &StudentEntity{})
}

//-------- One by One ------

func GetUserListEntity() {
	var userListEntity []map[string]interface{}
	utils.DB.Model(&UserEntity{}).Select("user.id", "user.name", "user.age", "id_card.no").
		Joins("LEFT JOIN id_card ON user.id = id_card.user_id").Scan(&userListEntity)
	fmt.Println(utils.MapToJSON(userListEntity))
}

func PreloadIdCardListEntity() {
	var idCardListEntity []IdCardEntity
	utils.DB.Preload("User").Find(&idCardListEntity, 3)
	fmt.Println(utils.MapToJSON(idCardListEntity))
}

//-------- One to many ------

func GetStudents() {
	var students []StudentEntity
	utils.DB.Find(&students)
	fmt.Println(utils.MapToJSON(students))
}

func GetClassList() {
	var classList []ClassEntity
	utils.DB.Preload("StudentList").Find(&classList, 3)
	fmt.Println(utils.MapToJSON(classList))
}
