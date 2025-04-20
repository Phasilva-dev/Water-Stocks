package misc

type Tuple[K comparable, V any] struct {
    key   K
    value V
}

func NewTuple[K comparable, V any](k K, v V) *Tuple[K, V] {
    return &Tuple[K, V]{
        key:   k,
        value: v,
    }
}

func (t *Tuple[K, V]) Key() K {
    return t.key
}

func (t *Tuple[K, V]) Value() V {
    return t.value
}