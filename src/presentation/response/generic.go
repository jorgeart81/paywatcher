package response

type Generic struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Token   interface{} `json:"token,omitempty"`
}

func (u *Generic) Err() *Generic {
	return &Generic{
		Status:  "error",
		Message: u.Message,
	}
}

func (u *Generic) Ok() *Generic {
	return &Generic{
		Status:  "success",
		Message: u.Message,
		Data:    u.Data,
		Token:   u.Token,
	}
}
