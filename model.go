package api

type Short struct {
	Id        string `json:"_" db:"id"`
	ShortCode string `json:"short_code" binding:"required"`
	LongURL   string `json:"long_url" binding:"required"`
	ShortURL  string `json:"short_url" binding:"required"`
}
