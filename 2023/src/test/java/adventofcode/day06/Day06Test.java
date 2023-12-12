package adventofcode.day06;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day06Test {
    private static final int DAY_NUMBER = 6;
    private final AbstractDay day = new Day06();

    @Test
    void testInputPart1() {
        String testInput = """
                Time:      7  15   30
                Distance:  9  40  200
                """;
        assertEquals(288, day.solvePart1(testInput));
    }

    @Test
    void testRealInputPart1() {
        assertEquals(3317888, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
    }

    @Test
    void testInputPart2() {
        String testInput = """
                Time:      7  15   30
                Distance:  9  40  200
                """;
        assertEquals(71503, day.solvePart2(testInput));
    }

    @Test
    void testRealInputPart2() {
        assertEquals(24655068, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
    }
}