package wordgroups

type WordGroupResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"createdAt"`
}

func (wg *WordGroup) ToResponse() WordGroupResponse {
	return WordGroupResponse{
		ID:        wg.ID.Hex(),
		UserID:    wg.UserID.Hex(),
		Name:      wg.Name,
		CreatedAt: wg.CreatedAt,
	}
}
