package moduletest

type Curle struct {
	name string
	age  int
}

func (c *Curle) Getname() string {
	return c.name
}

func (c *Curle) SetName(name string) {
	c.name = name
}

func (c *Curle) GetAge() int {
	return c.age
}

func (c *Curle) SetAge(age int) {
	c.age = age
}
