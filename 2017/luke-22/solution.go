package main

import "fmt"

func containsT(count map[byte]int) bool {
	return count['A'] > 1 && count['B'] > 0 && count['C'] > 0 && count['D'] > 0
}

func main() {
	S := "FJKAUNOJDCUTCRHBYDLXKEODVBWTYPTSHASQQFCPRMLDXIJMYPVOHBDUGSMBLMVUMMZYHULSUIZIMZTICQORLNTOVKVAMQTKHVRIFMNTSLYGHEHFAHWWATLYAPEXTHEPKJUGDVWUDDPRQLUZMSZOJPSIKAIHLTONYXAULECXXKWFQOIKELWOHRVRUCXIAASKHMWTMAJEWGEESLWRTQKVHRRCDYXNTLDSUPXMQTQDFAQAPYBGXPOLOCLFQNGNKPKOBHZWHRXAWAWJKMTJSLDLNHMUGVVOPSAMRUJEYUOBPFNEHPZZCLPNZKWMTCXERPZRFKSXVEZTYCXFRHRGEITWHRRYPWSVAYBUHCERJXDCYAVICPTNBGIODLYLMEYLISEYNXNMCDPJJRCTLYNFMJZQNCLAGHUDVLYIGASGXSZYPZKLAWQUDVNTWGFFYFFSMQWUNUPZRJMTHACFELGHDZEJWFDWVPYOZEVEJKQWHQAHOCIYWGVLPSHFESCGEUCJGYLGDWPIWIDWZZXRUFXERABQJOXZALQOCSAYBRHXQQGUDADYSORTYZQPWGMBLNAQOFODSNXSZFURUNPMZGHTAJUJROIGMRKIZHSFUSKIZJJTLGOEEPBMIXISDHOAIFNFEKKSLEXSJLSGLCYYFEQBKIZZTQQXBQZAPXAAIFQEIXELQEZGFEPCKFPGXULLAHXTSRXDEMKFKABUTAABSLNQBNMXNEPODPGAORYJXCHCGKECLJVRBPRLHORREEIZOBSHDSCETTTNFTSMQPQIJBLKNZDMXOTRBNMTKHHCZQQMSLOAXJQKRHDGZVGITHYGVDXRTVBJEAHYBYRYKJAVXPOKHFFMEPHAGFOOPFNKQAUGYLVPWUJUPCUGGIXGRAMELUTEPYILBIUOCKKUUBJROQFTXMZRLXBAMHSDTEKRRIKZUFNLGTQAEUINMBPYTWXULQNIIRXHHGQDPENXAJNWXULFBNKBRINUMTRBFWBYVNKNKDFR"
	shortestLen := 999
	for startPos, _ := range S {
		count := map[byte]int{}
		pos := startPos
		for pos != len(S)-1 {
			count[S[pos]]++
			if containsT(count) {
				length := pos - startPos + 1
				if length <= shortestLen {
					fmt.Println(S[startPos:pos+1], pos-startPos+1)
					shortestLen = length
				}
				break
			}
			pos++
		}
	}
}
