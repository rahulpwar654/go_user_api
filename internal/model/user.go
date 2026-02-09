package model

// User represents a person using the system.
// swagger:model User
type User struct {
	// Example: 1
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement" example:"1"`
	// Example: Alice
	Name string `json:"name" gorm:"size:255;not null" example:"Alice"`
	// Example: alice@example.com
	Email string `json:"email" gorm:"size:255;uniqueIndex;not null" example:"alice@example.com"`
}

// UsersListResponse is the paginated response for listing users.
// swagger:model UsersListResponse
type UsersListResponse struct {
	// data
	Data []User `json:"data"`
	// page number
	Page int `json:"page" example:"1"`
	// page size
	Limit int `json:"limit" example:"20"`
	// total items
	Total int64 `json:"total" example:"100"`
	// total pages
	TotalPages int `json:"total_pages" example:"5"`
	// next page url
	Next string `json:"next" example:"/users?page=2&limit=20"`
	// prev page url
	Prev string `json:"prev" example:"/users?page=1&limit=20"`
}

// UserPatch represents fields allowed for partial update.
// swagger:model UserPatch
type UserPatch struct {
	Name  *string `json:"name" example:"Alice Updated"`
	Email *string `json:"email" example:"alice.new@example.com"`
}
