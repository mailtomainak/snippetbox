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
	var snippet = &models.Snippet{}
	stmt := `SELECT id,content,title,created,expires FROM snippets where expires > UTC_TIMESTAMP and id=?`
	if err := m.DB.QueryRow(stmt, snippetId).Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created,
		&snippet.Expires); err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return snippet, nil
}

// Get the 10  latest created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {

	stmt := `SELECT id,content,title,created,expires FROM snippets where expires > UTC_TIMESTAMP ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Expires, &s.Created)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
