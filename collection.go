package genderize

import "time"

// Gender type.
type Gender struct {
	Name        string  `json:"name,omitempty"`
	Gender      string  `json:"gender,omitempty"`
	Probability float64 `json:"probability,omitempty"`
	Count       int64   `json:"count,omitempty"`
}

// Collection of genders.
type Collection struct {
	info    *Info
	genders map[string]*Gender
}

// Limit returns the amount of names available in the current time window.
func (c *Collection) Limit() (l int64) {
	if c.info != nil {
		l = c.info.Limit
	}

	return
}

// LimitRemaining returns the number of names left in the current time window.
func (c *Collection) LimitRemaining() (r int64) {
	if c.info != nil {
		r = c.info.Remaining
	}

	return
}

// LimitReset returns seconds remaining until a new time window opens.
func (c *Collection) LimitReset() (d time.Duration) {
	if c.info != nil {
		d = c.info.Reset
	}

	return
}

// Length of collection.
func (c *Collection) Length() int {
	return len(c.genders)
}

// Find gender info by name.
func (c *Collection) Find(name string) (g *Gender, err error) {
	var ok bool
	if g, ok = c.genders[name]; !ok {
		err = ErrNothingFound
	}

	return
}

// FindX like Find, but panics when error.
func (c *Collection) FindX(name string) *Gender {
	g, err := c.Find(name)
	if err != nil {
		panic(err)
	}

	return g
}

// First gender of collection.
func (c *Collection) First() (g *Gender, err error) {
	for _, g = range c.genders {
		return
	}

	err = ErrNothingFound

	return
}

// FirstX like First, but panics when error.
func (c *Collection) FirstX() *Gender {
	g, err := c.First()
	if err != nil {
		panic(err)
	}

	return g
}

// CollectionEachCallback iteration callback.
type CollectionEachCallback func(g *Gender)

// Each iterate over collection.
func (c *Collection) Each(fn CollectionEachCallback) error {
	if c.Length() == 0 {
		return ErrNothingFound
	}

	for _, g := range c.genders {
		fn(g)
	}

	return nil
}

// EachX like Each, but panics when error.
func (c *Collection) EachX(fn CollectionEachCallback) {
	if err := c.Each(fn); err != nil {
		panic(err)
	}
}
