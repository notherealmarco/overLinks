package structures

type OverLink struct {
	ID          int64  `json:"id"`
	Type        string `json:"type"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

type OverLinkInput struct {
	Type        string `json:"type"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

type GenericResponse struct {
	Status string `json:"status"`
}
