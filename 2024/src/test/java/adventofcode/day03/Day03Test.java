package adventofcode.day03;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

public class Day03Test {

	private static final int DAY_NUMBER = 3;
	private final AbstractDay day = new Day03();

	@Test
	void testInputPart1() {
		String testInput = """
				xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
				""";
		assertEquals(161, day.solvePart1(testInput));
	}

	@Test
	void testRealInputPart1() {
		assertEquals(174561379, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
	}

	@Test
	void testInputPart2() {
		String testInput = """
				xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
				""";
		assertEquals(48, day.solvePart2(testInput));
	}

	@Test
	void testRealInputPart2() {
		assertEquals(106921067, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
	}
}