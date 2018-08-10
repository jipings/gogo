package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"imooc.com/tutorial/crawler/engine"
	"imooc.com/tutorial/crawler/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	if err == nil {
		*result = "ok"
	}

	return err

}
