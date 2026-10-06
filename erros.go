package main

import "errors"

var ErrPileNotSorted = errors.New("pile not sorted")
var ErrValueNotFound = errors.New("value not found")
var ErrEmptyPile = errors.New("empty pile")
var ErrNegativeHeight = errors.New("la hauteur ne peut pas être négative")
var ErrOverflow = errors.New("le résultat dépasse la capacité d'un int")
var ErrInvalid = errors.New("pile invalide")