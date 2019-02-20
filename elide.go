package main

import "unicode/utf8"

const (
	VarcharLimit = 500
	ThreeDotsElision  = "..."
	FreeTextMaxLength = VarcharLimit - 3 // make room for three dots
)

func Elide(field string) string {

	if len(field) > FreeTextMaxLength {

		// Convert to byte array since it's required by the utf8 func
		// Also it's always true that len(bytes) >= len(chars)
		bytes := []byte(field)

		for len(bytes) > FreeTextMaxLength {
			_, size := utf8.DecodeLastRune(bytes)
			bytes = bytes[:len(bytes)-size]
		}
		return string(bytes) + ThreeDotsElision
	}
	return field
}