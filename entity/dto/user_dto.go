package dto

type UserCreatedorUpdatedDTO struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
	//ProfilePicture string `json:"profile_picture" from:"profile_picture"`
}

type UserDTO struct {
	ID    uint64 `json:"id" form:"id" binding:"required"`
	Name  string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
	//ProfilePicture string `json:"profile_picture" from:"profile_picture"`
	UserType uint64 `json:"user_type" form:"user_type" binding:"required"`
}

type UserIdDTO struct {
	UserId uint64 `gorm:"not null" json:"user_id"`
}
