package adventofcode.day03;

import adventofcode.util.AbstractDay;

import java.util.ArrayList;
import java.util.List;

public class Day03 extends AbstractDay {
    @Override
    public int solvePart1(String input) {
        System.out.println(foo(input));
        return foo(input).stream().mapToInt(Integer::intValue).sum();
    }

    @Override
    public int solvePart2(String input) {
        return 0;
    }

    private List<Integer> foo(String input) {
        List<Integer> partNumbers = new ArrayList<>();
        StringBuilder currentNumber = new StringBuilder();
        boolean adjacentToSymbol = false;

        String[] matrix = input.split("\n");
        for (String s : matrix) {
            assert s.length() == matrix[0].length();
        }

        for (int y = 0; y < matrix.length; y++) {
            for (int x = 0; x < matrix[y].length(); x++) {
                char current = matrix[y].charAt(x);
                if (Character.isDigit(current)) {
                    currentNumber.append(current);
                    if (isAdjacentToSymbol(matrix, y, x)) {
                        adjacentToSymbol = true;
                    }
                } else {
                    if (adjacentToSymbol) {
                        partNumbers.add(Integer.valueOf(currentNumber.toString()));
                        adjacentToSymbol = false;
                    }
                    currentNumber = new StringBuilder();
                }
            }
            if (adjacentToSymbol) {
                partNumbers.add(Integer.valueOf(currentNumber.toString()));
                adjacentToSymbol = false;
            }
            currentNumber = new StringBuilder();
        }

        if (adjacentToSymbol) {
            partNumbers.add(Integer.valueOf(currentNumber.toString()));
        }
        return partNumbers;
    }

    private boolean isAdjacentToSymbol(String[] matrix, int y, int x) {
        int[][] directions = {
                {-1, -1}, {-1, 0}, {-1, 1},
                {0, -1}, /*0, 0*/ {0, 1},
                {1, -1}, {1, 0}, {1, 1}
        };

        for (int[] direction : directions) {
            int newY = y + direction[0];
            int newX = x + direction[1];

            if (isInBounds(matrix, newY, newX)) {
                char currentChar = matrix[newY].charAt(newX);
                if (currentChar != '.' && !Character.isDigit(currentChar)) {
                    return true;
                }
            }
        }
        return false;
    }

    private boolean isInBounds(String[] matrix, int y, int x) {
        return 0 <= y && y < matrix.length && 0 <= x && x < matrix[0].length();
    }
}
