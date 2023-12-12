package adventofcode.day01;

import adventofcode.util.AbstractDay;

import java.util.Arrays;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day01 extends AbstractDay {

    private static final Pattern NUMBER_START = Pattern.compile("^(zero|one|two|three|four|five|six|seven|eight|nine|ten).*");
    private static final Pattern NUMBER_END = Pattern.compile(".*(zero|one|two|three|four|five|six|seven|eight|nine|ten)$");

    public long solvePart1(String input) {
        return Arrays.stream(input.split("\n")).mapToInt(this::getBothDigits).sum();
    }

    private int getBothDigits(String line) {
        int[] array = line.chars().filter(Character::isDigit).toArray();
        return Integer.parseInt(Character.toString(array[0]) + Character.toString(array[array.length - 1]));
    }

    public long solvePart2(String input) {
        return Arrays.stream(input.split("\n")).mapToInt(line -> getFirstNumberValue(line) * 10 + getLastNumberValue(line)).sum();
    }

    private int getFirstNumberValue(String line) {
        final char firstChar = line.charAt(0);
        if (Character.isDigit(firstChar)) {
            return Character.getNumericValue(firstChar);
        }
        Matcher matcher = NUMBER_START.matcher(line);
        if (matcher.matches()) {
            return wordToDigit(matcher.group(1));
        }
        return getFirstNumberValue(line.substring(1));
    }

    private int getLastNumberValue(String line) {
        final char lastChar = line.charAt(line.length() - 1);
        if (Character.isDigit(lastChar)) {
            return Character.getNumericValue(lastChar);
        }
        Matcher matcher = NUMBER_END.matcher(line);
        if (matcher.matches()) {
            return wordToDigit(matcher.group(1));
        }
        return getLastNumberValue(line.substring(0, line.length() - 1));
    }

    private int wordToDigit(String word) {
        return switch (word) {
            case "zero" -> 0;
            case "one" -> 1;
            case "two" -> 2;
            case "three" -> 3;
            case "four" -> 4;
            case "five" -> 5;
            case "six" -> 6;
            case "seven" -> 7;
            case "eight" -> 8;
            case "nine" -> 9;
            default -> -1;
        };
    }
}