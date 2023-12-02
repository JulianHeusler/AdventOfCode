package adventofcode.util;

import java.io.File;
import java.util.Scanner;

public class Parser {
    public static String readInputFile(int day) {
        StringBuilder input = new StringBuilder();
        try {
            File file = new File(String.format("src/main/resources/day%02d/input.txt", day));
            Scanner scanner = new Scanner(file);
            while (scanner.hasNextLine()) {
                input.append(scanner.nextLine()).append("\n");
            }
            scanner.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return input.toString();
    }
}
