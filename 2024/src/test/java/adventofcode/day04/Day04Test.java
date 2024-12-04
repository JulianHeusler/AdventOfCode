package adventofcode.day04;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Arrays;
import java.util.stream.Stream;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import adventofcode.util.ParseUtil;

public class Day04Test {

	private final Day04 day04 = new Day04();

	@Test
	void testInputPart1() {
		String testInput = """
				MMMSXXMASM
				MSAMXMSMSA
				AMXSXMAAMM
				MSAMASMSMX
				XMASAMXAMM
				XXAMMXXAMA
				SMSMSASXSS
				SAXAMASAAA
				MAMMMXMMMM
				MXMXAXMASX
				""";

		assertEquals(18, day04.solvePart1(testInput));
	}

	@Test
	void testInputPart1X() {
		String testInput = """
				..X...
				.SAMX.
				.A..A.
				XMAS.S
				.X....
				""";

		assertEquals(4, day04.solvePart1(testInput));
	}


	private static Stream<Arguments> testData() {
		return Stream.of( //
				Arguments.of("""
						XXXX
						XXXM
						XXXX
						AXXX
						XSXX
						""", 0),
				Arguments.of("""
						XMAS
						XMAS
						XMAS
						XMAS
						""", 6),
				Arguments.of("""
						SAMX
						SAMX
						SAMX
						SAMX
						""", 6),
				Arguments.of("""
						XXXX
						MMMM
						AAAA
						SSSS
						""", 6),
				Arguments.of("""
						XXXX
						MMMM
						AAAA
						SSSS
						""", 6),
				Arguments.of("""
						XXXX
						MMMM
						AAAA
						SSSS
						""", 6),
				Arguments.of("""
						SSSS
						AAAA
						MMMM
						XXXX
						""", 6),
				Arguments.of("""
						XXXX
						XMMM
						XMAA
						XMAS
						""", 3),
				Arguments.of("""
						XXXX
						MMMX
						AAMX
						SAMX
						""", 3),
				Arguments.of("""
						SAMX
						AAMX
						MMMX
						XXXX
						""", 3),
				Arguments.of("""
						XMAS
						XMAA
						XMMM
						XXXX
						""", 3)
		);
	}

	@ParameterizedTest
	@MethodSource("testData")
	void testInputPart1Extra(String testInput, long expected) {
		System.out.println(Arrays.toString(testInput.split("\n")));
		assertEquals(expected, day04.solvePart1(testInput));
	}

	@Test
	void testRealInputPart1() {
		assertEquals(2642, day04.solvePart1(ParseUtil.readInputFile(4)));
	}

	@Test
	void testInputPart2() {
		String testInput = """
				MMMSXXMASM
				MSAMXMSMSA
				AMXSXMAAMM
				MSAMASMSMX
				XMASAMXAMM
				XXAMMXXAMA
				SMSMSASXSS
				SAXAMASAAA
				MAMMMXMMMM
				MXMXAXMASX
				""";

		assertEquals(0, day04.solvePart2(testInput));
	}


	@Test
	void testRealInputPart2() {
		assertEquals(0, day04.solvePart2(ParseUtil.readInputFile(4)));
	}
}