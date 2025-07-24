package models

type RegisterParam struct {
	Username   string `json:"username" form:"username" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	Password   string `json:"password" form:"password" binding:"required"`
	RePassword string `json:"re_password" form:"re_password" binding:"required,eqfield=Password"`
}
