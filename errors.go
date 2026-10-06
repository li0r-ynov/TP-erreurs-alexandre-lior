package main

import "errors"

var ErrEmptyPile = errors.New("Pile vide")
var ErrSinglePile = errors.New("Pile unique")
