package persists

import (
	"go-spider/persist"
	"log"

	"go-spider/engine"

	"github.com/olivere/elastic"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("error saveing item %v: %v", item, err)
	}
	return err
}
