package schemas

import (
	"path/filepath"

	"github.com/composeDB/coredb/internal/belt"
	"github.com/composeDB/coredb/internal/utilities"
)

type CoreDB struct {
	Name        string `json:"name"`
	IndexCursor *belt.DB
	Pages       uint64 `json:"pages"`
	PageSize    uint64 `json:"page_size"`
}

func (c *CoreDB) Connect() {
	config, err := utilities.LoadConfig(utilities.GetHomeDir())
	utilities.CheckError(err)
	db, err := belt.Open(filepath.Join(config.Datafolder, c.Name), 0600, nil)
	utilities.CheckError(err)
	defer db.Close()
	c.IndexCursor = db
}

func (c *CoreDB) Disconnect() {
	c.IndexCursor.Close()
}

func (c *CoreDB) Insert(data [][]byte) {
	err := c.IndexCursor.Batch(func(tx *belt.Tx) error {
		b := tx.Bucket([]byte("data"))
		for i := range data {
			id, _ := b.NextSequence()
			b.Put(utilities.IntToBytes(id), []byte(data[i]))
			return nil
		}
		return nil
	})
	utilities.CheckError(err)
}
