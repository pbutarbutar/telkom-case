package middleware

//organization
type Organization struct {
	ID                       string `json:"id"`
	Name                     string `json:"name"`
	LeadsOwner               string `json:"leads_owner"`
	ShowCommission           bool   `json:"show_commission"`
	EnableQa                 bool   `json:"enable_qa"`
	ShowLevelInDashboard     bool   `json:"show_level_in_dashboard"`
	ShowShortcutsInDashboard bool   `json:"show_shortcuts_in_dashboard"`
	UseSimpleLeadStatus      bool   `json:"use_simple_lead_status"`
	ShowMemberStructure      bool   `json:"show_member_structure"`
}

// User contains user's information
type User struct {
	TokenType    string       `json:"token_type"`
	Exp          int          `json:"exp"`
	Jti          string       `json:"jti"`
	UserID       string       `json:"user_id"`
	HumanID      string       `json:"human_id"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	FullName     string       `json:"full_name"`
	Email        string       `json:"email"`
	IsActive     bool         `json:"is_active"`
	IsStaff      bool         `json:"is_staff"`
	IsSuperuser  bool         `json:"is_superuser"`
	PhoneNumber  string       `json:"phone_number"`
	Groups       []string     `json:"groups"`
	Organization Organization `json:"organization"`
	TokenAuth    string
}

// NewUser returns a new user
func NewUser(userID string, humanID string, email string) (*User, error) {
	user := &User{
		UserID:  userID,
		HumanID: humanID,
		Email:   email,
	}

	return user, nil
}

// Clone returns a clone of this user
func (user *User) Clone() *User {
	return &User{
		UserID:  user.UserID,
		HumanID: user.HumanID,
		Email:   user.Email,
	}
}
