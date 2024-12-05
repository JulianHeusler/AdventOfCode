package adventofcode.day05;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

public class Day05Test {

	private static final int DAY_NUMBER = 5;
	private final AbstractDay day = new Day05();

	@Test
	void testInputPart1() {
		String testInput = """
				47|53
				97|13
				97|61
				97|47
				75|29
				61|13
				75|53
				29|13
				97|29
				53|29
				61|53
				97|53
				61|29
				47|13
				75|47
				97|75
				47|61
				75|61
				47|29
				75|13
				53|13
				
				75,47,61,53,29
				97,61,53,29,13
				75,29,13
				75,97,47,61,53
				61,13,29
				97,13,75,29,47
				""";
		assertEquals(143, day.solvePart1(testInput));
	}

	@Test
	void testRealInputPart1() {
		assertEquals(7074, day.solvePart1(ParseUtil.readInputFile(DAY_NUMBER)));
	}

	@Test
	void testInputPart2() {
		String testInput = """
				47|53
				97|13
				97|61
				97|47
				75|29
				61|13
				75|53
				29|13
				97|29
				53|29
				61|53
				97|53
				61|29
				47|13
				75|47
				97|75
				47|61
				75|61
				47|29
				75|13
				53|13
				
				75,47,61,53,29
				97,61,53,29,13
				75,29,13
				75,97,47,61,53
				61,13,29
				97,13,75,29,47
				""";
		assertEquals(123, day.solvePart2(testInput));
	}

	@Test
	void testRealInputPart2() {
		assertEquals(4828, day.solvePart2(ParseUtil.readInputFile(DAY_NUMBER)));
	}
}