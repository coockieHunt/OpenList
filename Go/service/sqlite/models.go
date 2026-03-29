package sqlite

import "time"

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type List struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `json:"title"`
	Items []Item `gorm:"foreignKey:ListID" json:"items"`
}

type Item struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ListID    uint   `json:"list_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	Validated bool   `json:"validated"`
}

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string `gorm:"uniqueIndex" json:"username"`
	PasswordHash string `json:"-"`
	FirstLogin   bool   `json:"first_login"`
}

type Session struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Token     string    `gorm:"uniqueIndex" json:"token"`
	UserID    uint      `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}
