package adventofcode.day04;

import adventofcode.util.AbstractDay;
import adventofcode.util.Parser;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day04Test {

    private final AbstractDay day = new Day04();
    private static final int DAY_NUMBER = 4;

    @Test
    void testInputPart1() {
        String testInput = """
                Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
                Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
                Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
                Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
                Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
                Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
                """;

        assertEquals(13, day.solvePart1(testInput));
    }

    @Test
    void testRealInputPart1() {
        assertEquals(17803, day.solvePart1(Parser.readInputFile(DAY_NUMBER)));
    }

    @Test
    void testInputPart2() {
        String testInput = """
                Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
                Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
                Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
                Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
                Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
                Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
                """;
        assertEquals(30, day.solvePart2(testInput));
    }

    @Test
    void testRealInputPart2() {
        assertEquals(0, day.solvePart2(Parser.readInputFile(DAY_NUMBER)));
    }

}