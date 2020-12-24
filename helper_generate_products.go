package main

import (
	"math/rand"
	"strconv"
	"strings"
)

func getRandomWord() string {
	wordsLatin := []string{
		"ipsum", "semper", "habeo", "duo", "ut", "vis", "aliquyam", "eu", "splendide", "Ut", "mei", "eteu", "nec", "antiopam", "corpora", "kasd", "pretium", "cetero", "qui", "arcu", "assentior", "ei", "his", "usu", "invidunt", "kasd", "justo", "ne", "eleifend", "per", "ut", "eam", "graeci", "tincidunt", "impedit", "temporibus", "duo", "et", "facilisis", "insolens", "consequat", "cursus", "partiendo", "ullamcorper", "Vulputate", "facilisi", "donec", "aliquam", "labore", "inimicus", "voluptua", "penatibus", "sea", "vel", "amet", "his", "ius", "audire", "in", "mea", "repudiandae", "nullam", "sed", "assentior", "takimata", "eos", "at", "odio", "consequat", "iusto", "imperdiet", "dicunt", "abhorreant", "adipisci", "officiis", "rhoncus", "leo", "dicta", "vitae", "clita", "elementum", "mauris", "definiebas", "uonsetetur", "te", "inimicus", "nec", "mus", "usu", "duo", "aenean", "corrumpit", "aliquyam", "est", "eum",
	}

	return wordsLatin[rand.Intn(len(wordsLatin))]
}

func generateWords(length int) string {
	result := "Lorem "

	for i := 0; i < length-1; i++ {
		result = result + getRandomWord() + " "
	}

	return result
}

func generateSequenceInt(begin, no uint) []uint {
	result := make([]uint, no)
	for i := 0; i < int(no); i, begin = i+1, begin+1 {
		result[i] = begin
	}
	return result
}

func generatorProducts(no uint) []*Product {
	rangeID := generateSequenceInt(1, no)

	result := make([]*Product, int(no))
	for i, id := range rangeID {
		result[i] = &Product{
			ID:              id,
			SKU:             "SKU" + strconv.FormatUint(id, 10),
			Category:        "cosmetics",
			Name:            generateWords(2),
			Slug:            strings.Join(strings.Split(Name, " "), "-"),
			Description:     generateWords(10 + id),
			Quantity:        id,
			PriceCents:      id * 100,
			SalesPriceCents: id*100 - id,
		}

	}

	return result
}
