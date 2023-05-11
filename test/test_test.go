package test_test

import (
	"testing"

	"github.com/misseven0/httpcache/httpcache"
	"github.com/misseven0/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
