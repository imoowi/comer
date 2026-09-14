package components

import "testing"

func TestNewRedisClient(t *testing.T) {
	c := NewRedisClient(&RedisConfig{Addr: "localhost:6379", Password: "pw", DB: 1, PoolSize: 10})
	if c == nil {
		t.Fatal("NewRedisClient returned nil")
	}
	opts := c.Options()
	if opts.Addr != "localhost:6379" || opts.Password != "pw" || opts.DB != 1 || opts.PoolSize != 10 {
		t.Errorf("NewRedisClient options = %+v", opts)
	}
}
