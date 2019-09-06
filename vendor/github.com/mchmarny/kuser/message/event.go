package message

import "time"

// KUserEvent represents service input
type KUserEvent struct {
	ID     string       `json:"id" firestore:"id"`
	UserID string       `json:"userId" firestore:"userId"`
	On     time.Time    `json:"ts" firestore:"ts"`
	Data   []*KDataItem `json:"data" firestore:"data"`
}

// KDataItem represents single item in the KUserEvent
type KDataItem struct {
	Key   string `json:"key" firestore:"key"`
	Value string `json:"value" firestore:"value"`
}
