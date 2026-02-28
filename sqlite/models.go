package sqlite

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
