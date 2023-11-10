package cache

type Cache struct {
	cacheMap map[string]interface{}
}

func New() *Cache {
	cache := make(map[string]interface{})
	return &Cache{cache}
}

func (c Cache) Set(key string, value interface{}) {
	c.cacheMap[key] = value
}

func (c Cache) Get(key string) interface{} {
	if value, ok := c.cacheMap[key]; ok {
		return value
	} else {
		return nil
	}
}

func (c Cache) Delete(key string) {
	if _, ok := c.cacheMap[key]; ok {
		delete(c.cacheMap, key)
	}
}
