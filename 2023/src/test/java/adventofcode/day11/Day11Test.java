package adventofcode.day11;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day11Test {
    private final AbstractDay day = new Day11();
    private static final int DAY_NUMBER = 11;

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
        assertEquals(0, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
    }

    @Test
    void testInputPart2() {
        String testInput = """
                                
                """;
        assertEquals(0, day.solvePart2(testInput));
    }

    @Test
    void testRealInputPart2() {
        assertEquals(0, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
    }
}