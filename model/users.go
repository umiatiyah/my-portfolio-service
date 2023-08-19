package model

type Users struct {
	ID        int    `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Email     string `json:"email" db:"email"`
	Status    int    `json:"status" db:"status"`
	CreatedBy string `json:"created_by" db:"created_by"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedBy string `json:"updated_by" db:"updated_by"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
