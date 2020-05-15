package interfacetest

import "fmt"

type PolClass struct {
	a    int
	b    int
	Name string
}

func (c *PolClass) setNameInternal(n string) {
	c.Name = n
	fmt.Println("Set name internal", c.Name)
}

func (c *PolClass) SetTheName(n string) {
	c.setNameInternal(n)
}

// SetNumbers is used to set the numbers a and b
func (c *PolClass) SetNumbers(v1 int, v2 int) {
	c.a = v1
	c.b = v2
}

// Calculate a + b
// Comment line 2
func (c *PolClass) Calculate() int {
	return c.a + c.b
}
