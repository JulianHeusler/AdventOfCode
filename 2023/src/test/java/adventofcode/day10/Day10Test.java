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
                ...........
                .S-------7.
                .|F-----7|.
                .||.....||.
                .||.....||.
                .|L-7.F-J|.
                .|..|.|..|.
                .L--J.L--J.
                ...........
                """;
        assertEquals(4, day.solvePart2(testInput));
    }

    @Test
    void testInputPart2Enclosed() {
        String testInput = """
                ..........
                .S------7.
                .|F----7|.
                .||OOOO||.
                .||OOOO||.
                .|L-7F-J|.
                .|II||II|.
                .L--JL--J.
                ..........
                """;
        assertEquals(4, day.solvePart2(testInput));
    }

    @Test
    void testInputPart2Larger() {
        String testInput = """
                .F----7F7F7F7F-7....
                .|F--7||||||||FJ....
                .||.FJ||||||||L7....
                FJL7L7LJLJ||LJ.L-7..
                L--J.L7...LJS7F-7L7.
                ....F-J..F7FJ|L7L7L7
                ....L7.F7||L7|.L7L7|
                .....|FJLJ|FJ|F7|.LJ
                ....FJL-7.||.||||...
                ....L---J.LJ.LJLJ...
                """;
        assertEquals(8, day.solvePart2(testInput));
    }

    @Test
    void testInputPart2LargerWithJunkPipes() {
        String testInput = """
                FF7FSF7F7F7F7F7F---7
                L|LJ||||||||||||F--J
                FL-7LJLJ||||||LJL-77
                F--JF--7||LJLJ7F7FJ-
                L---JF-JLJ.||-FJLJJ7
                |F|F-JF---7F7-L7L|7|
                |FFJF7L7F-JF7|JL---7
                7-L-JL7||F7|L7F-7F7|
                L.L7LFJ|||||FJL7||LJ
                L7JLJL-JLJLJL--JLJ.L
                """;
        assertEquals(10, day.solvePart2(testInput));
    }

    @Test
    void testInputPart2dawdaw() {
        assertEquals(1, day.solvePart2("""
                .....
                .....
                .S-7.
                .|-|.
                .L-J.
                .....
                """));
    }

    @Test
    void testRealInputPart2() {
        assertEquals(0, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
    }
}