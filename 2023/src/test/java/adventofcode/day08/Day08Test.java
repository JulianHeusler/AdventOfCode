package adventofcode.day08;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class Day08Test {
    private final AbstractDay day = new Day08();
    private static final int DAY_NUMBER = 8;

    @Test
    void testInputPart1() {
        String testInput = """
                RL
                                                                
                AAA = (BBB, CCC)
                BBB = (DDD, EEE)
                CCC = (ZZZ, GGG)
                DDD = (DDD, DDD)
                EEE = (EEE, EEE)
                GGG = (GGG, GGG)
                ZZZ = (ZZZ, ZZZ)
                """;
        assertEquals(2, day.solvePart1(testInput));
    }

    @Test
    void testAdditionalInputPart1() {
        String testInput = """
                LLR
                                
                AAA = (BBB, BBB)
                BBB = (AAA, ZZZ)
                ZZZ = (ZZZ, ZZZ)
                """;
        assertEquals(6, day.solvePart1(testInput));
    }

    @Test
    void testRealInputPart1() {
        assertEquals(17873, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
    }

}