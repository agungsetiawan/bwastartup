package user

type RegisterUserInput struct {
	Name       string `json:"name" form:"name" binding:"required"`
	Occupation string `json:"occupation" form:"occupation" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	Password   string `json:"password" form:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
