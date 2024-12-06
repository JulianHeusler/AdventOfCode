package adventofcode.day06;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

public class Day06Test {
	private static final int DAY_NUMBER = 6;
	private final AbstractDay day = new Day06();

	@Test
	void testInputPart1() {
		String testInput = """
				....#.....
				.........#
				..........
				..#.......
				.......#..
				..........
				.#..^.....
				........#.
				#.........
				......#...
				""";
		assertEquals(41, day.solvePart1(testInput));
	}

	@Test
	void testRealInputPart1() {
		assertEquals(5095, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
	}

	@Test
	void testInputPart2() {
		String testInput = """
				....#.....
				.........#
				..........
				..#.......
				.......#..
				..........
				.#..^.....
				........#.
				#.........
				......#...
				""";
		assertEquals(0, day.solvePart2(testInput));
	}

	@Test
	void testRealInputPart2() {
		assertEquals(0, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
	}

}