// Package that contains tests for the cache package.
package cache_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/bauenlabs/echo/lib/cache"
	"os"
	"reflect"
	"testing"
)

// Test the initialization function.
func TestInit(t *testing.T) {
	assert.Equal(t, cache.RedisPort, os.Getenv("ECHO_REDIS_PORT"))
	assert.Equal(t, cache.RedisHost, os.Getenv("ECHO_REDIS_HOST"))
	assert.Equal(t, reflect.TypeOf(cache.Client).String(), "*redis.Client")
}
