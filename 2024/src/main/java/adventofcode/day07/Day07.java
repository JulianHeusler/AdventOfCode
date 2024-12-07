package adventofcode.day07;

import java.util.ArrayList;
import java.util.List;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

public class Day07 extends AbstractDay {

    private enum Operator {
        ADD,
        MULTIPLY,
        CONCATENATION;

        private long apply(long firstNumber, long secondNumber) {
            return switch (this) {
                case ADD -> firstNumber + secondNumber;
                case MULTIPLY -> firstNumber * secondNumber;
                case CONCATENATION -> Long.valueOf(String.valueOf(firstNumber) + String.valueOf(secondNumber));
            };
        }
    }

    @Override
    public long solvePart1(String input) {
        long sum = 0;
        for (Callibritation callibritation : parseCallibritations(input)) {
            if (isValid(callibritation, List.of(Operator.ADD, Operator.MULTIPLY))) {
                sum += callibritation.result();
            }
        }
        return sum;
    }

    private boolean isValid(Callibritation callibritation, List<Operator> operators) {
        return testRecursive(callibritation.result(), operators, removeFirst(callibritation.numbers()),
                callibritation.numbers().getFirst());
    }

    private boolean testRecursive(long result, List<Operator> operators, List<Long> numbers, long current) {
        if (numbers.isEmpty()) {
            return current == result;
        }

        if (current > result) {
            return false;
        }

        return operators.stream()
                .anyMatch(operator -> testRecursive(result, operators,
                        removeFirst(numbers), operator.apply(current, numbers.getFirst())));
    }

    @Override
    public long solvePart2(String input) {
        long sum = 0;
        for (Callibritation validCallibritation : parseCallibritations(input).stream()
                .filter(callibritation -> isValid(callibritation,
                        List.of(Operator.ADD, Operator.MULTIPLY, Operator.CONCATENATION)))
                .toList()) {
            sum += validCallibritation.result();
        }
        return sum;
    }

    private List<Long> removeFirst(List<Long> list) {
        if (list.isEmpty()) {
            return list;
        }
        return list.subList(1, list.size());
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

    private record Callibritation(long result, List<Long> numbers) {
    }
}
