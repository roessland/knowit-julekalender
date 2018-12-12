package main

import "fmt"

func main() {
	emojis := "😡😚😀😷😨😥😮😀😩😀😤😩😥😌😀😩😀😷😡😮😮😡😀😤😩😥😀😬😩😫😥😀😣😡😥😳😡😲😎😀😱😚😀😨😯😷😀😣😯😭😥😟😀😡😚😀😨😥😀😤😩😥😤😀😡😭😯😮😧😀😨😩😳😀😦😲😩😥😮😤😳😎"
	fmt.Println(emojis)

	for off := 0; off < 70; off++ {
		for _, r := range emojis {
			fmt.Printf("%c", r-128540+28+rune(off))
		}
		fmt.Println()
	}
}
