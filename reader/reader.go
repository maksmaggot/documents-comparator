package reader

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/maksmaggot/documents-comparator/entity"
)

// PricelistReader is interface
type PricelistReader interface {
	Read(string) *entity.Pricelist
}

// CsvReader is PricelistReader for CSV
type CsvReader struct {
}

func (r *CsvReader) Read(filepath string) *entity.Pricelist {
	file, error := os.Open(filepath)
	if error != nil {
		log.Fatal("Cannot open file")
	}
	defer file.Close()

	pricelist := entity.Pricelist{List: make(map[string]entity.Product)}

	reader := csv.NewReader(file)
	reader.Comma = ';'

	header := false
	for {
		row, error := reader.Read()
		if !header {
			header = true
			continue
		}

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal("Can't read file")
		}

		parsedPrice, error := strconv.ParseFloat(row[2], 64)
		if error != nil {
			log.Fatal("Can't parse price")
		}

		item := entity.Product{Name: row[0], Description: row[1], Price: parsedPrice}
		pricelist.Add(item)
	}

	return &pricelist
}
