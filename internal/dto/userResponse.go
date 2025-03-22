package dto

type LoginResponse struct {
	Message string      `json:"message"`
	Token   string      `json:"token,omitempty"`
	Role    uint        `json:"role,omitempty"`
	Users   interface{} `json:"users,omitempty"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}
