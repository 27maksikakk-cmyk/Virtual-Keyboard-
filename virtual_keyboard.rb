#!/usr/bin/env ruby
# virtual_keyboard.rb
# encoding: UTF-8

require 'io/console'

COLORS = {
  reset: "\e[0m",
  green: "\e[92m",
  blue: "\e[94m",
  yellow: "\e[93m",
  bold: "\e[1m"
}

def colorize(text, color)
  "#{COLORS[color]}#{text}#{COLORS[:reset]}"
end

KEYBOARD_EN = [
  %w[1 2 3 4 5 6 7 8 9 0],
  %w[Q W E R T Y U I O P],
  %w[A S D F G H J K L],
  %w[Z X C V B N M]
]

KEYBOARD_RU = [
  %w[1 2 3 4 5 6 7 8 9 0],
  %w[Й Ц У К Е Н Г Ш Щ З],
  %w[Ф Ы В А П Р О Л Д],
  %w[Я Ч С М И Т Ь]
]

class VirtualKeyboard
  attr_accessor :lang, :text, :shift, :layout

  def initialize(lang = 'en')
    @lang = lang
    @text = []
    @shift = false
    @layout = lang == 'en' ? KEYBOARD_EN : KEYBOARD_RU
  end

  def display
    system('clear') || system('cls')
    puts colorize('🖥️  ВИРТУАЛЬНАЯ КЛАВИАТУРА', :bold)
    puts colorize('Управление: номер клавиши → добавить, 0 → пробел, - → удалить, . → выход', :yellow)
    puts colorize('8 → переключить раскладку, 9 → переключить регистр (Shift)\n', :yellow)
    puts colorize("Текст: #{@text.join}", :green)
    puts

    idx = 1
    @layout.each do |row|
      line = ''
      row.each do |ch|
        display_ch = @shift ? ch.upcase : ch.downcase
        line += colorize(idx.to_s.rjust(2, '0'), :blue) + ':' + colorize(display_ch, :blue) + '  '
        idx += 1
      end
      puts line
    end
    puts "\nСпециальные: 0-пробел, --удалить, .-выход, 8-язык, 9-Shift"
  end

  def run
    loop do
      display
      print colorize('Введите номер команды: ', :bold)
      cmd = gets.chomp.strip
      case cmd
      when '.'
        break
      when '-'
        @text.pop if @text.any?
      when '0'
        @text << ' '
      when '8'
        @lang = @lang == 'en' ? 'ru' : 'en'
        @layout = @lang == 'en' ? KEYBOARD_EN : KEYBOARD_RU
      when '9'
        @shift = !@shift
      else
        num = cmd.to_i
        if num > 0
          flat = @layout.flatten
          if num <= flat.size
            ch = flat[num-1]
            ch = @shift ? ch.upcase : ch.downcase
            @text << ch
          else
            puts colorize('Неверный номер!', :red)
          end
        else
          puts colorize('Неверный ввод!', :red)
        end
      end
    end
    puts colorize("\nИтоговый текст: #{@text.join}", :green)
  end
end

if __FILE__ == $0
  kb = VirtualKeyboard.new('en')
  kb.run
end
