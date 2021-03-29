package user

// Role ...
type Role string

// User Roles
const (
	ROLEUSER  Role = "ROLE_USER"
	ROLEADMIN      = "ROLE_ADMIN"
)

// User ...
type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Fullname   string `json:"fullname"`
	Mobile     string `json:"mobile"`
	CreatedAt  int64  `json:"createdAt"`
	ModifiedAt int64  `json:"modifiedAt"`
	Blocked    bool   `json:"blocked"`
	Roles      []Role `json:"roles"`
	isDeleted  bool
}
