// virtual_keyboard.js
#!/usr/bin/env node
'use strict';

const readline = require('readline');

const COLORS = {
    reset: '\x1b[0m',
    green: '\x1b[92m',
    blue: '\x1b[94m',
    yellow: '\x1b[93m',
    bold: '\x1b[1m'
};

function colorize(text, color) {
    return COLORS[color] + text + COLORS.reset;
}

const KEYBOARD_EN = [
    ['1','2','3','4','5','6','7','8','9','0'],
    ['Q','W','E','R','T','Y','U','I','O','P'],
    ['A','S','D','F','G','H','J','K','L'],
    ['Z','X','C','V','B','N','M']
];

const KEYBOARD_RU = [
    ['1','2','3','4','5','6','7','8','9','0'],
    ['Й','Ц','У','К','Е','Н','Г','Ш','Щ','З'],
    ['Ф','Ы','В','А','П','Р','О','Л','Д'],
    ['Я','Ч','С','М','И','Т','Ь']
];

class VirtualKeyboard {
    constructor(lang = 'en') {
        this.lang = lang;
        this.text = [];
        this.shift = false;
        this.layout = lang === 'en' ? KEYBOARD_EN : KEYBOARD_RU;
    }

    display() {
        console.clear();
        console.log(colorize('🖥️  ВИРТУАЛЬНАЯ КЛАВИАТУРА', 'bold'));
        console.log(colorize('Управление: номер клавиши → добавить, 0 → пробел, - → удалить, . → выход', 'yellow'));
        console.log(colorize('8 → переключить раскладку, 9 → переключить регистр (Shift)\n', 'yellow'));
        console.log(colorize('Текст: ' + this.text.join(''), 'green'));
        console.log();

        let idx = 1;
        for (const row of this.layout) {
            let line = '';
            for (const ch of row) {
                let displayCh = this.shift ? ch.toUpperCase() : ch.toLowerCase();
                line += colorize(String(idx).padStart(2, '0'), 'blue') + ':' + colorize(displayCh, 'blue') + '  ';
                idx++;
            }
            console.log(line);
        }
        console.log('\nСпециальные: 0-пробел, --удалить, .-выход, 8-язык, 9-Shift');
    }

    run() {
        const rl = readline.createInterface({
            input: process.stdin,
            output: process.stdout
        });
        const prompt = () => {
            this.display();
            rl.question(colorize('Введите номер команды: ', 'bold'), (cmd) => {
                cmd = cmd.trim();
                if (cmd === '.') {
                    rl.close();
                    console.log(colorize('\nИтоговый текст: ' + this.text.join(''), 'green'));
                    return;
                }
                if (cmd === '-') {
                    if (this.text.length) this.text.pop();
                    prompt();
                    return;
                }
                if (cmd === '0') {
                    this.text.push(' ');
                    prompt();
                    return;
                }
                if (cmd === '8') {
                    this.lang = this.lang === 'en' ? 'ru' : 'en';
                    this.layout = this.lang === 'en' ? KEYBOARD_EN : KEYBOARD_RU;
                    prompt();
                    return;
                }
                if (cmd === '9') {
                    this.shift = !this.shift;
                    prompt();
                    return;
                }
                const num = parseInt(cmd);
                if (!isNaN(num)) {
                    const flat = this.layout.flat();
                    if (num >= 1 && num <= flat.length) {
                        let ch = flat[num-1];
                        ch = this.shift ? ch.toUpperCase() : ch.toLowerCase();
                        this.text.push(ch);
                    } else {
                        console.log(colorize('Неверный номер!', 'red'));
                    }
                } else {
                    console.log(colorize('Неверный ввод!', 'red'));
                }
                prompt();
            });
        };
        prompt();
    }
}

const kb = new VirtualKeyboard('en');
kb.run();
