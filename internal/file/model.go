package file

type Record struct {
	Name      string `json:"name"`
	Size      int    `json:"size"`
	MD5       string `json:"md5"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}