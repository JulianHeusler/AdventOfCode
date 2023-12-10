package adventofcode.day10;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day10Test {
    private final AbstractDay day = new Day10();
    private static final int DAY_NUMBER = 10;

    @Test
    void testInputPart1() {
        assertEquals(4, day.solvePart1("""
                .....
                .S-7.
                .|.|.
                .L-J.
                .....
                """));
        assertEquals(4, day.solvePart1("""
                -L|F7
                7S-7|
                L|7||
                -L-J|
                L|-JF
                """));
    }

    @Test
    void testInputPart12() {
        assertEquals(8, day.solvePart1("""
                ..F7.
                .FJ|.
                SJ.L7
                |F--J
                LJ...
                """));
        assertEquals(8, day.solvePart1("""
                7-F7-
                .FJ|7
                SJLL7
                |F--J
                LJ.LJ
                """));
    }

    @Test
    void testRealInputPart1() {
        assertEquals(6823, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
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