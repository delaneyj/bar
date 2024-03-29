package entity

type Ref struct {
	id         uint64
	generation uint64
	components [3]uint64
}

func (r Ref) Is(other Ref) bool {
    return r.id == other.id && r.generation == other.generation
}

func (r Ref) Zero() bool {
    return r.id == 0 && r.generation == 0
}