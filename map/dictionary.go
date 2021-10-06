package main

import (
	"errors"
)

type Dictionary map[string]string

var (
	ErrNotFound     = errors.New("could not find the word you were looking for")
	ErrWordExists   = errors.New("cannot add word because it alredy exists")
	ErrWordNotExist = errors.New("cannot update definition because word is not in dictionary")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, value string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
