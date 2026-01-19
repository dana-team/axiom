package types

type QueryParams struct {
	ClusterID   string `form:"clusterId"`
	KubeVersion string `form:"kubeVersion"`
	Name        string `form:"name"`
	Limit       int    `form:"limit,default=100"`
	SortBy      string `form:"sortBy,default=lastUpdated"`
	SortOrder   string `form:"sortOrder,default=desc"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Count   int64       `json:"count,omitempty"`
}
