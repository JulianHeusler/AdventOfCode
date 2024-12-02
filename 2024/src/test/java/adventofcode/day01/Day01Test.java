package adventofcode.day01;

import adventofcode.util.ParseUtil;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day01Test {

    private final Day01 day01 = new Day01();

    @Test
    void testInputPart1() {
        String testInput = """
                3   4
                4   3
                2   5
                1   3
                3   9
                3   3
                """;

        assertEquals(11, day01.solvePart1(testInput));
    }

    @Test
    void testRealInputPart1() {
        assertEquals(3574690, day01.solvePart1(ParseUtil.readInputFile(1)));
    }

    @Test
    void testInputPart2() {
        assertEquals(0, day01.solvePart2(
                """
                        
                        """));
    }

    @Test
    void testRealInputPart2() {
        assertEquals(0, day01.solvePart2(ParseUtil.readInputFile(1)));
    }
}