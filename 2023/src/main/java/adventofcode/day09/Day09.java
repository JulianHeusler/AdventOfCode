package adventofcode.day09;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Day09 extends AbstractDay {


    @Override
    public long solvePart1(String input) {
        return parseHistories(input).stream()
                .map(this::predictNextFutureValue)
                .mapToInt(Integer::intValue)
                .sum();
    }

    @Override
    public long solvePart2(String input) {
        return parseHistories(input).stream()
                .map(this::predictNextPastValue)
                .mapToInt(Integer::intValue)
                .sum();
    }

    private int predictNextFutureValue(List<Integer> history) {
        return getSequences(history)
                .reversed()
                .stream()
                .map(List::getLast)
                .reduce(Integer::sum)
                .orElseThrow();
    }

    private int predictNextPastValue(List<Integer> history) {
        return getSequences(history).reversed()
                .stream()
                .map(List::getFirst)
                .mapToInt(Integer::intValue)
                .reduce((a, b) -> b - a)
                .orElseThrow();
    }

    private List<List<Integer>> getSequences(List<Integer> history) {
        List<List<Integer>> sequences = new ArrayList<>();
        List<Integer> nextSequence = history;
        while (!isEmpty(nextSequence)) {
            sequences.add(nextSequence);
            nextSequence = getSequence(nextSequence);
        }
        sequences.forEach(System.out::println);
        return sequences;
    }

    private boolean isEmpty(List<Integer> nextSequence) {
        return nextSequence.stream().allMatch(i -> i == 0);
    }

    private List<Integer> getSequence(List<Integer> line) {
        List<Integer> newSequence = new ArrayList<>();
        if (line.size() == 1) {
            return List.of(0);
        }
        for (int i = 0; i < line.size() - 1; i++) {
            newSequence.add(line.get(i + 1) - line.get(i));
        }
        return newSequence;
    }

    private List<List<Integer>> parseHistories(String input) {
        return Arrays.stream(input.split("\n"))
                .map(ParseUtil::parseIntegerNumbers)
                .toList();
    }
}
