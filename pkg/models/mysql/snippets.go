package mysql

import (
	"database/sql"

	"github.com/sparkycj328/Snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// Insert will insert a new snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	//create the statement to be used
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?,?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
