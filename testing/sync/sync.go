package v1

// Counter will increment a number.
type Counter struct {
	value int
}

// Inc the count.
func (c *Counter) Inc()  {
	c.value++
}

// Value returns the current value
func (c *Counter) Value() int {
	return c.value
}