package gocd

type PGPipeline struct {
	Name  string  `json:"name,omitempty"`
	Links PGLinks `json:"_links"`
}

type Authorization struct {
	View   Permissions `json:"view"`
	Admins Permissions `json:"admins"`
}

type Permissions struct {
	Users []string `json:"users"`
	Roles []string `json:"roles"`
}

type PipelineGroup struct {
	Links         PGLinks       `json:"_links"`
	Name          string        `json:"name,omitempty"`
	Pipelines     []PGPipeline  `json:"pipelines,omitempty"`
	Authorization Authorization `json:"authorization"`
}

type PGResponse struct {
	Links    PGLinks    `json:"_links"`
	Embedded PGEmbedded `json:"_embedded"`
}

type PGEmbedded struct {
	Groups []PipelineGroup `json:"groups"`
}

type PGLinks struct {
	Self Link `json:"self"`
	Doc  Link `json:"doc"`
	Find Link `json:"find"`
}

// GetPipelineGroups List pipeline groups along with the pipelines, stages and materials for each pipeline.
func (c *DefaultClient) GetPipelineGroups() ([]PipelineGroup, error) {

	var res PGResponse
	err := c.getJSON("/go/api/admin/pipeline_groups", map[string]string{"Accept": "application/vnd.go.cd.v1+json"}, &res)

	if err != nil {
		return []PipelineGroup{}, err
	}

	groups := res.Embedded.Groups

	return groups, nil

}
