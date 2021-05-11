package types

type Register struct {
	UserName   string `json:"username" binding:"required,min=4,max=50,excludesrune=@"`
	Password   string `json:"password" binding:"required,min=6,max=18"`
	ConfirmPwd string `json:"confirm_pwd" binding:"required,eqfield=Password"`
	Email      string `json:"email" binding:"required,email"`
}

type Login struct {
	Account  string `json:"account" binding:"required,min=4,max=50"`
	Password string `json:"password" binding:"required,min=6,max=18"`
}

type Token struct {
	Token string `json:"token"`
}
