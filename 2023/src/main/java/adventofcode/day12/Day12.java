package adventofcode.day12;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

import java.util.ArrayList;
import java.util.List;

public class Day12 extends AbstractDay {
    @Override
    public long solvePart1(String input) {

        long result = 0;

        for (String line : input.split("\n")) {
            String[] split = line.split(" ");
            List<Integer> numbers = ParseUtil.parseIntegerNumbers(split[1]);

            result += getPermuations(split[0]).stream().filter(s -> isValidLine(s, numbers)).toList().size();
        }
        return result;
    }


    private List<String> getPermuations(String line) {
        List<String> perm = new ArrayList<>();
        long questionCount = line.chars().filter(c -> c == '?').count();

        for (int q = 0; q < questionCount; q++) {
            int count = (int) line.chars().count();
            for (int i = 0; i < count; i++) {
                if (line.charAt(i) == '?') {
                    if (i == 0) {
                        perm.add('.' + line.substring(i + 1, count));
                        perm.add('#' + line.substring(i + 1, count));
                    } else {
                        perm.add(line.substring(i - 1, i) + '.' + line.substring(i + 1, count));
                        perm.add(line.substring(i - 1, i) + '#' + line.substring(i + 1, count));
                    }
                }
            }
        }
        return perm;
    }

    private boolean isValidLine(String line, List<Integer> numbers) {
        String[] split = line.split(".");
        for (int i = 0; i < numbers.size(); i++) {
            if (split[i].length() != numbers.get(i)) {
                return false;
            }
        }
        return true;
    }

    @Override
    public long solvePart2(String input) {
        return 0;
    }
}
