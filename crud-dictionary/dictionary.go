package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrUpdatingNonExistentWord = DictionaryErr("cannot update the value of the word which does not exist")
	ErrDeletingNonExistingWord = DictionaryErr("the word you are attempting to delete does not exist")
)

func (d Dictionary) Search(key string) (string, error) {
	val, isPresent := d[key]

	if !isPresent {
		return "", ErrNotFound
	}

	return val, nil
}

func (d Dictionary) Add(key, value string) error {
	_, isPresent := d[key]
	if !isPresent {
		d[key] = value
		return nil
	}
	return ErrWordExists
}

func (d Dictionary) Update(key, value string) error {
	_, isPresent := d[key]
	if !isPresent {
		return ErrUpdatingNonExistentWord
	}
	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, isPresent := d[key]
	if !isPresent {
		return ErrDeletingNonExistingWord
	}
	delete(d, key)
	return nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}