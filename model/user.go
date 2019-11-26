package model

const (
	// PermissionAddUser describes a permission that the user is allowed
	// to create a new non administrator user.
	PermissionAddUser uint32 = 0
)

// User is the basic user definition of the system, including authorization
// informaiton and role owned by the user.
type User struct {
	// The unique identifier of the entity, for DB storing and indexing.
	ID uint64 `json:"id"`
	// The name of the user, in this case, we use real name for internal
	// management purpose. For general usage, it maybe a nickname or any
	// fake string.
	Name string `json:"name" binding:"required"`
	// The encrypted password of the user by cryptographic library.
	Password string `json:"-"`
	// The permissions of the user, for API access control.
	Permissions []uint32 `json:"permissions"`
}
