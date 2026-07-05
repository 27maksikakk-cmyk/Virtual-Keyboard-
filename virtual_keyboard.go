// virtual_keyboard.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	reset  = "\033[0m"
	green  = "\033[92m"
	blue   = "\033[94m"
	yellow = "\033[93m"
	bold   = "\033[1m"
)

func colorize(text, color string) string {
	return color + text + reset
}

var keyboardEN = [][]string{
	{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
	{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P"},
	{"A", "S", "D", "F", "G", "H", "J", "K", "L"},
	{"Z", "X", "C", "V", "B", "N", "M"},
}

var keyboardRU = [][]string{
	{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
	{"Й", "Ц", "У", "К", "Е", "Н", "Г", "Ш", "Щ", "З"},
	{"Ф", "Ы", "В", "А", "П", "Р", "О", "Л", "Д"},
	{"Я", "Ч", "С", "М", "И", "Т", "Ь"},
}

type VirtualKeyboard struct {
	lang   string
	text   []string
	shift  bool
	layout [][]string
}

func NewVirtualKeyboard(lang string) *VirtualKeyboard {
	v := &VirtualKeyboard{lang: lang, shift: false}
	if lang == "en" {
		v.layout = keyboardEN
	} else {
		v.layout = keyboardRU
	}
	return v
}

func (v *VirtualKeyboard) display() {
	fmt.Print("\033[H\033[2J")
	fmt.Println(colorize("🖥️  ВИРТУАЛЬНАЯ КЛАВИАТУРА", bold))
	fmt.Println(colorize("Управление: номер клавиши → добавить, 0 → пробел, - → удалить, . → выход", yellow))
	fmt.Println(colorize("8 → переключить раскладку, 9 → переключить регистр (Shift)\n", yellow))
	fmt.Println(colorize("Текст: "+strings.Join(v.text, ""), green))
	fmt.Println()

	idx := 1
	for _, row := range v.layout {
		line := ""
		for _, ch := range row {
			displayCh := ch
			if !v.shift {
				displayCh = strings.ToLower(ch)
			}
			line += colorize(fmt.Sprintf("%2d", idx), blue) + ":" + colorize(displayCh, blue) + "  "
			idx++
		}
		fmt.Println(line)
	}
	fmt.Println("\nСпециальные: 0-пробел, --удалить, .-выход, 8-язык, 9-Shift")
}

func (v *VirtualKeyboard) run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		v.display()
		fmt.Print(colorize("Введите номер команды: ", bold))
		scanner.Scan()
		cmd := strings.TrimSpace(scanner.Text())
		if cmd == "." {
			break
		}
		if cmd == "-" {
			if len(v.text) > 0 {
				v.text = v.text[:len(v.text)-1]
			}
			continue
		}
		if cmd == "0" {
			v.text = append(v.text, " ")
			continue
		}
		if cmd == "8" {
			if v.lang == "en" {
				v.lang = "ru"
				v.layout = keyboardRU
			} else {
				v.lang = "en"
				v.layout = keyboardEN
			}
			continue
		}
		if cmd == "9" {
			v.shift = !v.shift
			continue
		}
		num, err := strconv.Atoi(cmd)
		if err == nil {
			flat := []string{}
			for _, row := range v.layout {
				for _, ch := range row {
					flat = append(flat, ch)
				}
			}
			if num >= 1 && num <= len(flat) {
				ch := flat[num-1]
				if v.shift {
					ch = strings.ToUpper(ch)
				} else {
					ch = strings.ToLower(ch)
				}
				v.text = append(v.text, ch)
			} else {
				fmt.Println(colorize("Неверный номер!", "red"))
			}
		} else {
			fmt.Println(colorize("Неверный ввод!", "red"))
		}
	}
	fmt.Println(colorize("\nИтоговый текст: "+strings.Join(v.text, ""), green))
}

func main() {
	kb := NewVirtualKeyboard("en")
	kb.run()
}
