package adventofcode.day09;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Day09 extends AbstractDay {


    @Override
    public int solvePart1(String input) {
        return parseHistories(input).stream()
                .map(this::predictNextValue)
                .mapToInt(Integer::intValue)
                .sum();
    }

    @Override
    public int solvePart2(String input) {
        return 0;
    }

    private int predictNextValue(List<Integer> history) {
        List<List<Integer>> sequences = getSequences(history);
        sequences.forEach(System.out::println);
        return sequences.reversed()
                .stream()
                .map(List::getLast)
                .mapToInt(Integer::intValue)
                .sum();
    }

    private List<List<Integer>> getSequences(List<Integer> history) {
        List<List<Integer>> sequences = new ArrayList<>();
        List<Integer> nextSequence = history;
        while (!isEmpty(nextSequence)) {
            sequences.add(nextSequence);
            nextSequence = getSequence(nextSequence);
        }
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
