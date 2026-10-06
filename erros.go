package main

import "errors"

var ErrPileNotSorted = errors.New("pile not sorted")
var ErrValueNotFound = errors.New("value not found")
var ErrEmptyPile = errors.New("empty pile")