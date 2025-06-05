package database

import (
	"database/sql"

	"github.com/black-dev-x/go-graphql/graph/model"
	"github.com/google/uuid"
)

type CourseDB struct {
	db *sql.DB
}

func NewCourseDB(db *sql.DB) *CourseDB {
	return &CourseDB{db: db}
}

func (c *CourseDB) Create(name string, description string, categoryID string) (*model.Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, categoryID)
	if err != nil {
		return nil, err
	}
	return &model.Course{ID: id, Name: name, Description: description}, nil
}

func (c *CourseDB) FindAll() ([]*model.Course, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []*model.Course{}
	for rows.Next() {
		var course model.Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description); err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}

func (c *CourseDB) FindByCategoryID(categoryID string) ([]*model.Course, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM courses WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []*model.Course{}
	for rows.Next() {
		var course model.Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description); err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	return courses, nil
}
