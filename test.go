package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
)

func main() {
	width, _ := getTerminal()
	if len(os.Args) != 4 {
		fmt.Println(`
Usage: go run . [OPTION] [STRING] [BANNER]
Example: go run . --align=<center> something standard`)
		return
	}
	file := flag.String("align", "left", "where to print")
	flag.Parse()
	str := os.Args[2]
	split := strings.Split(str, "\\n")
	switch os.Args[3] {
	case "standard":
		if str == "\\n" {
			fmt.Print("\n")
			return
		}
		for _, word := range split {
			if word == "" {
				fmt.Print("\n")
			} else {
				if *file == "right" {
					for k := 0; k < 8; k++ {
						printInRight12(word, width)
						for i := 0; i < len(word); i++ {
							fmt.Print(standard2(rune(word[i]))[k])
						}
						fmt.Print("\n")
					}
				} else if *file == "center" {
					for k := 0; k < 8; k++ {
						printInMid12(word, width)
						for i := 0; i < len(word); i++ {
							fmt.Print(standard2(rune(word[i]))[k])
						}
						fmt.Print("\n")
					}
				} else if *file == "left" {
					for k := 0; k < 8; k++ {
						for i := 0; i < len(word); i++ {
							fmt.Print(standard2(rune(word[i]))[k])
						}
						fmt.Print("\n")
					}
				} else if *file == "justify" {
					arr := strings.Fields(word)
					for k := 0; k < 8; k++ {
						for index := 0; index < len(arr); index++ {
							for i := 0; i < len(arr[index]); i++ {
								if index != 0 {
									if i == 0 {
										printInJustify12(word, width)
									}
									fmt.Print(standard2(rune(arr[index][i]))[k])
								} else {
									fmt.Print(standard2(rune(arr[index][i]))[k])
								}
							}
						}
						fmt.Print("\n")
					}
				} else {
					fmt.Print("enter a valid value: right, left, center, justify")
				}
			}
		}
		/*else if *file == "justify" {
			arr := strings.Fields(word)
			for index := range arr {
				for k := 0; k < 8; k++ {
					if index != 0 {
						printInJustify12(word, width)
					}
					for i := 0; i < len(word); i++ {
						fmt.Print(standard2(rune(word[i]))[k])
					}
					fmt.Print("\n")
				}
			}
		} else {
			fmt.Print("enter a valid value: right, left, center, justify")
		}*/
	case "shadow":
		if str == "\\n" {
			fmt.Print("\n")
			return
		}
		for _, word := range split {
			if word == "" {
				fmt.Print("\n")
			} else {
				if *file == "right" {
					for k := 0; k < 8; k++ {
						printInRight12(word, width)
						for i := 0; i < len(word); i++ {
							fmt.Print(shadow2(rune(word[i]))[k])
						}
						fmt.Print("\n")
					}
				} else if *file == "center" {
					for k := 0; k < 8; k++ {
						printInMid12(word, width)
						for i := 0; i < len(word); i++ {
							fmt.Print(shadow2(rune(word[i]))[k])
						}
						fmt.Print("\n")
					}
				} else if *file == "left" {
					for k := 0; k < 8; k++ {
						for i := 0; i < len(word); i++ {
							fmt.Print(shadow2(rune(word[i]))[k])
						}
						fmt.Print("\n")
					}
				} else if *file == "justify" {
					arr := strings.Fields(word)
					for k := 0; k < 8; k++ {
						for index := 0; index < len(arr); index++ {
							for i := 0; i < len(arr[index]); i++ {
								if index != 0 {
									if i == 0 {
										printInJustify12(word, width)
									}
									fmt.Print(shadow2(rune(arr[index][i]))[k])
								} else {
									fmt.Print(shadow2(rune(arr[index][i]))[k])
								}
							}
						}
						fmt.Print("\n")
					}
				} else {
					fmt.Print("enter a valid value: right, left, center, justify")
				}
			}
		}
	case "thinkertoy":
		if str == "\\n" {
			fmt.Print("\n")
			return
		}
		for _, word := range split {
			if word == "" {
				fmt.Print("\n")
			} else {
				if *file == "right" {
					for k := 0; k < 8; k++ {
						printInRight12(word, width)
						for i := 0; i < len(word); i++ {
							fmt.Print(thinkertoy2(rune(word[i]))[k])
						}
						fmt.Print("\n")
					}
				} else if *file == "center" {
					for k := 0; k < 8; k++ {
						printInMid12(word, width)
						for i := 0; i < len(word); i++ {
							fmt.Print(thinkertoy2(rune(word[i]))[k])
						}
						fmt.Print("\n")
					}
				} else if *file == "left" {
					for k := 0; k < 8; k++ {
						for i := 0; i < len(word); i++ {
							fmt.Print(thinkertoy2(rune(word[i]))[k])
						}
						fmt.Print("\n")
					}
				} else if *file == "justify" {
					arr := strings.Fields(word)
					for k := 0; k < 8; k++ {
						for index := 0; index < len(arr); index++ {
							for i := 0; i < len(arr[index]); i++ {
								if index != 0 {
									if i == 0 {
										printInJustify12(word, width)
									}
									fmt.Print(thinkertoy2(rune(arr[index][i]))[k])
								} else {
									fmt.Print(thinkertoy2(rune(arr[index][i]))[k])
								}
							}
						}
						fmt.Print("\n")
					}
				} else {
					fmt.Print("enter a valid value: right, left, center, justify")
				}
			}
		}
	default:
		fmt.Println(`
Usage: go run . [STRING] [BANNER]
EX: go run . something standard2`)
	}
}
func get() int {
	columnsStr := os.Getenv("COLUMNS")
	columns, _ := strconv.Atoi(columnsStr)
	return columns
}
func getTerminal() (int, error) {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		width, _, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			return 0, err
		}
		return width, nil
	}
	// If not a terminal, try getting the width from the environment variable
	columnsStr := os.Getenv("COLUMNS")
	columns, err := strconv.Atoi(columnsStr)
	if err != nil {
		return 0, err
	}
	return columns, nil
}
func printInMid12(s string, width int) {
	spaces := (width - (len(s) * 7)) / 2
	if spaces < 0 {
		spaces = 0
	}
	fmt.Print(strings.Repeat(" ", spaces))
}
func printInRight12(s string, width int) {
	spaces := (width - (len(s) * 8))
	if spaces < 0 {
		spaces = 0
	}
	fmt.Print(strings.Repeat(" ", spaces))
}

