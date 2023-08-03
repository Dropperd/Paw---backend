package dto

type ImageUpdate struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Description string `gorm:"type:varchar(255)" form:"description" json:"description" binding:"required"`
	UserID      uint64 `gorm:"not null" json:"user_id"`
	Added_At    string `gorm:"type:timestamp" json:"added_at"`
	Updated_At  string `gorm:"type:timestamp" json:"updated_at"`
}
