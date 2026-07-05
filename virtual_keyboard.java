// virtual_keyboard.java
import java.io.*;
import java.util.*;

public class virtual_keyboard {
    private static final String RESET = "\u001B[0m";
    private static final String GREEN = "\u001B[92m";
    private static final String BLUE = "\u001B[94m";
    private static final String YELLOW = "\u001B[93m";
    private static final String BOLD = "\u001B[1m";

    private static String colorize(String text, String color) {
        return color + text + RESET;
    }

    private static final List<List<String>> KEYBOARD_EN = Arrays.asList(
        Arrays.asList("1","2","3","4","5","6","7","8","9","0"),
        Arrays.asList("Q","W","E","R","T","Y","U","I","O","P"),
        Arrays.asList("A","S","D","F","G","H","J","K","L"),
        Arrays.asList("Z","X","C","V","B","N","M")
    );

    private static final List<List<String>> KEYBOARD_RU = Arrays.asList(
        Arrays.asList("1","2","3","4","5","6","7","8","9","0"),
        Arrays.asList("Й","Ц","У","К","Е","Н","Г","Ш","Щ","З"),
        Arrays.asList("Ф","Ы","В","А","П","Р","О","Л","Д"),
        Arrays.asList("Я","Ч","С","М","И","Т","Ь")
    );

    private String lang;
    private List<String> text = new ArrayList<>();
    private boolean shift;
    private List<List<String>> layout;

    public virtual_keyboard(String lang) {
        this.lang = lang;
        this.shift = false;
        this.layout = lang.equals("en") ? KEYBOARD_EN : KEYBOARD_RU;
    }

    public void display() {
        System.out.print("\033[H\033[2J");
        System.out.flush();
        System.out.println(colorize("🖥️  ВИРТУАЛЬНАЯ КЛАВИАТУРА", BOLD));
        System.out.println(colorize("Управление: номер клавиши → добавить, 0 → пробел, - → удалить, . → выход", YELLOW));
        System.out.println(colorize("8 → переключить раскладку, 9 → переключить регистр (Shift)\n", YELLOW));
        System.out.println(colorize("Текст: " + String.join("", text), GREEN));
        System.out.println();

        int idx = 1;
        for (List<String> row : layout) {
            String line = "";
            for (String ch : row) {
                String displayCh = shift ? ch.toUpperCase() : ch.toLowerCase();
                line += colorize(String.format("%02d", idx), BLUE) + ":" + colorize(displayCh, BLUE) + "  ";
                idx++;
            }
            System.out.println(line);
        }
        System.out.println("\nСпециальные: 0-пробел, --удалить, .-выход, 8-язык, 9-Shift");
    }

    public void run() throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        while (true) {
            display();
            System.out.print(colorize("Введите номер команды: ", BOLD));
            String cmd = reader.readLine().trim();
            if (cmd.equals(".")) break;
            if (cmd.equals("-")) {
                if (!text.isEmpty()) text.remove(text.size() - 1);
                continue;
            }
            if (cmd.equals("0")) {
                text.add(" ");
                continue;
            }
            if (cmd.equals("8")) {
                lang = lang.equals("en") ? "ru" : "en";
                layout = lang.equals("en") ? KEYBOARD_EN : KEYBOARD_RU;
                continue;
            }
            if (cmd.equals("9")) {
                shift = !shift;
                continue;
            }
            try {
                int num = Integer.parseInt(cmd);
                List<String> flat = new ArrayList<>();
                for (List<String> row : layout) flat.addAll(row);
                if (num >= 1 && num <= flat.size()) {
                    String ch = flat.get(num - 1);
                    ch = shift ? ch.toUpperCase() : ch.toLowerCase();
                    text.add(ch);
                } else {
                    System.out.println(colorize("Неверный номер!", "red"));
                }
            } catch (NumberFormatException e) {
                System.out.println(colorize("Неверный ввод!", "red"));
            }
        }
        System.out.println(colorize("\nИтоговый текст: " + String.join("", text), GREEN));
        reader.close();
    }

    public static void main(String[] args) throws IOException {
        virtual_keyboard kb = new virtual_keyboard("en");
        kb.run();
    }
}
