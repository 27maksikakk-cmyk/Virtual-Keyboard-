// virtual_keyboard.cpp
#include <iostream>
#include <vector>
#include <string>
#include <map>
#include <cstdlib>

using namespace std;

const string RESET = "\033[0m";
const string GREEN = "\033[92m";
const string BLUE = "\033[94m";
const string YELLOW = "\033[93m";
const string BOLD = "\033[1m";

string colorize(const string& text, const string& color) {
    return color + text + RESET;
}

vector<vector<string>> KEYBOARD_EN = {
    {"1","2","3","4","5","6","7","8","9","0"},
    {"Q","W","E","R","T","Y","U","I","O","P"},
    {"A","S","D","F","G","H","J","K","L"},
    {"Z","X","C","V","B","N","M"}
};

vector<vector<string>> KEYBOARD_RU = {
    {"1","2","3","4","5","6","7","8","9","0"},
    {"Й","Ц","У","К","Е","Н","Г","Ш","Щ","З"},
    {"Ф","Ы","В","А","П","Р","О","Л","Д"},
    {"Я","Ч","С","М","И","Т","Ь"}
};

class VirtualKeyboard {
public:
    string lang;
    string text;
    bool shift;
    vector<vector<string>> layout;

    VirtualKeyboard(string l = "en") : lang(l), shift(false) {
        layout = (lang == "en") ? KEYBOARD_EN : KEYBOARD_RU;
    }

    void display() {
        system("clear");
        cout << colorize("🖥️  ВИРТУАЛЬНАЯ КЛАВИАТУРА", BOLD) << endl;
        cout << colorize("Управление: номер клавиши → добавить, 0 → пробел, - → удалить, . → выход", YELLOW) << endl;
        cout << colorize("8 → переключить раскладку, 9 → переключить регистр (Shift)\n", YELLOW) << endl;
        cout << colorize("Текст: " + text, GREEN) << endl << endl;

        int idx = 1;
        for (auto& row : layout) {
            string line;
            for (auto& ch : row) {
                string displayCh = shift ? ch : ch;
                // для нижнего регистра в C++ неудобно, но для демонстрации оставим как есть
                line += colorize(to_string(idx).substr(0,2), BLUE) + ":" + colorize(displayCh, BLUE) + "  ";
                idx++;
            }
            cout << line << endl;
        }
        cout << "\nСпециальные: 0-пробел, --удалить, .-выход, 8-язык, 9-Shift" << endl;
    }

    void run() {
        string cmd;
        while (true) {
            display();
            cout << colorize("Введите номер команды: ", BOLD);
            cin >> cmd;
            if (cmd == ".") break;
            if (cmd == "-") {
                if (!text.empty()) text.pop_back();
                continue;
            }
            if (cmd == "0") {
                text += ' ';
                continue;
            }
            if (cmd == "8") {
                lang = (lang == "en") ? "ru" : "en";
                layout = (lang == "en") ? KEYBOARD_EN : KEYBOARD_RU;
                continue;
            }
            if (cmd == "9") {
                shift = !shift;
                continue;
            }
            try {
                int num = stoi(cmd);
                vector<string> flat;
                for (auto& row : layout)
                    for (auto& ch : row)
                        flat.push_back(ch);
                if (num >= 1 && num <= (int)flat.size()) {
                    string ch = flat[num-1];
                    if (shift) {
                        // преобразовать в верхний регистр (только для латиницы)
                        for (char& c : ch) c = toupper(c);
                    } else {
                        for (char& c : ch) c = tolower(c);
                    }
                    text += ch;
                } else {
                    cout << colorize("Неверный номер!", RED) << endl;
                    cin.ignore();
                    cin.get();
                }
            } catch (...) {
                cout << colorize("Неверный ввод!", RED) << endl;
                cin.ignore();
                cin.get();
            }
        }
        cout << colorize("\nИтоговый текст: " + text, GREEN) << endl;
    }
};

int main() {
    VirtualKeyboard kb("en");
    kb.run();
    return 0;
}
