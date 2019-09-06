package message

// LogoRequest represents service input
type LogoRequest struct {
	ID       string `json:"id"`
	ImageURL string `json:"url"`
}
