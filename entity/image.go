package entity

type Image struct {
	ID           uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Image        string `gorm:"type:mediumblob" form:"image" json:"image"`
	Description  string `gorm:"type:varchar(255)" form:"description" json:"description" binding:"required"`
	BodyPosition string `gorm:"type:varchar(255)" form:"body_position" json:"body_position" binding:"required"`
	UserID       uint64 `gorm:"not null" json:"user_id"`
	Added_At     string `gorm:"type:timestamp" json:"added_at"`
	Updated_At   string `gorm:"type:timestamp" json:"updated_at"`
}
