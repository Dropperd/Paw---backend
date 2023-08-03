package entity

type BodyPos struct {
	Id     uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Membro string `gorm:"type:varchar(255)" json:"membro"`
}
