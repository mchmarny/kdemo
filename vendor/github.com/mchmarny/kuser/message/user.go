package message

import "time"

// KUser represents service input
type KUser struct {
	ID      string    `json:"id" firestore:"id"`
	Email   string    `json:"email" firestore:"email"`
	Name    string    `json:"name" firestore:"name"`
	Created time.Time `json:"created" firestore:"created"`
	Updated time.Time `json:"updated" firestore:"updated"`
	Picture string    `json:"pic" firestore:"pic"`
}
