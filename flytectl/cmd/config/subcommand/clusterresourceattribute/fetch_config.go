package clusterresourceattribute

//go:generate pflags AttrFetchConfig --default-var DefaultFetchConfig

type AttrFetchConfig struct {
	AttrFile string `json:"attrFile" pflag:",attribute file name to be used for generating attribute for the resource type."`
}

var DefaultFetchConfig = &AttrFetchConfig{}
