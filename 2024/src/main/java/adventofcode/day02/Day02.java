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
		return parseInput(input).stream().filter(report -> isSafe(report.levels, false)).count();
	}

	private boolean isSafe(List<Integer> levels, boolean skippedOne) {
		List<Integer> deltas = new ArrayList<>();
		for (int i = 0; i < levels.size() - 1; i++) {
			int delta = levels.get(i) - levels.get(i + 1);
			deltas.add(delta);
			if (isInvalidDelta(delta) || haveDifferentSigns(deltas)) {
				if (skippedOne) {
					return false;
				}
				return isSafe(removeLevel(levels, i - 1), true)
						|| isSafe(removeLevel(levels, i), true)
						|| isSafe(removeLevel(levels, i + 1), true);
			}
		}
		return true;
	}


	private List<Integer> removeLevel(List<Integer> list, int index) {
		ArrayList<Integer> copy = new ArrayList<>();
		copy.addAll(list.subList(0, Math.max(0, index)));
		copy.addAll(list.subList(Math.min(index + 1, list.size()), list.size()));
		return copy;
	}

	private boolean isSafe(Report report) {
		List<Integer> deltas = new ArrayList<>();
		for (int i = 0; i < report.levels.size() - 1; i++) {
			deltas.add(report.levels.get(i) - report.levels.get(i + 1));
		}

		if (haveDifferentSigns(deltas)) {
			return false;
		}
		return deltas.stream().noneMatch(Day02::isInvalidDelta);
	}

	private static boolean isInvalidDelta(Integer integer) {
		int abs = Math.abs(integer);
		return 1 > abs || abs > 3;
	}

	private static boolean haveDifferentSigns(List<Integer> deltas) {
		return !allPositive(deltas) && !allNegative(deltas);
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
