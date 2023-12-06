package adventofcode.day06;

import adventofcode.util.AbstractDay;
import adventofcode.util.Parser;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day06Test {
    private final AbstractDay day = new Day06();
    private static final int DAY_NUMBER = 6;

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
        assertEquals(0, day.solvePart1(Parser.readInputFile(DAY_NUMBER)));
    }

    @Test
    void testInputPart2() {
        String testInput = """
                Time:      7  15   30
                Distance:  9  40  200
                """;
        assertEquals(0, day.solvePart2(testInput));
    }

    @Test
    void testRealInputPart2() {
        assertEquals(0, day.solvePart2(Parser.readInputFile(DAY_NUMBER)));
    }
}