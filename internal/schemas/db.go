package schemas

import (
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
	db, err := belt.Open(config.Datafolder, 0600, nil)
	utilities.CheckError(err)
	defer db.Close()
	c.IndexCursor = db
}
