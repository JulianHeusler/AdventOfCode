package adventofcode.day03;

import adventofcode.util.AbstractDay;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public class Day03 extends AbstractDay {
    @Override
    public long solvePart1(String input) {
        return getPartNumbers(input).stream().mapToInt(Integer::intValue).sum();
    }

    @Override
    public long solvePart2(String input) {
        List<Integer> gearRatios = getPartNumbersAdjacentToGear(input)
                .stream()
                .collect(Collectors.groupingBy(PartNumber::gearPosition))
                .values()
                .stream()
                .filter(partNumbersPerGear -> partNumbersPerGear.size() == 2)
                .map(partNumbersPerGear -> partNumbersPerGear
                        .stream()
                        .mapToInt(PartNumber::number)
                        .reduce(1, Math::multiplyExact))
                .toList();
        return gearRatios.stream().mapToInt(Integer::intValue).sum();
    }

    private List<PartNumber> getPartNumbersAdjacentToGear(String input) {
        List<PartNumber> partNumbers = new ArrayList<>();
        StringBuilder currentNumber = new StringBuilder();
        Position gearPosition = null;

        String[] matrix = input.split("\n");
        for (int y = 0; y < matrix.length; y++) {
            for (int x = 0; x < matrix[y].length(); x++) {
                char current = matrix[y].charAt(x);
                if (Character.isDigit(current)) {
                    currentNumber.append(current);
                    Position position = isAdjacentToGear(matrix, y, x);
                    if (position != null) {
                        gearPosition = position;
                    }
                } else {
                    if (gearPosition != null) {
                        partNumbers.add(new PartNumber(Integer.parseInt(currentNumber.toString()), gearPosition));
                        gearPosition = null;
                    }
                    currentNumber = new StringBuilder();
                }
            }
            // at the end of each row
            if (gearPosition != null) {
                partNumbers.add(new PartNumber(Integer.parseInt(currentNumber.toString()), gearPosition));
                gearPosition = null;
            }
            currentNumber = new StringBuilder();
        }
        return partNumbers;
    }

    private List<Integer> getPartNumbers(String input) {
        List<Integer> partNumbers = new ArrayList<>();
        StringBuilder currentNumber = new StringBuilder();
        boolean adjacentToSymbol = false;

        String[] matrix = input.split("\n");
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
            // at the end of each row
            if (adjacentToSymbol) {
                partNumbers.add(Integer.valueOf(currentNumber.toString()));
                adjacentToSymbol = false;
            }
            currentNumber = new StringBuilder();
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

    private Position isAdjacentToGear(String[] matrix, int y, int x) {
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
                if (currentChar == '*') {
                    return new Position(newY, newX);
                }
            }
        }
        return null;
    }

    private boolean isInBounds(String[] matrix, int y, int x) {
        return 0 <= y && y < matrix.length && 0 <= x && x < matrix[0].length();
    }

    record PartNumber(int number, Position gearPosition) {
    }

    record Position(int y, int x) {
    }
}
