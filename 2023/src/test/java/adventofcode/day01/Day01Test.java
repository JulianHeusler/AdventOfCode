package adventofcode.day01;

import adventofcode.util.Parser;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day01Test {

    private final Day01 day01 = new Day01();

    @Test
    void testInputPart1() {
        String testInput = """
                1abc2
                pqr3stu8vwx
                a1b2c3d4e5f
                treb7uchet
                    """;

        assertEquals(142, day01.solvePart1(testInput));
    }

    @Test
    void testRealInputPart1() {
        assertEquals(55607, day01.solvePart1(Parser.readInputFile(1)));
    }

    @Test
    void testInputPart2() {
        assertEquals(281, day01.solvePart2(
                """
                        two1nine
                        eightwothree
                        abcone2threexyz
                        xtwone3four
                        4nineeightseven2
                        zoneight234
                        7pqrstsixteen
                        """));
    }

    @Test
    void testRealInputPart2() {
        assertEquals(55291, day01.solvePart2(Parser.readInputFile(1)));
    }
}