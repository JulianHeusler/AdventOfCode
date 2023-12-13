package adventofcode.day12;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

class Day12Test {

	private static final int DAY_NUMBER = 12;
	private final AbstractDay day = new Day12();

	@Test
	void testInputPart1() {
		String testInput = """
				#.#.### 1,1,3
				.#...#....###. 1,1,3
				.#.###.#.###### 1,3,1,6
				####.#...#... 4,1,1
				#....######..#####. 1,6,5
				.###.##....# 3,2,1
				""";
		assertEquals(6, day.solvePart1(testInput));
	}

	@Test
	void testInputPart1234234() {
		assertEquals(21, day.solvePart1("""
				???.### 1,1,3
				.??..??...?##. 1,1,3
				?#?#?#?#?#?#?#? 1,3,1,6
				????.#...#... 4,1,1
				????.######..#####. 1,6,5
				?###???????? 3,2,1
				"""));
	}

	@Test
	void testRealInputPart1() {
		assertEquals(7017, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
	}

	@Test
	void testInputPart2() {
		String testInput = """
				0 3 6 9 12 15
				1 3 6 10 15 21
				10 13 16 21 30 45
				""";
		assertEquals(0, day.solvePart2(testInput));
	}

	@Test
	void testRealInputPart2() {
		assertEquals(0, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
	}
}