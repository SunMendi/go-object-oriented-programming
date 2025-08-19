// repo/base_repo.go
// BaseRepo provides common repository functionality, demonstrating inheritance via embedding.

package repo

import "fmt"

// BaseRepo struct can be embedded in other repos to reuse methods.
type BaseRepo struct{}

// Save prints the entity, simulating a save operation (could be extended for DB logic).
func (r BaseRepo) Save(entity interface{}) {
	fmt.Println("Saved entity:", entity)
}
