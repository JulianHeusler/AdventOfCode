package adventofcode.day09;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

import java.util.ArrayList;
import java.util.List;

public class Day09 extends AbstractDay {


    @Override
    public int solvePart1(String input) {
        List<List<Integer>> histories = parseHistories(input);
        long result = 0;
        for (List<Integer> history : histories) {
            result += solve(history);
        }
        return Math.toIntExact(result);
    }

    private int solve(List<Integer> history) {
        List<List<Integer>> sequences = getSequences(history);
        sequences.forEach(System.out::println);

        int a = 0;
        for (List<Integer> sequence : sequences.reversed()) {
            a += sequence.getLast();
        }
        return a;
    }


    private List<List<Integer>> getSequences(List<Integer> history) {
        List<List<Integer>> sequences = new ArrayList<>();
        sequences.add(history);
        List<Integer> nextSequence = history;
        while (!isEmpty(nextSequence)) {
            nextSequence = getSequence(nextSequence);
            sequences.add(nextSequence);
        }
        return sequences;
    }

    private boolean isEmpty(List<Integer> nextSequence) {
        return nextSequence.stream().allMatch(i -> i == 0);
    }

    @Override
    public int solvePart2(String input) {
        return 0;
    }

    private List<Integer> getSequence(List<Integer> line) {
        List<Integer> newSequence = new ArrayList<>();
        if (line.size() == 1) {
            return List.of(0);
        }

        for (int i = 0; i < line.size() - 1; i++) {
            int x = line.get(i + 1) - line.get(i);
            newSequence.add(x);
        }
        return newSequence;
    }

    private List<List<Integer>> parseHistories(String input) {
        List<List<Integer>> histories = new ArrayList<>();
        for (String line : input.split("\n")) {
            histories.add(ParseUtil.parseIntegerNumbers(line));
        }
        return histories;
    }
}
