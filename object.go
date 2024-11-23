package gopen

type Object struct {
	ApiVersion string     `json:"api-version"`
	Kind       string     `json:"kind"`
	Meta       ObjectMeta `json:"meta"`
}

type ObjectUpdate struct {
	Kind       string           `json:"kind"`
	ApiVersion string           `json:"api-version"`
	Meta       ObjectUpdateMeta `json:"meta"`
}

type ListMeta struct {
	TotalCount int `json:"total-count"`
}

type ObjectMeta struct {
	Name            string            `json:"name"`
	Tenant          string            `json:"tenant"`
	Namespace       string            `json:"namespace"`
	GenerationId    string            `json:"generation-id"`
	ResourceVersion string            `json:"resource-version"`
	UUID            string            `json:"uuid"`
	Labels          map[string]string `json:"labels"`
	CreationTime    string            `json:"creation-time"`
	ModTime         string            `json:"mod-time"`
	SelfLink        string            `json:"self-link"`
}

type ObjectUpdateMeta struct {
	Labels []string `json:"labels"`
}

type ApiResponseList struct {
	Kind       string   `json:"kind"`
	ApiVersion string   `json:"api-version"`
	ListMeta   ListMeta `json:"list-meta"`
}
type ApiResponseObject struct {
	Kind       string     `json:"kind"`
	ApiVersion string     `json:"api-version"`
	Meta       ObjectMeta `json:"meta'`
}

type PropogationStatus struct {
	GenerationId string `json:"generation-id"`
	Updated      int    `json:"updated"`
	Pending      int    `json:"pending'`
	MinVersion   string `json:"min-version"`
	Status       string `json:"status"`
}
