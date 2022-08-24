package user

type Login struct {
	UserName string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}
