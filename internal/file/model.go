package file

import "time"

type Record struct {
	Name      string    `json:"name"`
	Size      int       `json:"size"`
	MD5       string    `json:"md5"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
