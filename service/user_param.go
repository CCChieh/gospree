package service

type CreateUserParams struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserInfoByIDParams struct {
	ID uint `form:"id" binding:"required"`
}

type CreateTokenParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
