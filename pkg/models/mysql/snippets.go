package mysql

import (
	"database/sql"
	"github.com/mailtomainak/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// Insert a snippet.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
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

// Get a snippet
func (m *SnippetModel) Get(snippetId int) (*models.Snippet, error) {
	return nil, nil
}

// Get the 10  latest created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
