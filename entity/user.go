package entity

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `admin_users`
func (Admin_users) TableName() string {
	return "admin_users"
}

//User represents users table in database
type Admin_users struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Token    string `json:"token,omitempty"`
}
