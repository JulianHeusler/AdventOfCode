package adventofcode.day03;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day03Test {
    private final AbstractDay day = new Day03();
    private static final int DAY_NUMBER = 3;

    @Test
    void testInputPart1() {
        String testInput = """
                467..114..
                ...*......
                ..35..633.
                ......#...
                617*......
                .....+.58.
                ..592.....
                ......755.
                ...$.*....
                .664.598..
                """;

        assertEquals(4361, day.solvePart1(testInput));
        assertEquals(4, day.solvePart1("""
                ........
                .24..4..
                ......*.
                """));
        assertEquals(221, day.solvePart1("""
                ....221
                ......*
                """));

        assertEquals(503, day.solvePart1("503+"));
    }

    @Test
    void testRealInputPart1() {
        assertEquals(514969, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
    }

    @Test
    void testInput2() {
        assertEquals(925, day.solvePart1("""
                12.......*..
                +.........34
                .......-12..
                ..78........
                ..*....60...
                78.........9
                .5.....23..$
                8...90*12...
                ............
                2.2......12.
                .*.........*
                1.1..503+.56
                """));
    }

    @Test
    void testInput3() {
        assertEquals(5 + 7 + 13 + 15, day.solvePart1("""
                .......5......
                ..7*..*.......
                ...*13*.......
                .......15.....
                """));
    }

    @Test
    void testInput4() {
        assertEquals(4361, day.solvePart1("""
                467..114..
                ...*......
                ..35..633.
                ......#...
                617*......
                .....+.58.
                ..592.....
                ......755.
                ...$.*....
                .664.598..
                """));
    }

    @Test
    void lineWrap() {
        assertEquals(123, day.solvePart1("""
                .......123
                456....*..
                ..........
                """));
    }

    @Test
    void testInput5() {
        assertEquals(413, day.solvePart1("""
                12.......*..
                +.........34
                .......-12..
                ..78........
                ..*....60...
                78..........
                .......23...
                ....90*12...
                ............
                2.2......12.
                .*.........*
                1.1.......56
                """));
    }

    @Test
    void testInput6() {
        assertEquals(111 + 11 + 1, day.solvePart1("""
                ....#...#..#
                .111..11..1.
                ............
                """));
        assertEquals(111, day.solvePart1("""
                #...........
                .111..11..1.
                ............
                """));
    }


    @Test
    void testInputPart2() {
        String testInput = """
                467..114..
                ...*......
                ..35..633.
                ......#...
                617*......
                .....+.58.
                ..592.....
                ......755.
                ...$.*....
                .664.598..
                """;
        assertEquals(467835, day.solvePart2(testInput));
    }

    @Test
    void testRealInputPart2() {
        assertEquals(78915902, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
    }
}