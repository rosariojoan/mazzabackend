package notifications

import "time"

type User struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Tagline string   `json:"tagline"`
	Avatar  *string  `json:"avatar"`
	Skills  []string `json:"skills"`
}

type FinishedProject struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Category    string    `json:"category"`
	Subcategory string    `json:"subcategory"`
	Price       float64   `json:"price"`
	CompletedAt time.Time `json:"completedAt"`
	Freelancer  *User     `json:"freelancer"`
	Client      *User     `json:"client"`
}
