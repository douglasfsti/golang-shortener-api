package shortner

import "time"

type Redirect struct {
	Code      uint64    `json:"code"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
