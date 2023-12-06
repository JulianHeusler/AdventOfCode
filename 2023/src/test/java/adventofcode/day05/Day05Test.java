package adventofcode.day05;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class Day05Test {

    private final AbstractDay day = new Day05();
    private static final int DAY_NUMBER = 5;

    @Test
    void testInput() {
        String testInput = """
                seeds: 79 14 55 13
                                
                seed-to-soil map:
                50 98 2
                52 50 48
                				
                soil-to-fertilizer map:
                0 15 37
                37 52 2
                39 0 15
                				
                fertilizer-to-water map:
                49 53 8
                0 11 42
                42 0 7
                57 7 4
                				
                water-to-light map:
                88 18 7
                18 25 70
                				
                light-to-temperature map:
                45 77 23
                81 45 19
                68 64 13
                				
                temperature-to-humidity map:
                0 69 1
                1 0 69
                				
                humidity-to-location map:
                60 56 37
                56 93 4
                """;

        assertEquals(35, day.solvePart1(testInput));
        assertEquals(46, day.solvePart2(testInput));
    }

    @Test
    void testRealInputPart1() {
        assertEquals(535088217, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
    }

    @Test
    void testRealInputPart2() {
        assertEquals(0, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
    }


}