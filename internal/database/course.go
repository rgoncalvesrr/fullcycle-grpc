package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryId string) (*Course, error) {

	c.ID = uuid.NewString()
	c.Name = name
	c.Description = description
	c.CategoryId = categoryId

	stmt, err := c.db.Prepare("insert into courses (id, name, description, category_id) values ($1, $2, $3, $4)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.ID, c.Name, c.Description, c.CategoryId)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("select id, name, description, category_id from courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []Course{}

	for rows.Next() {
		var id, name, description, categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}

		result = append(result, Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryId:  categoryId,
		})
	}

	return result, nil
}

func (c *Course) FindByCategory(categoryID string) ([]Course, error) {

	rows, err := c.db.Query(`
		select 
			id, name, description 
		from 
			courses 
		where 
			category_id = $1`, categoryID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []Course{}

	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		result = append(result, Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryId:  categoryID,
		})
	}

	return result, nil
}
