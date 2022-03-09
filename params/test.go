package params

//


type Test struct {
	Email    string `form:"email" failed:"账号错误" validate:"required,email"`
	Password string `form:"password" failed:"密码错误" validate:"required"`
}