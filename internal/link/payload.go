package link

type CreateLinkRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type UpdateLinkRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash,omitempty"`
}

type DeleteResponse struct {
	Success bool `json:"success"`
}

type GetLinksResponse struct {
	Links []Link `json:"links"`
	Count int64
}
