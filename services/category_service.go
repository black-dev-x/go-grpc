package services

import (
	"context"

	"github.com/black-dev-x/go-grpc/database"
	"github.com/black-dev-x/go-grpc/internal/pb"
	"google.golang.org/grpc"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.CategoryDB
}

func NewCategoryService(db database.CategoryDB) *CategoryService {
	return &CategoryService{
		CategoryDB: db,
	}
}

func (s *CategoryService) GetCategoryById(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := s.CategoryDB.GetCategoryById(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (s *CategoryService) GetAllCategories(ctx context.Context, req *pb.Blank) (*pb.CategoryList, error) {
	categories, err := s.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	var categoryList []*pb.Category
	for _, category := range categories {
		categoryList = append(categoryList, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return &pb.CategoryList{
		Categories: categoryList,
	}, nil
}

func (s *CategoryService) CreateCategoryStream(stream grpc.BidiStreamingServer[pb.CreateCategoryRequest, pb.Category]) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		category, err := s.CategoryDB.Create(req.Name, req.Description)
		if err != nil {
			return err
		}

		categoryCreated := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		if err := stream.Send(categoryCreated); err != nil {
			return err
		}
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := s.CategoryDB.Create(req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	categoryCreated := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryCreated, nil
}
