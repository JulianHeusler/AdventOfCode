package adventofcode.util;

import java.io.File;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public final class ParseUtil {

    private ParseUtil() {
    }

    @SuppressWarnings("CallToPrintStackTrace")
    public static String readInputFile(int day) {
        StringBuilder input = new StringBuilder();
        try (Scanner scanner = new Scanner(new File(String.format("src/main/resources/day%02d/input.txt", day)))) {
            while (scanner.hasNextLine()) {
                input.append(scanner.nextLine()).append("\n");
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return input.toString();
    }

    public static List<Long> parseLongNumbers(String numberLine) {
        List<Long> numbers = new ArrayList<>();
        Matcher matcher = Pattern.compile("(\\d+)").matcher(numberLine);
        while (matcher.find()) {
            numbers.add(Long.parseLong(matcher.group()));
        }
        return numbers;
    }

    public static List<Integer> parseIntegerNumbers(String numberLine) {
        List<Integer> numbers = new ArrayList<>();
        Matcher matcher = Pattern.compile("(\\d+)").matcher(numberLine);
        while (matcher.find()) {
            numbers.add(Integer.parseInt(matcher.group()));
        }
        return numbers;
    }
}
