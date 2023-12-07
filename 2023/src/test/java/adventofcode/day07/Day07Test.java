package adventofcode.day07;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day07Test {
    private final AbstractDay day = new Day07();
    private static final int DAY_NUMBER = 7;

    @Test
    void testInputPart1() {
        String testInput = """
                32T3K 765
                T55J5 684
                KK677 28
                KTJJT 220
                QQQJA 483
                """;
        assertEquals(6440, day.solvePart1(testInput));
    }

    @Test
    void testInputPart1dadw() {
        String testInput = """
                22234 2
                22233 3
                """;
        assertEquals(8, day.solvePart1(testInput));
    }

    @Test
    void testRealInputPart1() {
        assertEquals(251055711, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
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