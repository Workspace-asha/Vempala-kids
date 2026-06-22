package pkg

type Success struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type Error struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
