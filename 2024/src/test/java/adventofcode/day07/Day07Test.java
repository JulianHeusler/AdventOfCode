package adventofcode.day07;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

public class Day07Test {

    private static final int DAY_NUMBER = 7;
    private final AbstractDay day = new Day07();

    @Test
    void testInputPart1() {
        String testInput = """
                190: 10 19
                3267: 81 40 27
                83: 17 5
                156: 15 6
                7290: 6 8 6 15
                161011: 16 10 13
                192: 17 8 14
                21037: 9 7 18 13
                292: 11 6 16 20
                """;
        assertEquals(3749, day.solvePart1(testInput));
    }

    @Test
    void testInputPart1_extra() {
        assertEquals(0, day.solvePart1(
                """
                            190: 10
                        """));
        assertEquals(0, day.solvePart1(
                """
                            190: 200
                        """));
        assertEquals(190, day.solvePart1(
                """
                            190: 190
                        """));
        assertEquals(100, day.solvePart1(
                """
                            100: 50, 50
                        """));
        assertEquals(100, day.solvePart1(
                """
                            100: 50, 2
                        """));
        assertEquals(0, day.solvePart1(
                """
                            100: 50, 20
                        """));
        assertEquals(400, day.solvePart1(
                """
                            100: 50, 2
                            100: 50, 2
                            100: 50, 2
                            100: 50, 2
                        """));
    }

    @Test
    void testRealInputPart1() {
        assertEquals(303766880536L, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
    }

    @Test
    void testInputPart2() {
        String testInput = """
                190: 10 19
                3267: 81 40 27
                83: 17 5
                156: 15 6
                7290: 6 8 6 15
                161011: 16 10 13
                192: 17 8 14
                21037: 9 7 18 13
                292: 11 6 16 20
                """;
        assertEquals(0, day.solvePart2(testInput));
    }

    @Test
    void testRealInputPart2() {
        assertEquals(0, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
    }
}
