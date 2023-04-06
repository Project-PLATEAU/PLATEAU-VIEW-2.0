package indexer

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

const (
	resultsDataCSV = "resultsData.csv"
	indexRootJSON  = "indexRoot.json"
)

type Writer struct {
	config *Config
	o      OutputFS
}

func NewWriter(config *Config, o OutputFS) *Writer {
	return &Writer{
		config: config,
		o:      o,
	}
}

func (w *Writer) Write(r Result) error {
	resultsDataUrl, err := w.WriteResultsData(r.Data)
	if err != nil {
		return err
	}

	indexes, err := w.WriteIndexes(r.IndexBuilders)
	if err != nil {
		return err
	}

	return w.writeIndexRoot(IndexRoot{
		ResultDataUrl: resultsDataUrl,
		IdProperty:    w.config.IdProperty,
		Indexes:       indexes,
	})
}

// Writes the data.csv file and returns its path.
func (w *Writer) WriteResultsData(data ResultData) (string, error) {
	fileName := resultsDataCSV
	f, err := w.o.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	cw := csv.NewWriter(f)
	defer cw.Flush()

	var keys []string
	for k := range data[0] {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	if err := cw.Write(keys); err != nil {
		return "", fmt.Errorf("error writing header for csv: %v", err)
	}

	for _, record := range data {
		row := make([]string, 0, 1+len(keys))

		for _, k := range keys {
			row = append(row, record[k])
		}

		if err := cw.Write(row); err != nil {
			return "", fmt.Errorf("error writing record to file: %v", err)
		}
	}

	return fileName, nil
}

func (w *Writer) writeIndexRoot(indexRoot IndexRoot) error {
	fileName := indexRootJSON

	fw, err := w.o.Open(fileName)
	if err != nil {
		return fmt.Errorf("error while writing the indexRoot: %v", err)
	}

	defer fw.Close()

	if err := json.NewEncoder(fw).Encode(indexRoot); err != nil {
		return fmt.Errorf("error while writing the indexRoot: %v", err)
	}

	return nil
}

// Write indexes using the index builders and returns a IndexRoot.indexes map
func (w *Writer) WriteIndexes(indexBuilders []IndexBuilder) (_ map[string]any, err error) {
	indexes := make(map[string]any)
	count := 0

	for _, b := range indexBuilders {
		switch t := b.(type) {
		case EnumIndexBuilder:
			indexes[t.Property], err = w.WriteIndex(t, count)
			if err != nil {
				return nil, fmt.Errorf("failed to write index: %v", err)
			}
		default:
			continue
		}
		count++
	}

	return indexes, nil
}

func (w *Writer) WriteIndex(enumBuilder EnumIndexBuilder, fileId int) (_ *EnumIndex, err error) {
	values := make(map[string]*EnumValue)
	count := 0

	for name, value := range enumBuilder.ValueIds {
		values[name], err = w.WriteValueIndex(fileId, count, value)
		if err != nil {
			return nil, fmt.Errorf("failed to write value index: %v", err)
		}
		count++
	}

	return &EnumIndex{
		Values: values,
		Kind:   "enum",
	}, nil
}

func (w *Writer) WriteValueIndex(fileId int, valueId int, ids []Ids) (*EnumValue, error) {
	fileName := strconv.Itoa(fileId) + "-" + strconv.Itoa(valueId) + ".csv"
	f, err := w.o.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %v", err)
	}

	defer f.Close()
	cw := csv.NewWriter(f)
	defer cw.Flush()

	if err := cw.Write([]string{"dataRowId"}); err != nil {
		return nil, fmt.Errorf("error writing header for csv: %v", err)
	}

	for _, record := range ids {
		row := []string{strconv.Itoa(record.DataRowId)}
		if err := cw.Write(row); err != nil {
			return nil, fmt.Errorf("error writing record to file: %v", err)
		}
	}

	return &EnumValue{
		Count: len(ids),
		Url:   fileName,
	}, nil
}
