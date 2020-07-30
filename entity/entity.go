package entity

import (
	"crypto/md5"
	"fmt"
	"sync"
)

// Product is simple item
type Product struct {
	Name        string
	Description string
	Price       float64
}

// GetHashString get md5 from name
func (p *Product) GetHashString() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(p.Name)))
}

// Equals is comparing products
func (p *Product) Equals(p2 *Product) bool {
	if p.Name != p2.Name || p.Price != p2.Price || p.Description != p2.Description {
		return false
	}
	return true

}

// Pricelist is collection of product items
type Pricelist struct {
	List map[string]Product
	sync.Mutex
}

// Delete item by key
func (pl *Pricelist) Delete(key string) {
	delete(pl.List, key)
}

// Add item to collection
func (pl *Pricelist) Add(p Product) {
	pl.Lock()
	defer pl.Unlock()
	pl.List[p.GetHashString()] = p
}
