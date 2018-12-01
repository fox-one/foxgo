package request

type Pagination struct {
	NextCursor string `json:"nextCursor"`
	HasNext    bool   `json:"hasNext"`
}
