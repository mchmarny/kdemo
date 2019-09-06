package message

// LogoResponse represents service output
type LogoResponse struct {
	Request     LogoRequest `json:"req"`
	Description string      `json:"desc"`
}
