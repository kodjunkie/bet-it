// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"elivate9ja-go/app/admin/controllers"
	"elivate9ja-go/app/admin/repositories"
	"elivate9ja-go/app/admin/services"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/uptrace/bun"
)

// Injectors from wire.go:

func InitializeAdminController(db bun.IDB) (*controllers.AdminController, error) {
	adminRepository, err := repositories.NewAdminRepository(db)
	if err != nil {
		return nil, err
	}
	adminService, err := services.NewAdminService(adminRepository)
	if err != nil {
		return nil, err
	}
	validate := provideValidator()
	adminController, err := controllers.NewAdminController(adminService, validate)
	if err != nil {
		return nil, err
	}
	return adminController, nil
}

// wire.go:

func provideValidator() *validator.Validate {
	return validator.New()
}

var AdminControllerSet = wire.NewSet(services.NewAdminService, wire.Bind(new(services.IAdminService), new(*services.AdminService)), repositories.NewAdminRepository, wire.Bind(new(repositories.IAdminRepository), new(*repositories.AdminRepository)), provideValidator)
