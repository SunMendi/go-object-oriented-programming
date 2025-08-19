// repo/product_repo.go
// ProductRepo demonstrates repository pattern and inheritance via embedding BaseRepo.

package repo

import "oop_with_go/models"

// ProductRepo embeds BaseRepo to inherit common repository methods.
type ProductRepo struct {
	BaseRepo // embedding BaseRepo for method reuse
}

// FindAll returns a list of products (dummy data for demonstration).
func (r ProductRepo) FindAll() []models.Product {
	return []models.Product{
		*models.NewProduct(1, "Shirt", 500, 10),
		*models.NewProduct(2, "Pant", 800, 5),
	}
}
