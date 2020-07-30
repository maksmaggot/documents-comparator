package main

import (
	"fmt"

	"github.com/maksmaggot/documents-comparator/comparator"
	"github.com/maksmaggot/documents-comparator/reader"
)

func main() {
	priceListReader := new(reader.CsvReader)
	price := priceListReader.Read("testdata/pricelist.csv")
	secondPrice := priceListReader.Read("testdata/pricelist2.csv")
	fmt.Println(secondPrice)
	comparator := new(comparator.PriceListsComparator)
	result := comparator.Compare(price, secondPrice)
	fmt.Println(result.Added)
}
