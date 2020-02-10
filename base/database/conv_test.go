package database_test

import (
	"his6/base/database"
	"testing"
)

func BenchmarkColx(b *testing.B) {
	s := "QsdwleiWnUlzkcn1D1OlkjcwijvCkwej23lXkwdjfh"
	for n := 0; n < b.N; n++ {
		database.PascalToColx(s)
	}
}

func BenchmarkColy(b *testing.B) {
	s := "QsdwleiWnUlzkcn1D1OlkjcwijvCkwej23lXkwdjfh"
	for n := 0; n < b.N; n++ {
		database.PascalToCol(s)
	}
}
