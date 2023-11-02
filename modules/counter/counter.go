package counter

import "github.com/sirupsen/logrus"

type Counter struct {
	count int64
}

func Init() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() *Counter {
	c.count++
	return c
}

func (c *Counter) Print(l *logrus.Logger) {
	l.Debugf("count: %d", c.count)
}
