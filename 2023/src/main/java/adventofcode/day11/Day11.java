package adventofcode.day11;

import adventofcode.util.AbstractDay;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.LongStream;

public class Day11 extends AbstractDay {

    private static boolean isEmptyRow(long rowNumber, List<Position> galaxies) {
        return galaxies.stream()
                .map(Position::y)
                .noneMatch(y -> y == rowNumber);
    }

    private static boolean isEmptyColumn(long columnNumber, List<Position> galaxies) {
        return galaxies.stream()
                .map(Position::x)
                .noneMatch(x -> x == columnNumber);
    }

    @Override
    public long solvePart1(String input) {
        return sumLengthBetweenGalaxies(input, 2);
    }

    @Override
    public long solvePart2(String input) {
        return sumLengthBetweenGalaxies(input, 1000000);
    }

    private long sumLengthBetweenGalaxies(String input, int weight) {
        List<Position> galaxies = getGalaxiesPositions(input);
        List<Position> galaxiesCopy = new ArrayList<>(galaxies);

        long lengths = 0;
        for (Position galaxy1 : galaxies) {
            galaxiesCopy.remove(galaxy1);
            for (Position galaxy2 : galaxiesCopy) {
                lengths += calculateDistance(galaxy1, galaxy2, galaxies, weight);
            }
        }
        return lengths;
    }

    private long calculateDistance(Position a, Position b, List<Position> galaxies, int weight) {
        long dx = Math.abs(a.x - b.x);
        long dy = Math.abs(a.y - b.y);

        long emptyColumns = LongStream.range(Math.min(a.x, b.x), Math.max(a.x, b.x))
                .filter(i -> isEmptyColumn(i, galaxies))
                .count();
        long emptyRows = LongStream.range(Math.min(a.y, b.y), Math.max(a.y, b.y))
                .filter(i -> isEmptyRow(i, galaxies))
                .count();

        return dx + dy - emptyColumns + emptyColumns * weight - emptyRows + emptyRows * weight;
    }

    private List<Position> getGalaxiesPositions(String input) {
        List<Position> galaxies = new ArrayList<>();
        String[] split = input.split("\n");
        for (int y = 0; y < split.length; y++) {
            for (int x = 0; x < split[0].length(); x++) {
                char current = split[y].charAt(x);
                if (current == '#') {
                    galaxies.add(new Position(y, x));
                }
            }
        }
        return galaxies;
    }

    record Position(long y, long x) {
    }
}
