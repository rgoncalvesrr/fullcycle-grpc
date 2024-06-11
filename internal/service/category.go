package service

import (
	"context"
	"io"

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

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindByID(in.CategoryId)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}

		if err != nil {
			return err
		}

		result, err := c.CategoryDB.Create(category.Name, category.Description)

		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          result.ID,
			Name:        result.Name,
			Description: result.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryStreamVaiVem(stream pb.CategoryService_CreateCategoryStreamVaiVemServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		result, err := c.CategoryDB.Create(category.Name, category.Description)

		if err != nil {
			return err
		}

		err = stream.Send(&pb.Category{
			Id:          result.ID,
			Name:        result.Name,
			Description: result.Description,
		})

		if err != nil {
			return err
		}
	}
}
