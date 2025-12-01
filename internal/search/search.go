package search

import (
	"github.com/zarazaex69/zuk/pkg/zuk"
)

type Result = zuk.Result

func Search(query string) ([]Result, error) {
	client := zuk.NewClient()
	return client.Search(query)
}
