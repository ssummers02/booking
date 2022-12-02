package restmodel

type User struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	Surname    string `json:"surname"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email" validate:"email,required"`
	Password   string `json:"password" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	RoleID     int64  `json:"role_id" validate:"required"`
	Token      string `json:"token"`
}
