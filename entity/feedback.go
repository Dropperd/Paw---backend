package entity

type Feedback struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Feedback   string `gorm:"type:text" json:"feedback"`
	ImageId    uint64 `gorm:"not null" json:"image_id"`
	UserId     uint64 `gorm:"not null" json:"user_id"`
	IdClinical uint64 `gorm:"not null" json:"id_clinical"`
	Added_At   string `gorm:"type:timestamp" json:"added_at"`
	Updated_At string `gorm:"type:timestamp" json:"updated_at"`
}
