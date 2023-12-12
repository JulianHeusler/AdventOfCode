package adventofcode.day08;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

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


	@Test
	void testInputPart2() {
		String testInput = """
				LR
				                
				11A = (11B, XXX)
				11B = (XXX, 11Z)
				11Z = (11B, XXX)
				22A = (22B, XXX)
				22B = (22C, 22C)
				22C = (22Z, 22Z)
				22Z = (22B, 22B)
				XXX = (XXX, XXX)
				""";
		assertEquals(6, day.solvePart2(testInput));
	}

	@Disabled
	@Test
	void testRealInputPart2() {
		assertEquals(0, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
	}
}