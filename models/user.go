package models

// User user model
type User struct {
	Id           string `json:"_id" bson:"_id,omitempty"`
	FirstName    string `json:"first_name" bson:"first_name" binding:"required"`
	LastName     string `json:"last_name" bson:"last_name" binding:"required"`
	Email        string `json:"email" bson:"email" binding:"required"`
	Password     string `json:"password" bson:"password" binding:"required"`
	MobileNumber string `json:"mobile_number" bson:"mobile_number" binding:"required"`
	Role         string `json:"role" bson:"role" binding:"required"`
	Joined       int64  `json:"joined" bson:"joined"`
}
type UserRes struct {
	Id           string `json:"_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Role         string `json:"role"`
	Joined       int64  `json:"joined"`
}

func (u User) Res() UserRes {
	return UserRes{
		Id:           u.Id,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		MobileNumber: u.MobileNumber,
		Role:         u.Role,
		Joined:       u.Joined,
	}
}

type LoginCredentials struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type VerifyRequest struct {
	Otp   string `json:"otp" binding:"required"`
	Email string `json:"email" binding:"required"`
}
