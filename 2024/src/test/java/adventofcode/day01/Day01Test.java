package adventofcode.day01;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import adventofcode.util.ParseUtil;

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
		String testInput = """
				3   4
				4   3
				2   5
				1   3
				3   9
				3   3
				""";

		assertEquals(31, day01.solvePart2(testInput));
	}

	@Test
	void testRealInputPart2() {
		assertEquals(22565391, day01.solvePart2(ParseUtil.readInputFile(1)));
	}
}