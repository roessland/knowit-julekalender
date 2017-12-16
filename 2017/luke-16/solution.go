package main

import "fmt"
import "strconv"
import "bufio"
import "os"

func main() {
	on := false
	hasTurnedOn := make([]bool, 101)
	turnedOffCount := 0
	scanner := bufio.NewScanner(os.Stdin)
	timesInRoom := 0
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		timesInRoom++
		if n == 1 {
			if on {
				on = false
				turnedOffCount++
			}
			if turnedOffCount == 99 {
				fmt.Println("We did it!")
				fmt.Println("People have been in the room", timesInRoom, "times")
				fmt.Println(hasTurnedOn)
				break
			}
		} else {
			if !on && !hasTurnedOn[n] {
				on = true
				hasTurnedOn[n] = true
			}
		}
	}
}
