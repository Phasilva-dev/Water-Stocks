package misc

import()

type Tuple struct {
	key   string
	value float64
}

func NewTuple(k string, v float64) *Tuple{
	return &Tuple{
		key: k,
		value: v,
	}
}

func (t *Tuple) Key() string{
	return t.key
}

func (t *Tuple) Value() float64{
	return t.value
}