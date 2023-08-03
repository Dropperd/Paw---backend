package entity

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	UserType uint64 `gorm:"not null" json:"user_type"`
	//ProfilePicture string `gorm:"type:varchar(1024);default:'https://cdn-icons-png.flaticon.com/512/5987/5987462.png'" json:"profile_picture"`
	Token string `gorm:"-" json:"token"`
}
