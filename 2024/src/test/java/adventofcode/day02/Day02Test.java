package adventofcode.day02;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import adventofcode.util.ParseUtil;

class Day02Test {

	private final Day02 day02 = new Day02();

	@Test
	void testInputPart1() {
		String testInput = """
				7 6 4 2 1
				1 2 7 8 9
				9 7 6 2 1
				1 3 2 4 5
				8 6 4 4 1
				1 3 6 7 9
				""";

		assertEquals(2, day02.solvePart1(testInput));
	}

	@Test
	void testRealInputPart1() {
		assertEquals(0, day02.solvePart1(ParseUtil.readInputFile(1)));
	}

	@Test
	void testInputPart2() {
		String testInput = """
				7 6 4 2 1
				1 2 7 8 9
				9 7 6 2 1
				1 3 2 4 5
				8 6 4 4 1
				1 3 6 7 9
				""";

		assertEquals(0, day02.solvePart2(testInput));
	}

	@Test
	void testRealInputPart2() {
		assertEquals(0, day02.solvePart2(ParseUtil.readInputFile(1)));
	}
}