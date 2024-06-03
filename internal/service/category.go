package service

import (
	"context"

	"github.com/rgoncalvesrr/fullcycle-grpc/internal/database"
	"github.com/rgoncalvesrr/fullcycle-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(context.Context, *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesResp []*pb.Category

	for _, categorie := range categories {
		categoriesResp = append(categoriesResp, &pb.Category{
			Id:          categorie.ID,
			Name:        categorie.Name,
			Description: categorie.Description,
		})
	}

	return &pb.CategoryList{
		Categories: categoriesResp,
	}, nil
}
