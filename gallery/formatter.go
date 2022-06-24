package gallery

type GalleryFormatter struct {
	Image     string `json:"image_url"`
	IsDefault bool   `json:"is_default"`
}
