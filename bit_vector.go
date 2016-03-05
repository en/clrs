package clrs

// [0, 15]
type bitVector uint16

func (b *bitVector) Insert(key int) error {
	if b.Search(key) {
		return errExist
	}
	*b |= 1 << uint(key)
	return nil
}

func (b *bitVector) Search(key int) bool {
	return (*b & (1 << uint(key))) != 0
}

func (b *bitVector) Delete(key int) error {
	if !b.Search(key) {
		return errNotExist
	}
	*b &= ^(1 << uint(key))
	return nil
}
