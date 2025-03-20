package link

type CreateLinkRequest struct {
	Url string `json:"url" validate:"required,url"`
}
