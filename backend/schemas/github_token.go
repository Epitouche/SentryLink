package schemas

import "time"

// GitHubTokenResponse represents the response from Github when a token is requested
type GitHubTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// GithubToken represents the GithubToken entity in the database
type GithubToken struct {
	Id          uint64    `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UserId      uint64    `json:"-"` // Foreign key for User
	User        User      `json:"user,omitempty" gorm:"foreignKey:UserId;references:Id"`
	AccessToken string    `json:"access_token"`
	Scope       string    `json:"scope"`
	TokenType   string    `json:"token_type"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
