# virtual_keyboard.py
#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import sys
import time

# ANSI-цвета
COLORS = {
    'reset': '\033[0m',
    'green': '\033[92m',
    'blue': '\033[94m',
    'yellow': '\033[93m',
    'bold': '\033[1m'
}

def colorize(text, color):
    return f"{COLORS.get(color, '')}{text}{COLORS['reset']}"

# Клавиатурные раскладки (три ряда)
KEYBOARD_EN = [
    ['1', '2', '3', '4', '5', '6', '7', '8', '9', '0'],
    ['Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P'],
    ['A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L'],
    ['Z', 'X', 'C', 'V', 'B', 'N', 'M']
]

KEYBOARD_RU = [
    ['1', '2', '3', '4', '5', '6', '7', '8', '9', '0'],
    ['Й', 'Ц', 'У', 'К', 'Е', 'Н', 'Г', 'Ш', 'Щ', 'З'],
    ['Ф', 'Ы', 'В', 'А', 'П', 'Р', 'О', 'Л', 'Д'],
    ['Я', 'Ч', 'С', 'М', 'И', 'Т', 'Ь']
]

class VirtualKeyboard:
    def __init__(self, lang='en'):
        self.lang = lang
        self.text = []
        self.shift = False
        self.layout = KEYBOARD_EN if lang == 'en' else KEYBOARD_RU

    def display(self):
        os.system('clear' if os.name == 'posix' else 'cls')
        print(colorize("🖥️  ВИРТУАЛЬНАЯ КЛАВИАТУРА", 'bold'))
        print(colorize("Управление: номер клавиши → добавить, 0 → пробел, - → удалить, . → выход", 'yellow'))
        print(colorize("8 → переключить раскладку, 9 → переключить регистр (Shift)\n", 'yellow'))
        print(colorize("Текст: " + ''.join(self.text), 'green'))
        print()

        # Отображение клавиатуры с номерами
        idx = 1
        for row in self.layout:
            line = ''
            for ch in row:
                display_ch = ch.upper() if self.shift else ch.lower()
                line += f"{colorize(str(idx).zfill(2), 'blue')}:{colorize(display_ch, 'blue')}  "
                idx += 1
            print(line)
        print("\nСпециальные: 0-пробел, --удалить, .-выход, 8-язык, 9-Shift")

    def run(self):
        while True:
            self.display()
            cmd = input(colorize("Введите номер команды: ", 'bold')).strip()
            if cmd == '.':
                break
            if cmd == '-':
                if self.text:
                    self.text.pop()
                continue
            if cmd == '0':
                self.text.append(' ')
                continue
            if cmd == '8':
                self.lang = 'ru' if self.lang == 'en' else 'en'
                self.layout = KEYBOARD_RU if self.lang == 'ru' else KEYBOARD_EN
                continue
            if cmd == '9':
                self.shift = not self.shift
                continue
            try:
                num = int(cmd)
                # Найти символ по номеру
                flat = [ch for row in self.layout for ch in row]
                if 1 <= num <= len(flat):
                    ch = flat[num-1]
                    if self.shift:
                        ch = ch.upper()
                    else:
                        ch = ch.lower()
                    self.text.append(ch)
                else:
                    print(colorize("Неверный номер!", 'red'))
                    time.sleep(0.5)
            except ValueError:
                print(colorize("Неверный ввод!", 'red'))
                time.sleep(0.5)

        print(colorize(f"\nИтоговый текст: {''.join(self.text)}", 'green'))

if __name__ == '__main__':
    try:
        kb = VirtualKeyboard('en')
        kb.run()
    except KeyboardInterrupt:
        print(colorize("\nВыход.", 'yellow'))
        sys.exit(0)
