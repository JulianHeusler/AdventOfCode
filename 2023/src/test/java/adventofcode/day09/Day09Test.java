package adventofcode.day09;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day09Test {
    private final AbstractDay day = new Day09();
    private static final int DAY_NUMBER = 9;

    @Test
    void testInputPart1() {
        String testInput = """
                0 3 6 9 12 15
                1 3 6 10 15 21
                10 13 16 21 30 45
                """;
        assertEquals(114, day.solvePart1(testInput));
    }

    @Test
    void testRealInputPart1() {
        assertEquals(2008960228, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
    }

    @Test
    void testInputPart2() {
        String testInput = """
                0 3 6 9 12 15
                1 3 6 10 15 21
                10 13 16 21 30 45
                """;
        assertEquals(2, day.solvePart2(testInput));
    }

    @Test
    void testRealInputPart2() {
        assertEquals(1097, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
    }

}