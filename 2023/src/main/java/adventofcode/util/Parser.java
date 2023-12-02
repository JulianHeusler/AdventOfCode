package adventofcode.util;

import java.io.File;
import java.util.Scanner;

public final class Parser {

    private Parser() {
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
}
