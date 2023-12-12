package adventofcode.day11;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day11Test {
    private static final int DAY_NUMBER = 11;
    private final AbstractDay day = new Day11();

    @Test
    void testInputPart1() {
        String testInput = """
                ...#......
                .......#..
                #.........
                ..........
                ......#...
                .#........
                .........#
                ..........
                .......#..
                #...#.....
                """;
        assertEquals(374, day.solvePart1(testInput));
    }

    @Test
    void testRealInputPart1() {
        assertEquals(9312968, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
    }

    @Test
    void testInputPart2() {
        String testInput = """
                ...#......
                .......#..
                #.........
                ..........
                ......#...
                .#........
                .........#
                ..........
                .......#..
                #...#.....
                """;
        assertEquals(82000210, day.solvePart2(testInput));
    }

    @Test
    void testRealInputPart2() {
        assertEquals(597714117556L, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
    }
}