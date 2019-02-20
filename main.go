package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main() {

	StringsInGo()
	RunesInGo()
	NaiveSubstring()
	Elide("some text")
}

func StringsInGo() {
	s := "Nulab" // s is assigned a string literal

	fmt.Println(reflect.TypeOf(s[0])) // uint8
	fmt.Printf("%x\n", s[0]) // 4e
	fmt.Printf("%d\n", s[0]) // 78
	fmt.Printf("%c\n", s[0]) // N

	bytes := []byte{0x4e, 0x75, 0x6c, 0x61, 0x62}
	fmt.Printf("%s\n", bytes) // Nulab
}

func RunesInGo() {
	var r rune = '乾' // U+4e7e
	fmt.Println(reflect.TypeOf(r)) // int32
	fmt.Printf("%x\n", r) // 4e7e

	var wide rune = 0x1f76a
	fmt.Printf("%c\n", wide)
	fmt.Println("Length: ", len(string(wide)))

	// var overflow rune = 0x1f76a0000
}

func NaiveSubstring() {
	const max = 5

	a := "截断错误"
	// UTF-8
	// e6 88 aa (截) e6 96 ad (断) e9 94 99 (错) e8 af af (误)

	b := a[:max]
	fmt.Println("Bad cut: %s", b) // 截�
}


const (
	VarcharLimit = 100
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