/*
	func printInJustify12(s string, width int) {
		arr := strings.Fields(s)
		spaces := (width - (len(s) * 7)) / (len(arr) - 1)
		for i := range arr {
			if i != 0 {
		fmt.Print(strings.Repeat(" ", spaces))
			}
		}
	}
*/
func printInJustify12(s string, width int) {
	arr := strings.Fields(s)
	spaces := (width - (len(s) * 7)) / (len(arr) - 1)
	if spaces < 0 {
		spaces = 0
	}
	fmt.Print(strings.Repeat(" ", spaces))
}
func standard2(r rune) []string {
	allLetters := make(map[rune][]string)
	allLetters[' '] = []string{
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
	}
	allLetters['!'] = []string{
		" _  ",
		"| | ",
		"| | ",
		"| | ",
		"|_| ",
		"(_) ",
		"    ",
		"    ",
	}
	allLetters['"'] = []string{
		" _ _  ",
		"( | ) ",
		" V V  ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	allLetters['#'] = []string{
		"     _  _    ",
		"   _| || |_  ",
		"  |_  __  _| ",
		"   _| || |_  ",
		"  |_  __  _| ",
		"    |_||_|   ",
		"             ",
		"             ",
	}
	allLetters['$'] = []string{
		"    _   ",
		"   | |  ",
		"  / __) ",
		"  \\__ \\ ",
		"  (   / ",
		"   |_|  ",
		"        ",
		"        ",
	}
	allLetters['%'] = []string{
		" _   __ ",
		"(_) / / ",
		"   / /  ",
		"  / /   ",
		" / / _  ",
		"/_/ (_) ",
		"        ",
		"        ",
	}
	allLetters['&'] = []string{
		"           ",
		"    ___    ",
		"   ( _ )   ",
		"   / _ \\/\\ ",
		"  | (_>  < ",
		"   \\___/\\/ ",
		"           ",
		"           ",
	}
	allLetters['\''] = []string{
		" _  ",
		"( ) ",
		"|/  ",
		"    ",
		"    ",
		"    ",
		"    ",
		"    ",
	}
	allLetters['('] = []string{
		"    __ ",
		"   / / ",
		"  | |  ",
		"  | |  ",
		"  | |  ",
		"  | |  ",
		"   \\_\\ ",
		"       ",
	}
	allLetters[')'] = []string{
		"__   ",
		"\\ \\  ",
		" | | ",
		" | | ",
		" | | ",
		" | | ",
		"/_/  ",
		"     ",
	}
	allLetters['*'] = []string{
		"    _     ",
		" /\\| |/\\  ",
		" \\ ` ' /  ",
		"|_     _| ",
		" / , . \\  ",
		" \\/|_|\\/  ",
		"          ",
		"          ",
	}
	allLetters['+'] = []string{
		"        ",
		"   _    ",
		" _| |_  ",
		"|_   _| ",
		"  |_|   ",
		"        ",
		"        ",
		"        ",
	}
	allLetters['-'] = []string{
		"         ",
		"         ",
		" ______  ",
		"|______| ",
		"         ",
		"         ",
		"         ",
		"         ",
	}
	allLetters['.'] = []string{
		"    ",
		"    ",
		"    ",
		"    ",
		" _  ",
		"(_) ",
		"    ",
		"    ",
	}
	allLetters['/'] = []string{
		"     __ ",
		"    / / ",
		"   / /  ",
		"  / /   ",
		" / /    ",
		"/_/     ",
		"        ",
		"        ",
	}
	allLetters['0'] = []string{
		"        ",
		"  ___   ",
		" / _ \\  ",
		"| | | | ",
		"| |_| | ",
		" \\___/  ",
		"        ",
		"        ",
	}
	allLetters['1'] = []string{
		"    ",
		" _  ",
		"/ | ",
		"| | ",
		"| | ",
		"|_| ",
		"    ",
		"    ",
	}
	allLetters['2'] = []string{
		"        ",
		" ____   ",
		"|___ \\  ",
		"  __) | ",
		" / __/  ",
		"|_____| ",
		"        ",
		"        ",
	}
	allLetters['3'] = []string{
		"        ",
		" _____  ",
		"|___ /  ",
		"  |_ \\  ",
		" ___) | ",
		"|____/  ",
		"        ",
		"        ",
	}
	allLetters['4'] = []string{
		"         ",
		" _  _    ",
		"| || |   ",
		"| || |_  ",
		"|__   _| ",
		"   |_|   ",
		"         ",
		"         ",
	}
	allLetters['5'] = []string{
		"        ",
		" ____   ",
		"| ___|  ",
		"|___ \\  ",
		"  __) | ",
		"|____/  ",
		"        ",
		"        ",
	}
	allLetters['6'] = []string{
		"        ",
		"  __    ",
		" / /    ",
		"| '_ \\  ",
		"| (_) | ",
		" \\___/  ",
		"        ",
		"        ",
	}
	allLetters['7'] = []string{
		"        ",
		" _____  ",
		"|___  | ",
		"   / /  ",
		"  / /   ",
		" /_/    ",
		"        ",
		"        ",
	}
	allLetters['8'] = []string{
		"          ",
		"    ___   ",
		"   ( _ )  ",
		"   / _ \\  ",
		"  | (_) | ",
		"   \\___/  ",
		"          ",
		"          ",
	}
	allLetters['9'] = []string{
		"    ___   ",
		"   / _ \\  ",
		"  | (_) | ",
		"   \\__, | ",
		"     / /  ",
		"    /_/   ",
		"          ",
		"          ",
	}
	allLetters[':'] = []string{
		" _  ",
		"(_) ",
		"    ",
		" _  ",
		"(_) ",
		"    ",
		"    ",
		"    ",
	}
	allLetters[';'] = []string{
		" _  ",
		"(_) ",
		"    ",
		" _  ",
		"( ) ",
		"|/  ",
		"    ",
		"    ",
	}
	allLetters['<'] = []string{
		"   __ ",
		"  / / ",
		" / /  ",
		"< <   ",
		" \\ \\  ",
		"  \\_\\ ",
		"      ",
		"      ",
	}
	allLetters['='] = []string{
		"         ",
		" ______  ",
		"|______| ",
		" ______  ",
		"|______| ",
		"         ",
		"         ",
		"         ",
	}
	allLetters['>'] = []string{
		"__    ",
		"\\ \\   ",
		" \\ \\  ",
		"  > > ",
		" / /  ",
		"/_/   ",
		"      ",
		"      ",
	}
	allLetters['?'] = []string{
		" ___   ",
		"|__ \\  ",
		"   ) | ",
		"  / /  ",
		" |_|   ",
		" (_)   ",
		"       ",
		"       ",
	}
	allLetters['@'] = []string{
		"          ",
		"   ____   ",
		"  / __ \\  ",
		" / / _` | ",
		"| | (_| | ",
		" \\ \\__,_| ",
		"  \\____/  ",
		"          ",
	}
	allLetters['A'] = []string{
		"           ",
		"    /\\     ",
		"   /  \\    ",
		"  / /\\ \\   ",
		" / ____ \\  ",
		"/_/    \\_\\ ",
		"           ",
		"           ",
	}
	allLetters['B'] = []string{
		" ____   ",
		"|  _ \\  ",
		"| |_) | ",
		"|  _ <  ",
		"| |_) | ",
		"|____/  ",
		"        ",
		"        ",
	}
	allLetters['C'] = []string{
		"  _____  ",
		" / ____| ",
		"| |      ",
		"| |      ",
		"| |____  ",
		" \\_____| ",
		"         ",
		"         ",
	}
	allLetters['D'] = []string{
		" _____   ",
		"|  __ \\  ",
		"| |  | | ",
		"| |  | | ",
		"| |__| | ",
		"|_____/  ",
		"         ",
		"         ",
	}
	allLetters['E'] = []string{
		" ______  ",
		"|  ____| ",
		"| |__    ",
		"|  __|   ",
		"| |____  ",
		"|______| ",
		"         ",
		"         ",
	}
	allLetters['F'] = []string{
		" ______  ",
		"|  ____| ",
		"| |__    ",
		"|  __|   ",
		"| |      ",
		"|_|      ",
		"         ",
		"         ",
	}
	allLetters['G'] = []string{
		"  _____  ",
		" / ____| ",
		"| |  __  ",
		"| | |_ | ",
		"| |__| | ",
		" \\_____| ",
		"         ",
		"         ",
	}
	allLetters['H'] = []string{
		" _    _  ",
		"| |  | | ",
		"| |__| | ",
		"|  __  | ",
		"| |  | | ",
		"|_|  |_| ",
		"         ",
		"         ",
	}
	allLetters['I'] = []string{
		" _____  ",
		"|_   _| ",
		"  | |   ",
		"  | |   ",
		" _| |_  ",
		"|_____| ",
		"        ",
		"        ",
	}
	allLetters['J'] = []string{
		"      _  ",
		"     | | ",
		"     | | ",
		" _   | | ",
		"| |__| | ",
		" \\____/  ",
		"         ",
		"         ",
	}
	allLetters['K'] = []string{
		" _  __ ",
		"| |/ / ",
		"| ' /  ",
		"|  <   ",
		"| . \\  ",
		"|_|\\_\\ ",
		"       ",
		"       ",
	}
	allLetters['L'] = []string{
		" _       ",
		"| |      ",
		"| |      ",
		"| |      ",
		"| |____  ",
		"|______| ",
		"         ",
		"         ",
	}
	allLetters['M'] = []string{
		" __  __  ",
		"|  \\/  | ",
		"| \\  / | ",
		"| |\\/| | ",
		"| |  | | ",
		"|_|  |_| ",
		"         ",
		"         ",
	}
	allLetters['N'] = []string{
		" _   _  ",
		"| \\ | | ",
		"|  \\| | ",
		"| . ` | ",
		"| |\\  | ",
		"|_| \\_| ",
		"        ",
		"        ",
	}
	allLetters['O'] = []string{
		"  ____   ",
		" / __ \\  ",
		"| |  | | ",
		"| |  | | ",
		"| |__| | ",
		" \\____/  ",
		"         ",
		"         ",
	}
	allLetters['P'] = []string{
		" _____   ",
		"|  __ \\  ",
		"| |__) | ",
		"|  ___/  ",
		"| |      ",
		"|_|      ",
		"         ",
		"         ",
	}
	allLetters['Q'] = []string{
		"  ____   ",
		" / __ \\  ",
		"| |  | | ",
		"| |  | | ",
		"| |__| | ",
		" \\___\\_\\ ",
		"         ",
		"         ",
	}
	allLetters['R'] = []string{
		" _____   ",
		"|  __ \\  ",
		"| |__) | ",
		"|  _  /  ",
		"| | \\ \\  ",
		"|_|  \\_\\ ",
		"         ",
		"         ",
	}
	allLetters['S'] = []string{
		"  _____  ",
		" / ____| ",
		"| (___   ",
		" \\___ \\  ",
		" ____) | ",
		"|_____/  ",
		"         ",
		"         ",
	}
	allLetters['T'] = []string{
		" _______  ",
		"|__   __| ",
		"   | |    ",
		"   | |    ",
		"   | |    ",
		"   |_|    ",
		"          ",
		"          ",
	}
	allLetters['U'] = []string{
		" _    _  ",
		"| |  | | ",
		"| |  | | ",
		"| |  | | ",
		"| |__| | ",
		" \\____/  ",
		"         ",
		"         ",
	}
	allLetters['V'] = []string{
		"__      __ ",
		"\\ \\    / / ",
		" \\ \\  / /  ",
		"  \\ \\/ /   ",
		"   \\  /    ",
		"    \\/     ",
		"           ",
		"           ",
	}
	allLetters['W'] = []string{
		"__          __ ",
		"\\ \\        / / ",
		" \\ \\  /\\  / /  ",
		"  \\ \\/  \\/ /   ",
		"   \\  /\\  /    ",
		"    \\/  \\/     ",
		"               ",
		"               ",
	}
	allLetters['X'] = []string{
		"__   __ ",
		"\\ \\ / / ",
		" \\ V /  ",
		"  > <   ",
		" / . \\  ",
		"/_/ \\_\\ ",
		"        ",
		"        ",
	}
	allLetters['Y'] = []string{
		"__     __ ",
		"\\ \\   / / ",
		" \\ \\_/ /  ",
		"  \\   /   ",
		"   | |    ",
		"   |_|    ",
		"          ",
		"          ",
	}
	allLetters['Z'] = []string{
		" ______ ",
		"|___  / ",
		"   / /  ",
		"  / /   ",
		" / /__  ",
		"/_____| ",
		"        ",
		"        ",
	}
	allLetters['['] = []string{
		" ___  ",
		"|  _| ",
		"| |   ",
		"| |   ",
		"| |   ",
		"| |_  ",
		"|___| ",
		"      ",
	}
	allLetters[']'] = []string{
		" ___  ",
		"|_  | ",
		"  | | ",
		"  | | ",
		"  | | ",
		" _| | ",
		"|___| ",
		"      ",
	}
	allLetters['\\'] = []string{
		"__      ",
		"\\ \\     ",
		" \\ \\    ",
		"  \\ \\   ",
		"   \\ \\  ",
		"    \\_\\ ",
		"        ",
		"        ",
	}
	allLetters['^'] = []string{
		" /\\   ",
		"|/\\|  ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	allLetters['_'] = []string{
		"         ",
		"         ",
		"         ",
		"         ",
		"         ",
		"         ",
		" ______  ",
		"|______| ",
	}
	allLetters[','] = []string{
		"    ",
		"    ",
		"    ",
		"    ",
		" _  ",
		"( ) ",
		"|/  ",
		"    ",
	}
	allLetters['`'] = []string{
		" _  ",
		"( ) ",
		" \\| ",
		"    ",
		"    ",
		"    ",
		"    ",
		"    ",
	}
	allLetters['a'] = []string{
		"        ",
		"        ",
		"  __ _  ",
		" / _` | ",
		"| (_| | ",
		" \\__,_| ",
		"        ",
		"        ",
	}
	allLetters['b'] = []string{
		" _      ",
		"| |     ",
		"| |__   ",
		"| '_ \\  ",
		"| |_) | ",
		"|_.__/  ",
		"        ",
		"        ",
	}
	allLetters['c'] = []string{
		"       ",
		"       ",
		"  ___  ",
		" / __| ",
		"| (__  ",
		" \\___| ",
		"       ",
		"       ",
	}
	allLetters['e'] = []string{
		"       ",
		"       ",
		"  ___  ",
		" / _ \\ ",
		"|  __/ ",
		" \\___| ",
		"       ",
		"       ",
	}
	allLetters['f'] = []string{
		"  __  ",
		" / _| ",
		"| |_  ",
		"|  _| ",
		"| |   ",
		"|_|   ",
		"      ",
		"      ",
	}
	allLetters['g'] = []string{
		"        ",
		"        ",
		"  __ _  ",
		" / _` | ",
		"| (_| | ",
		" \\__, | ",
		"  __/ | ",
		" |___/  ",
	}
	allLetters['d'] = []string{
		"     _  ",
		"    | | ",
		"  __| | ",
		" / _` | ",
		"| (_| | ",
		" \\__,_| ",
		"        ",
		"        ",
	}
	allLetters['h'] = []string{
		" _      ",
		"| |     ",
		"| |__   ",
		"|  _ \\  ",
		"| | | | ",
		"|_| |_| ",
		"        ",
		"        ",
	}
	allLetters['i'] = []string{
		" _  ",
		"(_) ",
		" _  ",
		"| | ",
		"| | ",
		"|_| ",
		"    ",
		"    ",
	}
	allLetters['j'] = []string{
		"   _  ",
		"  (_) ",
		"   _  ",
		"  | | ",
		"  | | ",
		"  | | ",
		" _/ | ",
		"|__/  ",
	}
	allLetters['k'] = []string{
		"       ",
		" _     ",
		"| | _  ",
		"| |/ / ",
		"|   <  ",
		"|_|\\_\\ ",
		"       ",
		"       ",
	}
	allLetters['l'] = []string{
		" _  ",
		"| | ",
		"| | ",
		"| | ",
		"| | ",
		"|_| ",
		"    ",
		"    ",
	}
	allLetters['m'] = []string{
		"            ",
		"            ",
		" _ __ ___   ",
		"| '_ ` _ \\  ",
		"| | | | | | ",
		"|_| |_| |_| ",
		"            ",
		"            ",
	}
	allLetters['n'] = []string{
		"        ",
		"        ",
		" _ __   ",
		"| '_ \\  ",
		"| | | | ",
		"|_| |_| ",
		"        ",
		"        ",
	}
	allLetters['o'] = []string{
		"        ",
		"        ",
		"  ___   ",
		" / _ \\  ",
		"| (_) | ",
		" \\___/  ",
		"        ",
		"        ",
	}
	allLetters['p'] = []string{
		"        ",
		"        ",
		" _ __   ",
		"| '_ \\  ",
		"| |_) | ",
		"| .__/  ",
		"| |     ",
		"|_|     ",
	}
	allLetters['q'] = []string{
		"        ",
		"        ",
		"  __ _  ",
		" / _` | ",
		"| (_| | ",
		" \\__, | ",
		"    | | ",
		"    |_| ",
	}
	allLetters['r'] = []string{
		"       ",
		"       ",
		" _ __  ",
		"| '__| ",
		"| |    ",
		"|_|    ",
		"       ",
		"       ",
	}
	allLetters['s'] = []string{
		"      ",
		"      ",
		" ___  ",
		"/ __| ",
		"\\__ \\ ",
		"|___/ ",
		"      ",
		"      ",
	}
	allLetters['t'] = []string{
		" _    ",
		"| |   ",
		"| |_  ",
		"| __| ",
		"\\ |_  ",
		" \\__| ",
		"      ",
		"      ",
	}
	allLetters['u'] = []string{
		"        ",
		"        ",
		" _   _  ",
		"| | | | ",
		"| |_| | ",
		" \\__,_| ",
		"        ",
		"        ",
	}
	allLetters['v'] = []string{
		"        ",
		"        ",
		"__   __ ",
		"\\ \\ / / ",
		" \\ V /  ",
		"  \\_/   ",
		"        ",
		"        ",
	}
	allLetters['w'] = []string{
		"           ",
		"           ",
		"__      __ ",
		"\\ \\ /\\ / / ",
		" \\ V  V /  ",
		"  \\_/\\_/   ",
		"           ",
		"           ",
	}
	allLetters['x'] = []string{
		"       ",
		"       ",
		"__  __ ",
		"\\ \\/ / ",
		" >  <  ",
		"/_/\\_\\ ",
		"       ",
		"       ",
	}
	allLetters['y'] = []string{
		"        ",
		"        ",
		" _   _  ",
		"| | | | ",
		"| |_| | ",
		" \\__, | ",
		" __/ /  ",
		"|___/   ",
	}
	allLetters['z'] = []string{
		"      ",
		"      ",
		" ____ ",
		"|_  / ",
		" / /  ",
		"/___| ",
		"      ",
		"      ",
	}
	allLetters['{'] = []string{
		"   __ ",
		"  / / ",
		" | |  ",
		"/ /   ",
		"\\ \\   ",
		" | |  ",
		"  \\_\\ ",
		"      ",
	}
	allLetters['}'] = []string{
		"__    ",
		"\\ \\   ",
		" | |  ",
		"  \\ \\ ",
		"  / / ",
		" | |  ",
		"/_/   ",
		"      ",
	}
	allLetters['|'] = []string{
		" _  ",
		"| | ",
		"| | ",
		"| | ",
		"| | ",
		"| | ",
		"| | ",
		"|_| ",
	}
	allLetters['~'] = []string{
		" /\\/| ",
		"|/\\/  ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	for harf, drawn := range allLetters {
		if harf == r {
			return drawn
		}
	}
	return allLetters[' ']
}
func shadow2(r rune) []string {
	allLetters := make(map[rune][]string)
	allLetters[' '] = []string{
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	allLetters['!'] = []string{
		"   ",
		"_| ",
		"_| ",
		"_| ",
		"   ",
		"_| ",
		"   ",
		"   ",
	}
	allLetters['"'] = []string{
		"_|  _| ",
		"_|  _| ",
		"       ",
		"       ",
		"       ",
		"       ",
		"       ",
		"       ",
	}
	allLetters['#'] = []string{
		"           ",
		"  _|  _|   ",
		"_|_|_|_|_| ",
		"  _|  _|   ",
		"_|_|_|_|_| ",
		"  _|  _|   ",
		"           ",
		"           ",
	}
	allLetters['$'] = []string{
		"       ",
		"  _|   ",
		"_|_|_| ",
		"_|_|   ",
		"  _|_| ",
		"_|_|_| ",
		"  _|   ",
		"       ",
	}
	allLetters['%'] = []string{
		"           ",
		"_|_|    _| ",
		"_|_|  _|   ",
		"    _|     ",
		"  _|  _|_| ",
		"_|    _|_| ",
		"           ",
		"           ",
	}
	allLetters['&'] = []string{
		"           ",
		"  _|       ",
		"_|  _|     ",
		"  _|_|  _| ",
		"_|    _|   ",
		"  _|_|  _| ",
		"           ",
		"           ",
	}
	allLetters['\''] = []string{
		"  _| ",
		"_|   ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
	}
	allLetters['('] = []string{
		"  _| ",
		"_|   ",
		"_|   ",
		"_|   ",
		"_|   ",
		"_|   ",
		"  _| ",
		"     ",
	}
	allLetters[')'] = []string{
		"_|   ",
		"  _| ",
		"  _| ",
		"  _| ",
		"  _| ",
		"  _| ",
		"_|   ",
		"     ",
	}
	allLetters['*'] = []string{
		"           ",
		"_|  _|  _| ",
		"  _|_|_|   ",
		"_|_|_|_|_| ",
		"  _|_|_|   ",
		"_|  _|  _| ",
		"           ",
		"           ",
	}
	allLetters['+'] = []string{
		"           ",
		"    _|     ",
		"    _|     ",
		"_|_|_|_|_| ",
		"    _|     ",
		"    _|     ",
		"           ",
		"           ",
	}
	allLetters[','] = []string{
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"  _| ",
		"_|   ",
		"     ",
	}
	allLetters['-'] = []string{
		"           ",
		"           ",
		"           ",
		"_|_|_|_|_| ",
		"           ",
		"           ",
		"           ",
		"           ",
	}
	allLetters['.'] = []string{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
		"_| ",
		"   ",
		"   ",
	}
	allLetters['/'] = []string{
		"           ",
		"        _| ",
		"      _|   ",
		"    _|     ",
		"  _|       ",
		"_|         ",
		"           ",
		"           ",
	}
	allLetters['0'] = []string{
		"       ",
		"  _|   ",
		"_|  _| ",
		"_|  _| ",
		"_|  _| ",
		"  _|   ",
		"       ",
		"       ",
	}
	allLetters['1'] = []string{
		"     ",
		"  _| ",
		"_|_| ",
		"  _| ",
		"  _| ",
		"  _| ",
		"     ",
		"     ",
	}
	allLetters['2'] = []string{
		"         ",
		"  _|_|   ",
		"_|    _| ",
		"    _|   ",
		"  _|     ",
		"_|_|_|_| ",
		"         ",
		"         ",
	}
	allLetters['3'] = []string{
		"         ",
		"_|_|_|   ",
		"      _| ",
		"  _|_|   ",
		"      _| ",
		"_|_|_|   ",
		"         ",
		"         ",
	}
	allLetters['4'] = []string{
		"         ",
		"_|  _|   ",
		"_|  _|   ",
		"_|_|_|_| ",
		"    _|   ",
		"    _|   ",
		"         ",
		"         ",
	}
	allLetters['5'] = []string{
		"         ",
		"_|_|_|_| ",
		"_|       ",
		"_|_|_|   ",
		"      _| ",
		"_|_|_|   ",
		"         ",
		"         ",
	}
	allLetters['6'] = []string{
		"         ",
		"  _|_|_| ",
		"_|       ",
		"_|_|_|   ",
		"_|    _| ",
		"  _|_|   ",
		"         ",
		"         ",
	}
	allLetters['7'] = []string{
		"           ",
		"_|_|_|_|_| ",
		"        _| ",
		"      _|   ",
		"    _|     ",
		"  _|       ",
		"           ",
		"           ",
	}
	allLetters['8'] = []string{
		"         ",
		"  _|_|   ",
		"_|    _| ",
		"  _|_|   ",
		"_|    _| ",
		"  _|_|   ",
		"         ",
		"         ",
	}
	allLetters['9'] = []string{
		"         ",
		"  _|_|   ",
		"_|    _| ",
		"  _|_|_| ",
		"      _| ",
		"_|_|_|   ",
		"         ",
		"         ",
	}
	allLetters[':'] = []string{
		"   ",
		"   ",
		"_| ",
		"   ",
		"   ",
		"_| ",
		"   ",
		"   ",
	}
	allLetters[';'] = []string{
		"     ",
		"     ",
		"  _| ",
		"     ",
		"     ",
		"  _| ",
		"_|   ",
		"     ",
	}
	allLetters['<'] = []string{
		"       ",
		"    _| ",
		"  _|   ",
		"_|     ",
		"  _|   ",
		"    _| ",
		"       ",
		"       ",
	}
	allLetters['='] = []string{
		"           ",
		"           ",
		"_|_|_|_|_| ",
		"           ",
		"_|_|_|_|_| ",
		"           ",
		"           ",
		"           ",
	}
	allLetters['>'] = []string{
		"       ",
		"_|     ",
		"  _|   ",
		"    _| ",
		"  _|   ",
		"_|     ",
		"       ",
		"       ",
	}
	allLetters['?'] = []string{
		"       ",
		"_|_|   ",
		"    _| ",
		"_|_|   ",
		"       ",
		"_|     ",
		"       ",
		"       ",
	}
	allLetters['@'] = []string{
		"                 ",
		"    _|_|_|_|_|   ",
		"  _|          _| ",
		"_|    _|_|_|  _| ",
		"_|  _|    _|  _| ",
		"_|    _|_|_|_|   ",
		"  _|             ",
		"    _|_|_|_|_|_| ",
	}
	allLetters['A'] = []string{
		"         ",
		"  _|_|   ",
		"_|    _| ",
		"_|_|_|_| ",
		"_|    _| ",
		"_|    _| ",
		"         ",
		"         ",
	}
	allLetters['B'] = []string{
		"         ",
		"_|_|_|   ",
		"_|    _| ",
		"_|_|_|   ",
		"_|    _| ",
		"_|_|_|   ",
		"         ",
		"         ",
	}
	allLetters['C'] = []string{
		"         ",
		"  _|_|_| ",
		"_|       ",
		"_|       ",
		"_|       ",
		"  _|_|_| ",
		"         ",
		"         ",
	}
	allLetters['D'] = []string{
		"         ",
		"_|_|_|   ",
		"_|    _| ",
		"_|    _| ",
		"_|    _| ",
		"_|_|_|   ",
		"         ",
		"         ",
	}
	allLetters['E'] = []string{
		"         ",
		"_|_|_|_| ",
		"_|       ",
		"_|_|_|   ",
		"_|       ",
		"_|_|_|_| ",
		"         ",
		"         ",
	}
	allLetters['F'] = []string{
		"         ",
		"_|_|_|_| ",
		"_|       ",
		"_|_|_|   ",
		"_|       ",
		"_|       ",
		"         ",
		"         ",
	}
	allLetters['G'] = []string{
		"         ",
		"  _|_|_| ",
		"_|       ",
		"_|  _|_| ",
		"_|    _| ",
		"  _|_|_| ",
		"         ",
		"         ",
	}
	allLetters['H'] = []string{
		"         ",
		"_|    _| ",
		"_|    _| ",
		"_|_|_|_| ",
		"_|    _| ",
		"_|    _| ",
		"         ",
		"         ",
	}
	allLetters['I'] = []string{
		"       ",
		"_|_|_| ",
		"  _|   ",
		"  _|   ",
		"  _|   ",
		"_|_|_| ",
		"       ",
		"       ",
	}
	allLetters['J'] = []string{
		"         ",
		"      _| ",
		"      _| ",
		"      _| ",
		"_|    _| ",
		"  _|_|   ",
		"         ",
		"         ",
	}
	allLetters['K'] = []string{
		"         ",
		"_|    _| ",
		"_|  _|   ",
		"_|_|     ",
		"_|  _|   ",
		"_|    _| ",
		"         ",
		"         ",
	}
	allLetters['L'] = []string{
		"         ",
		"_|       ",
		"_|       ",
		"_|       ",
		"_|       ",
		"_|_|_|_| ",
		"         ",
		"         ",
	}
	allLetters['M'] = []string{
		"           ",
		"_|      _| ",
		"_|_|  _|_| ",
		"_|  _|  _| ",
		"_|      _| ",
		"_|      _| ",
		"           ",
		"           ",
	}
	allLetters['N'] = []string{
		"           ",
		"_|      _| ",
		"_|_|    _| ",
		"_|  _|  _| ",
		"_|    _|_| ",
		"_|      _| ",
		"           ",
		"           ",
	}
	allLetters['O'] = []string{
		"         ",
		"  _|_|   ",
		"_|    _| ",
		"_|    _| ",
		"_|    _| ",
		"  _|_|   ",
		"         ",
		"         ",
	}
	allLetters['P'] = []string{
		"         ",
		"_|_|_|   ",
		"_|    _| ",
		"_|_|_|   ",
		"_|       ",
		"_|       ",
		"         ",
		"         ",
	}
	allLetters['Q'] = []string{
		"           ",
		"  _|_|     ",
		"_|    _|   ",
		"_|  _|_|   ",
		"_|    _|   ",
		"  _|_|  _| ",
		"           ",
		"           ",
	}
	allLetters['R'] = []string{
		"         ",
		"_|_|_|   ",
		"_|    _| ",
		"_|_|_|   ",
		"_|    _| ",
		"_|    _| ",
		"         ",
		"         ",
	}
	allLetters['S'] = []string{
		"         ",
		"  _|_|_| ",
		"_|       ",
		"  _|_|   ",
		"      _| ",
		"_|_|_|   ",
		"         ",
		"         ",
	}
	allLetters['T'] = []string{
		"           ",
		"_|_|_|_|_| ",
		"    _|     ",
		"    _|     ",
		"    _|     ",
		"    _|     ",
		"           ",
		"           ",
	}
	allLetters['U'] = []string{
		"         ",
		"_|    _| ",
		"_|    _| ",
		"_|    _| ",
		"_|    _| ",
		"  _|_|   ",
		"         ",
		"         ",
	}
	allLetters['V'] = []string{
		"           ",
		"_|      _| ",
		"_|      _| ",
		"_|      _| ",
		"  _|  _|   ",
		"    _|     ",
		"           ",
		"           ",
	}
	allLetters['W'] = []string{
		"               ",
		"_|          _| ",
		"_|          _| ",
		"_|    _|    _| ",
		"  _|  _|  _|   ",
		"    _|  _|     ",
		"               ",
		"               ",
	}
	allLetters['X'] = []string{
		"           ",
		"_|      _| ",
		"  _|  _|   ",
		"    _|     ",
		"  _|  _|   ",
		"_|      _| ",
		"           ",
		"           ",
	}
	allLetters['Y'] = []string{
		"           ",
		"_|      _| ",
		"  _|  _|   ",
		"    _|     ",
		"    _|     ",
		"    _|     ",
		"           ",
		"           ",
	}
	allLetters['Z'] = []string{
		"           ",
		"_|_|_|_|_| ",
		"      _|   ",
		"    _|     ",
		"  _|       ",
		"_|_|_|_|_| ",
		"           ",
		"           ",
	}
	allLetters['['] = []string{
		"_|_| ",
		"_|   ",
		"_|   ",
		"_|   ",
		"_|   ",
		"_|   ",
		"_|_| ",
		"     ",
	}
	allLetters['\\'] = []string{
		"           ",
		"_|         ",
		"  _|       ",
		"    _|     ",
		"      _|   ",
		"        _| ",
		"           ",
		"           ",
	}
	allLetters[']'] = []string{
		"_|_| ",
		"  _| ",
		"  _| ",
		"  _| ",
		"  _| ",
		"  _| ",
		"_|_| ",
		"     ",
	}
	allLetters['^'] = []string{
		"  _|   ",
		"_|  _| ",
		"       ",
		"       ",
		"       ",
		"       ",
		"       ",
		"       ",
	}
	allLetters['_'] = []string{
		"           ",
		"           ",
		"           ",
		"           ",
		"           ",
		"           ",
		"           ",
		"_|_|_|_|_| ",
	}
	allLetters['`'] = []string{
		"_|   ",
		"  _| ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
	}
	allLetters['a'] = []string{
		"         ",
		"         ",
		"  _|_|_| ",
		"_|    _| ",
		"_|    _| ",
		"  _|_|_| ",
		"         ",
		"         ",
	}
	allLetters['b'] = []string{
		"         ",
		"_|       ",
		"_|_|_|   ",
		"_|    _| ",
		"_|    _| ",
		"_|_|_|   ",
		"         ",
		"         ",
	}
	allLetters['c'] = []string{
		"         ",
		"         ",
		"  _|_|_| ",
		"_|       ",
		"_|       ",
		"  _|_|_| ",
		"         ",
		"         ",
	}
	allLetters['d'] = []string{
		"         ",
		"      _| ",
		"  _|_|_| ",
		"_|    _| ",
		"_|    _| ",
		"  _|_|_| ",
		"         ",
		"         ",
	}
	allLetters['e'] = []string{
		"         ",
		"         ",
		"  _|_|   ",
		"_|_|_|_| ",
		"_|       ",
		"  _|_|_| ",
		"         ",
		"         ",
	}
	allLetters['f'] = []string{
		"         ",
		"    _|_| ",
		"  _|     ",
		"_|_|_|_| ",
		"  _|     ",
		"  _|     ",
		"         ",
		"         ",
	}
	allLetters['g'] = []string{
		"         ",
		"         ",
		"  _|_|_| ",
		"_|    _| ",
		"_|    _| ",
		"  _|_|_| ",
		"      _| ",
		"  _|_|   ",
	}
	allLetters['h'] = []string{
		"         ",
		"_|       ",
		"_|_|_|   ",
		"_|    _| ",
		"_|    _| ",
		"_|    _| ",
		"         ",
		"         ",
	}
	allLetters['i'] = []string{
		"   ",
		"_| ",
		"   ",
		"_| ",
		"_| ",
		"_| ",
		"   ",
		"   ",
	}
	allLetters['j'] = []string{
		"     ",
		"  _| ",
		"     ",
		"  _| ",
		"  _| ",
		"  _| ",
		"  _| ",
		"_|   ",
	}
	allLetters['k'] = []string{
		"         ",
		"_|       ",
		"_|  _|   ",
		"_|_|     ",
		"_|  _|   ",
		"_|    _| ",
		"         ",
		"         ",
	}
	allLetters['l'] = []string{
		"   ",
		"_| ",
		"_| ",
		"_| ",
		"_| ",
		"_| ",
		"   ",
		"   ",
	}
	allLetters['m'] = []string{
		"               ",
		"               ",
		"_|_|_|  _|_|   ",
		"_|    _|    _| ",
		"_|    _|    _| ",
		"_|    _|    _| ",
		"               ",
		"               ",
	}
	allLetters['n'] = []string{
		"         ",
		"         ",
		"_|_|_|   ",
		"_|    _| ",
		"_|    _| ",
		"_|    _| ",
		"         ",
		"         ",
	}
	allLetters['o'] = []string{
		"         ",
		"         ",
		"  _|_|   ",
		"_|    _| ",
		"_|    _| ",
		"  _|_|   ",
		"         ",
		"         ",
	}
	allLetters['p'] = []string{
		"         ",
		"         ",
		"_|_|_|   ",
		"_|    _| ",
		"_|    _| ",
		"_|_|_|   ",
		"_|       ",
		"_|       ",
	}
	allLetters['q'] = []string{
		"         ",
		"         ",
		"  _|_|_| ",
		"_|    _| ",
		"_|    _| ",
		"  _|_|_| ",
		"      _| ",
		"      _| ",
	}
	allLetters['r'] = []string{
		"         ",
		"         ",
		"_|  _|_| ",
		"_|_|     ",
		"_|       ",
		"_|       ",
		"         ",
		"         ",
	}
	allLetters['s'] = []string{
		"         ",
		"         ",
		"  _|_|_| ",
		"_|_|     ",
		"    _|_| ",
		"_|_|_|   ",
		"         ",
		"         ",
	}
	allLetters['t'] = []string{
		"         ",
		"  _|     ",
		"_|_|_|_| ",
		"  _|     ",
		"  _|     ",
		"    _|_| ",
		"         ",
		"         ",
	}
	allLetters['u'] = []string{
		"         ",
		"         ",
		"_|    _| ",
		"_|    _| ",
		"_|    _| ",
		"  _|_|_| ",
		"         ",
		"         ",
	}
	allLetters['v'] = []string{
		"           ",
		"           ",
		"_|      _| ",
		"_|      _| ",
		"  _|  _|   ",
		"    _|     ",
		"           ",
		"           ",
	}
	allLetters['w'] = []string{
		"                   ",
		"                   ",
		"_|      _|      _| ",
		"_|      _|      _| ",
		"  _|  _|  _|  _|   ",
		"    _|      _|     ",
		"                   ",
		"                   ",
	}
	allLetters['x'] = []string{
		"         ",
		"         ",
		"_|    _| ",
		"  _|_|   ",
		"_|    _| ",
		"_|    _| ",
		"         ",
		"         ",
	}
	allLetters['y'] = []string{
		"         ",
		"         ",
		"_|    _| ",
		"_|    _| ",
		"_|    _| ",
		"  _|_|_| ",
		"      _| ",
		"  _|_|   ",
	}
	allLetters['z'] = []string{
		"         ",
		"         ",
		"_|_|_|_| ",
		"    _|   ",
		"  _|     ",
		"_|_|_|_| ",
		"         ",
		"         ",
	}
	allLetters['{'] = []string{
		"    _| ",
		"  _|   ",
		"  _|   ",
		"_|     ",
		"  _|   ",
		"  _|   ",
		"    _| ",
		"       ",
	}
	allLetters['|'] = []string{
		"_| ",
		"_| ",
		"_| ",
		"_| ",
		"_| ",
		"_| ",
		"_| ",
		"_| ",
	}
	allLetters['}'] = []string{
		"_|     ",
		"  _|   ",
		"  _|   ",
		"    _| ",
		"  _|   ",
		"  _|   ",
		"_|     ",
		"       ",
	}
	allLetters['~'] = []string{
		"  _|  _| ",
		"_|  _|   ",
		"         ",
		"         ",
		"         ",
		"         ",
		"         ",
		"         ",
	}
	for harf, drawn := range allLetters {
		if harf == r {
			return drawn
		}
	}
	return allLetters[' ']
}
func thinkertoy2(r rune) []string {
	allLetters := make(map[rune][]string)
	allLetters[' '] = []string{
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	allLetters['!'] = []string{
		"  ",
		"o ",
		"| ",
		"o ",
		"  ",
		"O ",
		"  ",
		"  ",
	}
	allLetters['"'] = []string{
		"o o ",
		"| | ",
		"    ",
		"    ",
		"    ",
		"    ",
		"    ",
		"    ",
	}
	allLetters['#'] = []string{
		"      ",
		" | |  ",
		"-O-O- ",
		" | |  ",
		"-O-O- ",
		" | |  ",
		"      ",
		"      ",
	}
	allLetters['$'] = []string{
		"  | |   ",
		" -O-O-  ",
		"o | |   ",
		" -O-O-  ",
		"  | | o ",
		" -O-O-  ",
		"  | |   ",
		"        ",
	}
	allLetters['%'] = []string{
		"      ",
		"    O ",
		"o  /  ",
		"  /   ",
		" /  o ",
		"O     ",
		"      ",
		"      ",
	}
	allLetters['&'] = []string{
		"     ",
		"     ",
		"  o  ",
		" /|  ",
		"o-O- ",
		"  |  ",
		"     ",
		"     ",
	}
	allLetters['\''] = []string{
		"o ",
		"| ",
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
	}
	allLetters['('] = []string{
		"   ",
		" / ",
		"o  ",
		"|  ",
		"o  ",
		" \\ ",
		"   ",
		"   ",
	}
	allLetters[')'] = []string{
		"   ",
		"\\  ",
		" o ",
		" | ",
		" o ",
		"/  ",
		"   ",
		"   ",
	}
	allLetters['*'] = []string{
		"      ",
		"o | o ",
		" \\|/  ",
		"--O-- ",
		" /|\\  ",
		"o | o ",
		"      ",
		"      ",
	}
	allLetters['+'] = []string{
		"    ",
		"    ",
		" |  ",
		"-o- ",
		" |  ",
		"    ",
		"    ",
		"    ",
	}
	allLetters[','] = []string{
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
		"o ",
		"| ",
		"  ",
	}
	allLetters['-'] = []string{
		"    ",
		"    ",
		"    ",
		"    ",
		"o-o ",
		"    ",
		"    ",
		"    ",
	}
	allLetters['.'] = []string{
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
		"O ",
		"  ",
		"  ",
	}
	allLetters['/'] = []string{
		"      ",
		"    o ",
		"   /  ",
		"  o   ",
		" /    ",
		"o     ",
		"      ",
		"      ",
	}
	allLetters['0'] = []string{
		"      ",
		" o-o  ",
		"o  /o ",
		"| / | ",
		"o/  o ",
		" o-o  ",
		"      ",
		"      ",
	}
	allLetters['1'] = []string{
		"      ",
		"  0   ",
		" /|   ",
		"o |   ",
		"  |   ",
		"o-o-o ",
		"      ",
		"      ",
	}
	allLetters['2'] = []string{
		"     ",
		" --  ",
		"o  o ",
		"  /  ",
		" /   ",
		"o--o ",
		"     ",
		"     ",
	}
	allLetters['3'] = []string{
		"     ",
		"o-o  ",
		"   | ",
		" oo  ",
		"   | ",
		"o-o  ",
		"     ",
		"     ",
	}
	allLetters['4'] = []string{
		"     ",
		"o  o ",
		"|  | ",
		"o--O ",
		"   | ",
		"   o ",
		"     ",
		"     ",
	}
	allLetters['5'] = []string{
		"     ",
		"o--o ",
		"|    ",
		"o-o  ",
		"   | ",
		"o-o  ",
		"     ",
		"     ",
	}
	allLetters['6'] = []string{
		"      ",
		"  o   ",
		" /    ",
		"O--o  ",
		"o   | ",
		" o-o  ",
		"      ",
		"      ",
	}
	allLetters['7'] = []string{
		"      ",
		"o---o ",
		"   /  ",
		"  o   ",
		"  |   ",
		"  o   ",
		"      ",
		"      ",
	}
	allLetters['8'] = []string{
		"      ",
		" o-o  ",
		"|   | ",
		" o-o  ",
		"|   | ",
		" o-o  ",
		"      ",
		"      ",
	}
	allLetters['9'] = []string{
		"      ",
		" o-o  ",
		"|   o ",
		" o--O ",
		"   /  ",
		"  o   ",
		"      ",
		"      ",
	}
	allLetters[':'] = []string{
		"  ",
		"  ",
		"O ",
		"  ",
		"O ",
		"  ",
		"  ",
		"  ",
	}
	allLetters[';'] = []string{
		"  ",
		"  ",
		"o ",
		"  ",
		"o ",
		"| ",
		"  ",
		"  ",
	}
	allLetters['<'] = []string{
		"    ",
		"  o ",
		" /  ",
		"O   ",
		" \\  ",
		"  o ",
		"    ",
		"    ",
	}
	allLetters['='] = []string{
		"     ",
		"     ",
		"     ",
		"o--o ",
		"o--o ",
		"     ",
		"     ",
		"     ",
	}
	allLetters['>'] = []string{
		"    ",
		"o   ",
		" \\  ",
		"  O ",
		" /  ",
		"o   ",
		"    ",
		"    ",
	}
	allLetters['?'] = []string{
		"      ",
		" o-o  ",
		"o   o ",
		"   /  ",
		"  o   ",
		"      ",
		"  O   ",
		"      ",
	}
	allLetters['@'] = []string{
		"      ",
		"  o   ",
		" / \\  ",
		"o O-o ",
		" \\    ",
		"  o-  ",
		"      ",
		"      ",
	}
	allLetters['A'] = []string{
		"      ",
		"  O   ",
		" / \\  ",
		"o---o ",
		"|   | ",
		"o   o ",
		"      ",
		"      ",
	}
	allLetters['B'] = []string{
		"      ",
		"o--o  ",
		"|   | ",
		"O--o  ",
		"|   | ",
		"o--o  ",
		"      ",
		"      ",
	}
	allLetters['C'] = []string{
		"      ",
		"  o-o ",
		" /    ",
		"O     ",
		" \\    ",
		"  o-o ",
		"      ",
		"      ",
	}
	allLetters['D'] = []string{
		"      ",
		"o-o   ",
		"|  \\  ",
		"|   O ",
		"|  /  ",
		"o-o   ",
		"      ",
		"      ",
	}
	allLetters['E'] = []string{
		"     ",
		"o--o ",
		"|    ",
		"O-o  ",
		"|    ",
		"o--o ",
		"     ",
		"     ",
	}
	allLetters['F'] = []string{
		"     ",
		"o--o ",
		"|    ",
		"O-o  ",
		"|    ",
		"o    ",
		"     ",
		"     ",
	}
	allLetters['G'] = []string{
		"      ",
		" o-o  ",
		"o     ",
		"|  -o ",
		"o   | ",
		" o-o  ",
		"      ",
		"      ",
	}
	allLetters['H'] = []string{
		"     ",
		"o  o ",
		"|  | ",
		"O--O ",
		"|  | ",
		"o  o ",
		"     ",
		"     ",
	}
	allLetters['I'] = []string{
		"      ",
		"o-O-o ",
		"  |   ",
		"  |   ",
		"  |   ",
		"o-O-o ",
		"      ",
		"      ",
	}
	allLetters['J'] = []string{
		"      ",
		"    o ",
		"    | ",
		"    | ",
		"\\   o ",
		" o-o  ",
		"      ",
		"      ",
	}
	allLetters['K'] = []string{
		"     ",
		"o  o ",
		"| /  ",
		"OO   ",
		"| \\  ",
		"o  o ",
		"     ",
		"     ",
	}
	allLetters['L'] = []string{
		"      ",
		"o     ",
		"|     ",
		"|     ",
		"|     ",
		"O---o ",
		"      ",
		"      ",
	}
	allLetters['M'] = []string{
		"      ",
		"o   o ",
		"|\\ /| ",
		"| O | ",
		"|   | ",
		"o   o ",
		"      ",
		"      ",
	}
	allLetters['N'] = []string{
		"      ",
		"o   o ",
		"|\\  | ",
		"| \\ | ",
		"|  \\| ",
		"o   o ",
		"      ",
		"      ",
	}
	allLetters['O'] = []string{
		"      ",
		" o-o  ",
		"o   o ",
		"|   | ",
		"o   o ",
		" o-o  ",
		"      ",
		"      ",
	}
	allLetters['P'] = []string{
		"      ",
		"o--o  ",
		"|   | ",
		"O--o  ",
		"|     ",
		"o     ",
		"      ",
		"      ",
	}
	allLetters['Q'] = []string{
		"      ",
		" o-o  ",
		"o   o ",
		"|   | ",
		"o   O ",
		" o-O\\ ",
		"      ",
		"      ",
	}
	allLetters['R'] = []string{
		"      ",
		"o--o  ",
		"|   | ",
		"O-Oo  ",
		"|  \\  ",
		"o   o ",
		"      ",
		"      ",
	}
	allLetters['S'] = []string{
		"      ",
		" o-o  ",
		"|     ",
		" o-o  ",
		"    | ",
		"o--o  ",
		"      ",
		"      ",
	}
	allLetters['T'] = []string{
		"      ",
		"o-O-o ",
		"  |   ",
		"  |   ",
		"  |   ",
		"  o   ",
		"      ",
		"      ",
	}
	allLetters['U'] = []string{
		"      ",
		"o   o ",
		"|   | ",
		"|   | ",
		"|   | ",
		" o-o  ",
		"      ",
		"      ",
	}
	allLetters['V'] = []string{
		"      ",
		"o   o ",
		"|   | ",
		"o   o ",
		" \\ /  ",
		"  o   ",
		"      ",
		"      ",
	}
	allLetters['W'] = []string{
		"          ",
		"o       o ",
		"|       | ",
		"o   o   o ",
		" \\ / \\ /  ",
		"  o   o   ",
		"          ",
		"          ",
	}
	allLetters['X'] = []string{
		"      ",
		"o   o ",
		" \\ /  ",
		"  O   ",
		" / \\  ",
		"o   o ",
		"      ",
		"      ",
	}
	allLetters['Y'] = []string{
		"      ",
		"o   o ",
		" \\ /  ",
		"  O   ",
		"  |   ",
		"  o   ",
		"      ",
		"      ",
	}
	allLetters['Z'] = []string{
		"      ",
		"o---o ",
		"   /  ",
		" -O-  ",
		" /    ",
		"o---o ",
		"      ",
		"      ",
	}
	allLetters['['] = []string{
		"    ",
		"O-o ",
		"|   ",
		"|   ",
		"|   ",
		"O-o ",
		"    ",
		"    ",
	}
	allLetters['\\'] = []string{
		"      ",
		"o     ",
		" \\    ",
		"  o   ",
		"   \\  ",
		"    o ",
		"      ",
		"      ",
	}
	allLetters[']'] = []string{
		"    ",
		"o-O ",
		"  | ",
		"  | ",
		"  | ",
		"o-O ",
		"    ",
		"    ",
	}
	allLetters['^'] = []string{
		"    ",
		" o  ",
		"/ \\ ",
		"    ",
		"    ",
		"    ",
		"    ",
		"    ",
	}
	allLetters['_'] = []string{
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"o---o ",
		"      ",
		"      ",
	}
	allLetters['`'] = []string{
		"  ",
		"0 ",
		"| ",
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
	}
	allLetters['a'] = []string{
		"     ",
		"     ",
		"     ",
		" oo  ",
		"| |  ",
		"o-o- ",
		"     ",
		"     ",
	}
	allLetters['b'] = []string{
		"     ",
		"o    ",
		"|    ",
		"O-o  ",
		"|  | ",
		"o-o  ",
		"     ",
		"     ",
	}
	allLetters['c'] = []string{
		"     ",
		"     ",
		"     ",
		" o-o ",
		"|    ",
		" o-o ",
		"     ",
		"     ",
	}
	allLetters['d'] = []string{
		"     ",
		"   o ",
		"   | ",
		" o-O ",
		"|  | ",
		" o-o ",
		"     ",
		"     ",
	}
	allLetters['e'] = []string{
		"    ",
		"    ",
		"    ",
		"o-o ",
		"|-' ",
		"o-o ",
		"    ",
		"    ",
	}
	allLetters['f'] = []string{
		"     ",
		" o-o ",
		" |   ",
		"-O-  ",
		" |   ",
		" o   ",
		"     ",
		"     ",
	}
	allLetters['g'] = []string{
		"     ",
		"     ",
		"     ",
		"o--o ",
		"|  | ",
		"o--O ",
		"   | ",
		"o--o ",
	}
	allLetters['h'] = []string{
		"     ",
		"o    ",
		"|    ",
		"O--o ",
		"|  | ",
		"o  o ",
		"     ",
		"     ",
	}
	allLetters['i'] = []string{
		"  ",
		"  ",
		"o ",
		"  ",
		"| ",
		"| ",
		"  ",
		"  ",
	}
	allLetters['j'] = []string{
		"      ",
		"      ",
		"    o ",
		"      ",
		"    o ",
		"    | ",
		"o   o ",
		" o-o  ",
	}
	allLetters['k'] = []string{
		"     ",
		"o    ",
		"| /  ",
		"OO   ",
		"| \\  ",
		"o  o ",
		"     ",
		"     ",
	}
	allLetters['l'] = []string{
		"  ",
		"o ",
		"| ",
		"| ",
		"| ",
		"o ",
		"  ",
		"  ",
	}
	allLetters['m'] = []string{
		"      ",
		"      ",
		"      ",
		"o-O-o ",
		"| | | ",
		"o o o ",
		"      ",
		"      ",
	}
	allLetters['n'] = []string{
		"     ",
		"     ",
		"     ",
		"o-o  ",
		"|  | ",
		"o  o ",
		"     ",
		"     ",
	}
	allLetters['o'] = []string{
		"    ",
		"    ",
		"    ",
		"o-o ",
		"| | ",
		"o-o ",
		"    ",
		"    ",
	}
	allLetters['p'] = []string{
		"     ",
		"     ",
		"     ",
		"o-o  ",
		"|  | ",
		"O-o  ",
		"|    ",
		"o    ",
	}
	allLetters['q'] = []string{
		"     ",
		"     ",
		"     ",
		" o-o ",
		"|  | ",
		" o-O ",
		"   | ",
		"   o ",
	}
	allLetters['r'] = []string{
		"    ",
		"    ",
		"    ",
		"o-o ",
		"|   ",
		"o   ",
		"    ",
		"    ",
	}
	allLetters['s'] = []string{
		"    ",
		"    ",
		"    ",
		"o-o ",
		" \\  ",
		"o-o ",
		"    ",
		"    ",
	}
	allLetters['t'] = []string{
		"    ",
		" o  ",
		" |  ",
		"-o- ",
		" |  ",
		" o  ",
		"    ",
		"    ",
	}
	allLetters['u'] = []string{
		"     ",
		"     ",
		"     ",
		"o  o ",
		"|  | ",
		"o--o ",
		"     ",
		"     ",
	}
	allLetters['v'] = []string{
		"      ",
		"      ",
		"      ",
		"o   o ",
		" \\ /  ",
		"  o   ",
		"      ",
		"      ",
	}
	allLetters['w'] = []string{
		"          ",
		"          ",
		"          ",
		"o   o   o ",
		" \\ / \\ /  ",
		"  o   o   ",
		"          ",
		"          ",
	}
	allLetters['x'] = []string{
		"    ",
		"    ",
		"    ",
		"\\ / ",
		" o  ",
		"/ \\ ",
		"    ",
		"    ",
	}
	allLetters['y'] = []string{
		"     ",
		"     ",
		"     ",
		"o  o ",
		"|  | ",
		"o--O ",
		"   | ",
		"o--o ",
	}
	allLetters['z'] = []string{
		"    ",
		"    ",
		"    ",
		"o-o ",
		" /  ",
		"o-o ",
		"    ",
		"    ",
	}
	allLetters['{'] = []string{
		"      ",
		"  o-o ",
		"  |   ",
		"o-O   ",
		"  |   ",
		"  o-o ",
		"      ",
		"      ",
	}
	allLetters['|'] = []string{
		"  ",
		"o ",
		"| ",
		"o ",
		"| ",
		"o ",
		"  ",
		"  ",
	}
	allLetters['}'] = []string{
		"      ",
		"o-o   ",
		"  |   ",
		"  O-o ",
		"  |   ",
		"o-o   ",
		"      ",
		"      ",
	}
	allLetters['~'] = []string{
		" o_ / ",
		"/  o  ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	for harf, drawn := range allLetters {
		if harf == r {
			return drawn
		}
	}
	return allLetters[' ']
}
