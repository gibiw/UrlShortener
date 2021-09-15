package repository

import (
	"log"

	link "github.com/gibiw/UrlShortener"
	"github.com/hashicorp/go-memdb"
)

const linkitemTable string = "link"

type LinkItemRepository struct {
	db *memdb.MemDB
}

func NewLinkItemRepository() *LinkItemRepository {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			linkitemTable: &memdb.TableSchema{
				Name: "link",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Original"},
					},
					"modification": &memdb.IndexSchema{
						Name:    "modification",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Modification"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &LinkItemRepository{db: db}
}

func (r *LinkItemRepository) Create(o, e string) (string, error) {
	txn := r.db.Txn(true)

	link := &link.LinkItem{
		Original:     o,
		Modification: e,
	}

	if err := txn.Insert(linkitemTable, link); err != nil {
		return "", err
	}

	txn.Commit()

	return e, nil
}

func (r *LinkItemRepository) GetByUrl(url string) (link.LinkItem, error) {
	txn := r.db.Txn(false)
	defer txn.Abort()

	var l link.LinkItem
	raw, err := txn.First(linkitemTable, "original", url)
	if err != nil {
		return l, err
	}

	l = *raw.(*link.LinkItem)

	return l, nil
}

func (r *LinkItemRepository) GetByHash(hash string) (string, error) {
	txn := r.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First(linkitemTable, "modification", hash)
	if err != nil {
		return "", err
	}

	l := *raw.(*link.LinkItem)

	return l.Original, nil
}
