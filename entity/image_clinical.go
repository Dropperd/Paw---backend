package entity

type ImageClinical struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserId     uint64 `gorm:"not null" json:"user_id"`
	ClinicalId uint64 `gorm:"not null" json:"clinical_id"`
}
