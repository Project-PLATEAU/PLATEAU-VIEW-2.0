package indexer

type IndexRoot struct {
	ResultDataUrl string                 `json:"resultDataUrl"`
	IdProperty    string                 `json:"idProperty"`
	Indexes       map[string]interface{} `json:"indexes"`
}

type EnumIndex struct {
	Kind   string                `json:"kind"`
	Values map[string]*EnumValue `json:"values"`
}

type EnumValue struct {
	Count int    `json:"count"`
	Url   string `json:"url"`
}
