package adventofcode.day07;

import java.util.ArrayList;
import java.util.List;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

public class Day07 extends AbstractDay {

    private enum Operator {
        ADD,
        MULTIPLY
    }

    private record Callibritation(long result, List<Long> numbers) {
    }

    @Override
    public long solvePart1(String input) {
        long sum = 0;
        for (Callibritation validCallibritation : parseCallibritations(input).stream()
                .filter(this::isValid).toList()) {
            sum += validCallibritation.result();
        }
        return sum;
    }

    private boolean isValid(Callibritation callibritation) {
        return testRecursive(removeFirst(callibritation.numbers()), callibritation.numbers().getFirst(),
                callibritation.result());
    }

    private boolean testRecursive(List<Long> numbers, long current, long result) {
        if (numbers.isEmpty()) {
            return current == result;
        }

        if (current > result) {
            // return false;
        }

        long added = current + numbers.getFirst();
        long multiplyed = current * numbers.getFirst();

        return testRecursive(numbers.subList(1, numbers.size()), added, result)
                || testRecursive(numbers.subList(1, numbers.size()), multiplyed, result);
    }

    private List<Long> removeFirst(List<Long> list) {
        if (list.isEmpty()) {
            return list;
        }
        return list.subList(1, list.size());
    }

    @Override
    public long solvePart2(String input) {
        return 0;
    }

    private List<Callibritation> parseCallibritations(String input) {
        List<Callibritation> callibritations = new ArrayList<>();
        for (String line : input.split("\n")) {
            String[] split = line.split(": ");
            long result = Long.valueOf(split[0].trim());
            List<Long> numbers = ParseUtil.parseLongNumbers(split[1]);
            callibritations.add(new Callibritation(result, numbers));
        }
        return callibritations;
    }
}
