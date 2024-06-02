package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (*Category, error) {
	c.ID = uuid.New().String()
	c.Name = name
	c.Description = description

	stmt, err := c.db.Prepare("insert into categories (id, name, description) values ($1, $2, $3)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.ID, c.Name, c.Description)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("select id, name, description from categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []Category{}

	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		result = append(result, Category{
			ID:          id,
			Name:        name,
			Description: description,
		})
	}

	return result, nil
}

func (c *Category) FindByCourse(courseID string) (*Category, error) {
	var id, name, description string

	err := c.db.QueryRow(`
		select
			ca.id, ca.name, ca.description
		from
			categories ca
		join
			courses co on
			co.category_id = ca.id and
			co.id = $1`, courseID).Scan(&id, &name, &description)

	if err != nil {
		return nil, err
	}

	result := &Category{
		ID:          id,
		Name:        name,
		Description: description,
	}

	return result, nil
}
