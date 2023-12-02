package adventofcode.day01;

public class Day01 {

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
        return Integer.parseInt(temp);
    }
}