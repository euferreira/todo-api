package entities

type Task struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	IsCompleted bool   `json:"isCompleted" bson:"isCompleted"`
	Status      string `json:"status" bson:"status"`
	Active      bool   `json:"active" bson:"active"`
}
type Status struct {
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
