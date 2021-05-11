package domain

import (
	"backend/internal/types"
	"gorm.io/gorm"
	"time"
)

type UserRepositoryFace interface {
	CheckExistedAccountAndRegister(*types.Register) error
	ChekLogin(*types.Login) (*User, error)
}

type UserLogicFace interface {
	UserRegister(*types.Register) error
	UserLogin(*types.Login) (string, error)
	UserLogout(string) bool
}

type User struct {
	gorm.Model
	UserName     string    `gorm:"column:username;type:varchar(50);not null;uniqueIndex"`
	Password     string    `gorm:"column:password;type:varchar(60);not null"`
	Age          uint8     `gorm:"column:age;type:tinyint;default:0;check:age >= 0"`
	Gender       string    `gorm:"column:gender;type:enum('no','male','female');default:no"`
	NickName     string    `gorm:"column:nickname;type:varchar(60);default:null"`
	Avatar       string    `gorm:"column:avatar;type:varchar(150)"`
	Birthday     time.Time `gorm:"column:birthday;type:date"`
	Address      []Address `gorm:"foreignKey:UserID"`
	Introduction string    `gorm:"column:introduction;type:varchar(500);default:null"`
	Email        string    `gorm:"column:email;type:varchar(60);uniqueIndex"`
	Phone        string    `gorm:"column:phone;type:char(11);uniqueIndex;default:null"`
	State        string    `gorm:"column:state;type:enum('0','1','2', '-1');default:0"`
	Inviter      uint      `gorm:"column:inviter;default:null"`
	InviteCode   string    `gorm:"column:invite_code;type:varchar(50);default:null"`
	Teams        []Team    `gorm:"many2many:user_teams"`
	Followers    []*User   `gorm:"many2many:user_followers"`
	RoleID       uint8     `gorm:"column:role_id;type:tinyint;default:null"`
	Role         Role      `gorm:"foreignKey:RoleID"`
}

type Role struct {
	ID        uint8          `gorm:"primaryKey"`
	Name      string         `gorm:"column:name;type:varchar(50);uniqueIndex"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Team struct {
	gorm.Model
	Name  string `gorm:"column:name;type:varchar(50);uniqueIndex"`
	Label string `gorm:"column:label;type:varchar(50)"`
	Tag   string `gorm:"column:tag;type:varchar(60)"`
}

type Address struct {
	gorm.Model
	UserID  uint
	Addr    string `gorm:"column:addr;type:varchar(200)"`
	Default string `gorm:"column:default;type:enum('0', '1');default:'1'"`
}

func NewUser() *User {
	return &User{}
}

func NewRole() *Role {
	return &Role{}
}

func NewTeam() *Team {
	return &Team{}
}

func NewAddress() *Address {
	return &Address{}
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Model(u).Where("id = ?", u.ID).Update("birthday", u.Birthday.Format("2006-01-02"))
	return err
}
