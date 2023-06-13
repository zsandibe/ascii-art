package banner

import (
	"fmt"
	"os"
	"strings"
)

func ReadBanner(banner string) (map[rune][]string, error) {
	input, err := os.ReadFile(banner)
	if err != nil {
		return nil, fmt.Errorf("in standart tead : %v", err)
	}
	standartInput := strings.ReplaceAll(string(input), "\r", "")
	standartInput1 := strings.Split(string(standartInput), "\n")
	s := make(map[rune][]string)
	startingPoint := ' '
	var val []string
	for i, ch := range standartInput1 {
		i++
		if i%9 == 0 {
			val = append(val, ch)
			s[startingPoint] = val
			startingPoint++
			val = nil
		} else if ch != "" {
			val = append(val, ch)
		}
	}
	return s, nil
}

func PrintBanner(text string, banner map[rune][]string) {
	if len(text) < 1 {
		return
	} else if len(text) > 52 {
		fmt.Println("Dude,don't exceed the character limit")
		return
	}
	line := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")

	for i, w := range line {
		if len(w) < 1 {
			if i != len(line)-1 {
				fmt.Println()
			} else {
				return
			}
			continue
		} else {
			for i := 0; i < 8; i++ {
				for _, char := range w {
					fmt.Print(banner[char][i])
				}
				fmt.Println()
			}
			if i != len(line)-1 {
				fmt.Println()
			}
		}
	}
}
