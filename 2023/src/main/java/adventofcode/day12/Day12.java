package adventofcode.day12;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

public class Day12 extends AbstractDay {
	@Override
	public long solvePart1(String input) {
		long result = 0;

		for (String line : input.split("\n")) {
			String[] split = line.split(" ");
			List<Integer> numbers = ParseUtil.parseIntegerNumbers(split[1]);

			List<String> permutations = generatePermutations(split[0], 0);
			List<String> validOnes = permutations.stream()
					.filter(s -> isValidLine(s, numbers))
					.toList();
			result += validOnes.size();
		}
		return result;
	}

	private List<String> getPermuations(String line) {
		List<String> perm = new ArrayList<>();
		long questionCount = line.chars().filter(c -> c == '?').count();

		for (int q = 0; q < questionCount; q++) {
			int count = (int) line.chars().count();
			for (int i = 0; i < count; i++) {
				if (line.charAt(i) == '?') {
					if (i == 0) {
						perm.add('.' + line.substring(i + 1, count));
						perm.add('#' + line.substring(i + 1, count));
					} else {
						perm.add(line.substring(i - 1, i) + '.' + line.substring(i + 1, count));
						perm.add(line.substring(i - 1, i) + '#' + line.substring(i + 1, count));
					}
				}
			}
		}
		return perm;
	}

	private List<String> generatePermutations(String current, int index) {
		List<String> permutations = new ArrayList<>();
		if (index == current.length()) {
			permutations.add(current);
		} else {
			if (current.charAt(index) == '?') {
				permutations.addAll(generatePermutations(setCharAt('.', index, current), index + 1));
				permutations.addAll(generatePermutations(setCharAt('#', index, current), index + 1));
			} else {
				permutations.addAll(generatePermutations(current, index + 1));
			}
		}
		return permutations;
	}

	private String setCharAt(char c, int index, String s) {
		return s.substring(0, index) + c + s.substring(index + 1);
	}

	private boolean isValidLine(String line, List<Integer> numbers) {
		List<String> split = Arrays.stream(line.split("\\."))
				.filter(s -> !s.isEmpty())
				.toList();
		if (split.size() != numbers.size()) {
			return false;
		}
		for (int i = 0; i < split.size(); i++) {
			if (split.get(i).length() != numbers.get(i)) {
				return false;
			}
		}
		return true;
	}

	@Override
	public long solvePart2(String input) {
		return 0;
	}
}
