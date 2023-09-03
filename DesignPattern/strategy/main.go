package strategy

// 策略模式
// 可以定义一系列算法，并将算法分别放入独立的类中，以使算法的对象能够相互替换

type EvictionAlgo interface {
	evict(c *Cache)
}

type Cache struct {
}
