package entities

import "time"

type Account struct {
	Id        uint      `gorm:"primary_key"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	RoleId    uint      `gorm:"column:role_id"`
	Verified  bool      `gorm:"column:verified"`
	Active    bool      `gorm:"column:active"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Approval struct {
	Id           uint      `gorm:"primary_key"`
	AdminId      uint      `gorm:"column:admin_id"`
	SuperAdminId uint      `gorm:"column:super_admin_id"`
	Status       bool      `gorm:"column:status"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`

	Admin      Account `gorm:"foreignKey:AdminId"`
	SuperAdmin Account `gorm:"foreignKey:SuperAdminId"`
}
