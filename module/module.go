package module

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Module interface define functions a Skel module needs to implement
type Module interface {
	Migrate(db *gorm.DB) error
	Seed(db *gorm.DB) error
	SetupRoutes(*gorm.DB, *gin.Engine) error
}

// Registry stores the list of active modules
type Registry struct {
	Modules []Module
}

// Register a module
func (r *Registry) Register(m Module) {
	r.Modules = append(r.Modules, m)
}

// RegisterAll registers a list of modules
func (r *Registry) RegisterAll(modules []Module) {
	for _, m := range modules {
		r.Register(m)
	}
}

// Migrate create and modify database tables according to the models
func (r *Registry) Migrate(db *gorm.DB) {
	for _, m := range r.Modules {
		m.Migrate(db)
	}
}

// Seed import seed files into database
func (r *Registry) Seed(db *gorm.DB) {
	for _, m := range r.Modules {
		m.Seed(db)
	}
}

// SetupRoutes import seed files into database
func (r *Registry) SetupRoutes(db *gorm.DB, app *gin.Engine) {
	for _, m := range r.Modules {
		m.SetupRoutes(db, app)
	}
}
