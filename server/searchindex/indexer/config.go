package indexer

type Config struct {
	IdProperty string           `json:"idProperty"`
	Indexes    map[string]Index `json:"indexes"`
}

type Index struct {
	Kind string `json:"kind"`
}
