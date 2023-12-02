package adventofcode.day01;

import java.io.File;
import java.util.Scanner;

public class Day01 {
    public static void main(String[] args) {
        String testInput = """
                1abc2
                pqr3stu8vwx
                a1b2c3d4e5f
                treb7uchet
                    """;
        String i = parse();
        System.out.println("Part 1: " + new Day01().solvePart1(i));
    }

    public int solvePart1(String input) {
        int sum = 0;
        for (String s : input.split("\n")) {
            sum += foo(s);
        }
        return sum;
    }

    private int foo(String line) {
        int[] array = line.chars().filter(Character::isDigit).toArray();

        String temp = Character.toString(array[0]) + Character.toString(array[array.length - 1]);
        return Integer.valueOf(temp);
    }

    private static String parse() {
        String input = "";
        try {
            File file = new File("src/main/resources/day01/input.txt");
            Scanner scanner = new Scanner(file);
            while (scanner.hasNextLine()) {
                input += scanner.nextLine() + "\n";
            }
            scanner.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return input;
    }
}