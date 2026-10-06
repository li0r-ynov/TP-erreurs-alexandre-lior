package main

import "errors"

var ErrEmptyPile = errors.New("Pile vide")
var ErrNegativePile = errors.New("Pile Négative")