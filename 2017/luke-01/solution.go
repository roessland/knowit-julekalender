package main

import "fmt"
import "strings"
import "os"
import "bufio"
import "sort"
import "log"

func Ngramify(n int, s string) string {
	if len(s) <= n {
		return ""
	}
	ngrams := make([]string, len(s)+1-n)
	for i := 0; i < len(ngrams); i++ {
		ngrams[i] = s[i : i+n]
	}
	return strings.Join(ngrams, "")
}

func Sorted(s string) string {
	rs := []rune(s)
	sort.Slice(rs, func(i, j int) bool { return rs[i] < rs[j] })
	return string(rs)
}

func main() {
	target := Sorted("aeteesasrsssstaesersrrsse")
	f, err := os.Open("wordlist.txt")
	if err != nil {
		log.Fatal("Couldn't open wordlist.txt", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		for i := 2; i < 7; i++ {
			if Sorted(Ngramify(i, strings.ToLower(scanner.Text()))) == target {
				fmt.Printf("%d-%s\n", i, scanner.Text())
			}
		}
	}
}
