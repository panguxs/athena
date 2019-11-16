package memory

import (
	"fmt"
	"time"

	"github.com/akyoto/cache"
)

func test() {
	// New cache
	c := cache.New(5 * time.Minute)

	// Put something into the cache
	c.Set("a", "b", 1*time.Minute)

	// Read from the cache
	obj, found := c.Get("a")

	fmt.Println(found)
	// Convert the type
	fmt.Println(obj.(string))
}
