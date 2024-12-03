package adventofcode.day02;

import java.util.ArrayList;
import java.util.List;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

public class Day02 extends AbstractDay {
	@Override
	public long solvePart1(String input) {
		return parseInput(input).stream().filter(this::isSafe).count();
	}

	@Override
	public long solvePart2(String input) {
		return 0;
	}


	private boolean isSafe(Report report) {
		List<Integer> deltas = new ArrayList<>();
		for (int i = 0; i < report.levels.size() - 1; i++) {
			deltas.add(report.levels.get(i) - report.levels.get(i + 1));
		}

		if (!allPositive(deltas) && !allNegative(deltas)) {
			return false;
		}
		if (deltas.stream().map(Math::abs).anyMatch(integer -> integer < 1 || 3 < integer)) {
			return false;
		}
		return true;
	}

	private static boolean allPositive(List<Integer> deltas) {
		return deltas.stream().allMatch(integer -> integer > 0);
	}

	private static boolean allNegative(List<Integer> deltas) {
		return deltas.stream().allMatch(integer -> integer < 0);
	}


	private List<Report> parseInput(String input) {
		List<Report> reports = new ArrayList<>();
		for (String line : input.split("\n")) {
			reports.add(new Report(ParseUtil.parseIntegerNumbers(line)));
		}
		return reports;
	}

	private record Report(List<Integer> levels) {
	}
}
