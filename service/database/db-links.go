package database

import "github.com/notherealmarco/overLinks/service/structures"

func (db *appdbimpl) GetLinks() (*[]structures.OverLink, error) {

	links := make([]structures.OverLink, 0)

	rows, err := db.c.Query("SELECT * FROM links")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var l structures.OverLink
		err = rows.Scan(&l.ID, &l.Type, &l.Link, &l.Description)
		if err != nil {
			return nil, err
		}
		links = append(links, l)
	}

	return &links, err
}

func (db *appdbimpl) AddLink(link *structures.OverLinkInput) error {
	_, err := db.c.Exec("INSERT INTO links (type, url, description) VALUES (?, ?, ?)", link.Type, link.Link, link.Description)
	return err
}

func (db *appdbimpl) DeleteLink(id int64) error {
	_, err := db.c.Exec("DELETE FROM links WHERE id=?", id)
	return err
}