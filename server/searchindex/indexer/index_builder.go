package indexer

type IndexBuilder interface {
	AddIndexValue(int, string)
	GetValueIds() *map[string][]Ids
}

type EnumIndexBuilder struct {
	Property string
	Config   Index
	ValueIds map[string][]Ids
}

type Ids struct {
	DataRowId int
}

func (enumBuilder EnumIndexBuilder) AddIndexValue(dataRowId int, value string) {
	enumBuilder.ValueIds[value] = append(enumBuilder.ValueIds[value], Ids{dataRowId})
}

func (enumBuilder EnumIndexBuilder) GetValueIds() *map[string][]Ids {
	return &enumBuilder.ValueIds
}

func createIndexBuilder(property string, indexConfig Index) IndexBuilder {
	switch indexConfig.Kind {
	case "enum":
		return EnumIndexBuilder{
			Property: property,
			Config:   indexConfig,
			ValueIds: make(map[string][]Ids),
		}
	}
	return nil
}
