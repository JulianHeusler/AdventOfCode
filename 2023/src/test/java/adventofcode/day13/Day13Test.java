package adventofcode.day13;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day13Test {
    private static final int DAY_NUMBER = 13;
    private final AbstractDay day = new Day13();

    @Test
    void testInputPart1() {
        String testInput = """
                #.##..##.
                ..#.##.#.
                ##......#
                ##......#
                ..#.##.#.
                ..##..##.
                #.#.##.#.
                                
                #...##..#
                #....#..#
                ..##..###
                #####.##.
                #####.##.
                ..##..###
                #....#..#
                """;
        assertEquals(405, day.solvePart1(testInput));
    }

    @Test
    void testInputPart1232() {
        String testInput = """
                ###.#.###
                .#..####.
                .##.#.#..
                .##.#.#..
                .#..####.
                ###.#.###
                .#.#.#.##
                ######.##
                #.#....#.
                #.#....#.
                ######.##
                .#.#.#.##
                ###.#.###
                .#..####.
                .##.#....
                """;
        assertEquals(500, day.solvePart1(testInput));
    }

    @Test
    void testInputPart1Example() {
        String testInput = """
                #.##..##.
                ..#.##.#.
                ##......#
                ##......#
                ..#.##.#.
                ..##..##.
                #.#.##.#.
                                 
                #...##..#
                #....#..#
                ..##..###
                #####.##.
                #####.##.
                ..##..###
                #....#..#
                               
                .#.##.#.#
                .##..##..
                .#.##.#..
                #......##
                #......##
                .#.##.#..
                .##..##.#
                                 
                #..#....#
                ###..##..
                .##.#####
                .##.#####
                ###..##..
                #..#....#
                #..##...#
                """;
        assertEquals(709, day.solvePart1(testInput));
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