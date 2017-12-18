package main

import "fmt"
import "os"
import "strconv"
import "strings"
import "log"
import "bufio"
import "io"
import "sort"
import "unicode"

var IcelandicAlphabetEncoding map[rune]int

func init() {
	IcelandicAlphabetEncoding = make(map[rune]int)
	i := 0
	for _, r := range "AÁBDÐEÉFGHIÍJKLMNOÓPRSTUÚVXYÝÞÆÖ" {
		IcelandicAlphabetEncoding[r] = i
		i++
	}
}

func main() {
	// Compute frequencies
	reader := bufio.NewReader(os.Stdin)
	freqs := map[rune]int{}
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if unicode.IsSpace(r) {
			continue
		}
		freqs[r]++
	}

	// Sort unique characters by descending frequency
	key := make([]rune, 0, len(freqs))
	for r, _ := range freqs {
		key = append(key, r)
	}
	sort.Slice(key, func(i, j int) bool {
		return freqs[key[i]] > freqs[key[j]]
	})

	// Convert from utf8 to the Magic Icelandic Knowit Encoding (MIKE) (tm)
	stringbuilder := []string{}
	for _, r := range key {
		stringbuilder = append(stringbuilder, fmt.Sprintf("%05b", IcelandicAlphabetEncoding[r]))
	}
	k := strings.Join(stringbuilder, "")

	y := "1110010101000001011000000011101110100101010011011010101101100000010001111101000001010010001011101001100100100011010000110101111101010011100010110001100111110010"

	// Take first 8 bits of key and encrypted string, xor them together.
	for i := 0; i < len(y)/8; i++ {
		yc := y[8*i : 8*i+8]
		kc := k[8*i : 8*i+8]
		yb, errY := strconv.ParseUint(yc, 2, 32)
		kb, errK := strconv.ParseUint(kc, 2, 32)
		if errY != nil || errK != nil {
			log.Fatal(errY, errK)
		}
		fmt.Printf("%c", byte(yb)^byte(kb))
	}
}
