package database

import (
	"database/sql"

	"github.com/black-dev-x/go-graphql/graph/model"
	"github.com/google/uuid"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (c *CategoryDB) Create(name string, description string) (*model.Category, error) {
	id := uuid.New().String()
	_, error := c.db.Exec("insert into categories (id, name, description) values ($1, $2, $3)", id, name, description)
	if error != nil {
		return nil, error
	}
	return &model.Category{ID: id, Name: name, Description: description}, error
}

func (c *CategoryDB) GetCategoryById(id string) (*model.Category, error) {
	row := c.db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id)
	var category model.Category
	if err := row.Scan(&category.ID, &category.Name, &category.Description); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No category found
		}
		return nil, err // Other error
	}
	return &category, nil
}

func (c *CategoryDB) FindAll() ([]*model.Category, error) {
	rows, err := c.db.Query("Select id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []*model.Category{}
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (c *CategoryDB) FindByCourseId(id string) (*model.Category, error) {
	row := c.db.QueryRow("SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1", id)
	var category model.Category
	if err := row.Scan(&category.ID, &category.Name, &category.Description); err != nil {
		return nil, err
	}
	return &category, nil
}
