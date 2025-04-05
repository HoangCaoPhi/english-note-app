package wordgroups

type CreateWordGroupRequest struct {
	Name string `json:"name" validate:"required"`
}
