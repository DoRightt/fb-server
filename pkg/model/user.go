package model

// User represents all user properties
type User struct {
	UserId    int32  `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Email     string `json:"email,omitempty" yaml:"email,omitempty"`
	Rank      string `json:"rank,omitempty" yaml:"rank,omitempty"`
	Claim     string `json:"claim,omitempty" yaml:"claim,omitempty"`
	Roles     uint64 `json:"roles,omitempty" yaml:"roles,omitempty"`
	Flags     uint64 `json:"flags,omitempty" yaml:"flags,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}

// UserRequest represents a request for user information.
type UserRequest struct {
	UserId int32  `json:"user_id,omitempty" yaml:"user_id,omitempty"`
	Name   string `json:"name,omitempty" yaml:"name,omitempty"`
}

// UsersRequest represents a request for multiple users' information.
type UsersRequest struct {
	UserIds           []int32 `json:"user_ids,omitempty" yaml:"user_ids,omitempty"`
	Name              string  `json:"name,omitempty" yaml:"name,omitempty"`
	Email             string  `json:"email,omitempty" yaml:"email,omitempty"`
	LastActivityFrom  int64   `json:"last_activity_from,omitempty" yaml:"last_activity_from,omitempty"`
	LastActivityUntil int64   `json:"last_activity_until,omitempty" yaml:"last_activity_until,omitempty"`
	ExtraData         bool    `json:"extra_data,omitempty" yaml:"extra_data,omitempty"`
	OnlyCount         bool    `json:"only_count,omitempty" yaml:"only_count,omitempty"`
	ListRequest
}

// UsersRequest represents a request for multiple users' information.
type UserResult struct {
	User     User          `json:"user"`
	Settings *UserSettings `json:"settings,omitempty"`
}

// UserSetttings represents user settings
type UserSettings struct {
	UserId     int32  `json:"user_id"`
	LoginEmail string `json:"login_email"`
	WebFlags   uint64 `json:"web_flags"`
	EmailFlags uint64 `json:"email_flags"`
	UpdatedAt  int64  `json:"updated_at"`
}
