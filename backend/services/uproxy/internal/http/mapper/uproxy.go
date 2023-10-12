package mapper

// Feature Toggles data
type ToggleList struct {
	Version  int `json:"version" example:"2"`
	Features []struct {
		Name         string `json:"name" example:"CreateUserToggle"`
		Description  string `json:"description" example:"Create User Toggle"`
		Environments []struct {
			Name    string `json:"name" example:"development"`
			Enabled bool   `json:"enabled" example:"true"`
			Type    string `json:"type" example:"development"`
		} `json:"environments"`
	} `json:"features"`
}
