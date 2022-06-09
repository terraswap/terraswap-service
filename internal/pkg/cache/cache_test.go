package cache

import (
	"testing"

	"github.com/delight-labs/terraswap-service/configs"
	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {

	cache := New(configs.CacheConfig{})

	testSets := []struct {
		key      string
		value    string
		expected string
	}{
		{key: "a", value: "a", expected: "a"},
		{key: "b", value: "b", expected: "b"},
	}

	for _, tcase := range testSets {
		cache.Set(tcase.key, tcase.value)
	}

	for _, tcase := range testSets {
		v := cache.Get(tcase.key)
		v = v.(string)
		assert.Equal(t, v, tcase.expected)
	}

}
