package resources

type NodeNetworkStateCurrentState struct {
	Interfaces []Interface `json:"interfaces,omitempty" yaml:"interfaces,omitempty"`
}

type Interface struct {
	Ipv4 Ipv4 `json:"ipv4,omitempty" yaml:"ipv4,omitempty"`
}

type Ipv4 struct {
	Address []Address `json:"address,omitempty" yaml:"address,omitempty" `
}

type Address struct {
	IP           string `json:"ip,omitempty" yaml:"ip,omitempty"`
	PrefixLength int    `json:"prefix-length,omitempty" yaml:"prefix-length,omitempty"`
}

type NetBoxResponse struct {
	Count   int            `json:"count,omitempty"`
	Results []NetBoxResult `json:"results,omitempty"`
}

type NetBoxResult struct {
	ID           int          `json:"id,omitempty"`
	Prefix       string       `json:"prefix,omitempty"`
	CustomFields CustomFields `json:"custom_fields,omitempty"`
}

type CustomFields struct {
	Cluster string `json:"Cluster,omitempty"`
}
