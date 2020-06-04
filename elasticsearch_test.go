package goelasticsearch

import (
	"testing"
)

func TestNewESClient(t *testing.T) {
	urls := []string{"http://127.0.0.1:9200"}
	_, err := NewESClient(urls, false)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
