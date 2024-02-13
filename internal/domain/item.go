package domain

type ItemsIdsRequest struct {
	Items []string `json:"items_ids"`
}

type ItemsIdsResponse struct {
	Prices []Price `json:"prices"`
}
