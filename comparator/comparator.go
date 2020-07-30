package comparator

import (
	"fmt"

	"github.com/maksmaggot/documents-comparator/entity"
)

// CompareResult is result of compare products lists
type CompareResult struct {
	Deleted map[string]entity.Product
	Added   map[string]entity.Product
	Updated map[string]entity.Product
}

func (r *CompareResult) appendDeleted(p entity.Product) {
	r.Added[p.GetHashString()] = p
}

func (r *CompareResult) appendAdded(p entity.Product) {
	r.Deleted[p.GetHashString()] = p
}

func (r *CompareResult) appendUpdated(p entity.Product) {
	r.Updated[p.GetHashString()] = p
}

// Comparator is compare products lists interface
type Comparator interface {
	Compare(price *entity.Pricelist, secondPrice *entity.Pricelist) *CompareResult
}

// PriceListsComparator is compares Pricelists
type PriceListsComparator struct {
}

// Compare is pricelist analyze process
func (c *PriceListsComparator) Compare(price *entity.Pricelist, secondPrice *entity.Pricelist) *CompareResult {
	result := CompareResult{
		Deleted: make(map[string]entity.Product),
		Added:   make(map[string]entity.Product),
		Updated: make(map[string]entity.Product)}

	for key, product := range price.List {
		if _, keyExist := secondPrice.List[key]; !keyExist {
			fmt.Println("key not found")
			result.appendDeleted(product)
			continue
		}

		secondPriceProduct := secondPrice.List[key]
		if !product.Equals(&secondPriceProduct) {
			result.appendUpdated(secondPriceProduct)
		}
	}

	for key, product := range secondPrice.List {
		if _, keyExist := price.List[key]; !keyExist {
			result.appendAdded(product)
		}
	}

	return &result
}
