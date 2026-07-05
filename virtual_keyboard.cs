// virtual_keyboard.cs
using System;
using System.Collections.Generic;
using System.Linq;

class VirtualKeyboard
{
    static string Colorize(string text, string color)
    {
        string col = color switch
        {
            "green" => "\x1b[92m",
            "blue" => "\x1b[94m",
            "yellow" => "\x1b[93m",
            "bold" => "\x1b[1m",
            _ => "\x1b[0m"
        };
        return col + text + "\x1b[0m";
    }

    static List<List<string>> KEYBOARD_EN = new List<List<string>>
    {
        new List<string>{"1","2","3","4","5","6","7","8","9","0"},
        new List<string>{"Q","W","E","R","T","Y","U","I","O","P"},
        new List<string>{"A","S","D","F","G","H","J","K","L"},
        new List<string>{"Z","X","C","V","B","N","M"}
    };

    static List<List<string>> KEYBOARD_RU = new List<List<string>>
    {
        new List<string>{"1","2","3","4","5","6","7","8","9","0"},
        new List<string>{"Й","Ц","У","К","Е","Н","Г","Ш","Щ","З"},
        new List<string>{"Ф","Ы","В","А","П","Р","О","Л","Д"},
        new List<string>{"Я","Ч","С","М","И","Т","Ь"}
    };

    private string lang;
    private List<string> text = new List<string>();
    private bool shift;
    private List<List<string>> layout;

    public VirtualKeyboard(string lang = "en")
    {
        this.lang = lang;
        this.shift = false;
        this.layout = lang == "en" ? KEYBOARD_EN : KEYBOARD_RU;
    }

    public void Display()
    {
        Console.Clear();
        Console.WriteLine(Colorize("🖥️  ВИРТУАЛЬНАЯ КЛАВИАТУРА", "bold"));
        Console.WriteLine(Colorize("Управление: номер клавиши → добавить, 0 → пробел, - → удалить, . → выход", "yellow"));
        Console.WriteLine(Colorize("8 → переключить раскладку, 9 → переключить регистр (Shift)\n", "yellow"));
        Console.WriteLine(Colorize("Текст: " + string.Join("", text), "green"));
        Console.WriteLine();

        int idx = 1;
        foreach (var row in layout)
        {
            string line = "";
            foreach (string ch in row)
            {
                string displayCh = shift ? ch.ToUpper() : ch.ToLower();
                line += Colorize(idx.ToString("D2"), "blue") + ":" + Colorize(displayCh, "blue") + "  ";
                idx++;
            }
            Console.WriteLine(line);
        }
        Console.WriteLine("\nСпециальные: 0-пробел, --удалить, .-выход, 8-язык, 9-Shift");
    }

    public void Run()
    {
        while (true)
        {
            Display();
            Console.Write(Colorize("Введите номер команды: ", "bold"));
            string cmd = Console.ReadLine().Trim();
            if (cmd == ".")
                break;
            if (cmd == "-")
            {
                if (text.Count > 0) text.RemoveAt(text.Count - 1);
                continue;
            }
            if (cmd == "0")
            {
                text.Add(" ");
                continue;
            }
            if (cmd == "8")
            {
                lang = lang == "en" ? "ru" : "en";
                layout = lang == "en" ? KEYBOARD_EN : KEYBOARD_RU;
                continue;
            }
            if (cmd == "9")
            {
                shift = !shift;
                continue;
            }
            if (int.TryParse(cmd, out int num))
            {
                var flat = layout.SelectMany(row => row).ToList();
                if (num >= 1 && num <= flat.Count)
                {
                    string ch = flat[num - 1];
                    ch = shift ? ch.ToUpper() : ch.ToLower();
                    text.Add(ch);
                }
                else Console.WriteLine(Colorize("Неверный номер!", "red"));
            }
            else Console.WriteLine(Colorize("Неверный ввод!", "red"));
        }
        Console.WriteLine(Colorize("\nИтоговый текст: " + string.Join("", text), "green"));
    }

    static void Main()
    {
        var kb = new VirtualKeyboard("en");
        kb.Run();
    }
}
