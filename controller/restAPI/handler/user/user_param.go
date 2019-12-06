package user

type createUserParams struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type getUserInfoByEmailParams struct {
	Email string `form:"email" binding:"required"`
}
