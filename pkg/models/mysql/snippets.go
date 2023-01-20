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
	return 0, nil
}

// Get a snippet
func (m *SnippetModel) Get(snippetId int) (*models.Snippet, error) {
	return nil, nil
}

// Get the 10  latest created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
