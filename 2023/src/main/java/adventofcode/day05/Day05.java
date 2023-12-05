package adventofcode.day05;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Optional;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import adventofcode.util.AbstractDay;

public class Day05 extends AbstractDay {


	record Map(String title, List<Range> ranges) {
	}

	record Range(int destinationRangeStart, int sourceRangeStart, int length) {
	}

	@Override
	public int solvePart1(String input) {
		parseInput(input);
		return 0;
	}

	@Override
	public int solvePart2(String input) {
		return 0;
	}


	private List<Map> parseInput(String input) {
		String[] split = input.split("\n");
		List<Integer> seeds = parseNumbers(split[0]);

		List<Map> maps = getMaps(Arrays.copyOfRange(split, 2, split.length));

		for (Integer seed : seeds) {
			for (Map map : maps) {
				Optional<Range> first = map.ranges.stream().filter(range -> isInDest(seed, range)).findFirst();

			}
		}

		return maps;
	}


	private boolean isInDest(int seed, Range range) {
		return range.destinationRangeStart <= seed && seed <= range.destinationRangeStart + range.length;
	}

	private List<Map> getMaps(String[] lines) {
		List<Map> maps = new ArrayList<>();

		List<Range> ranges = new ArrayList<>();
		String lastTitle = "";
		for (String line : lines) {
			if (line.equals("")) {
				maps.add(new Map(lastTitle, ranges));
				ranges = new ArrayList<>();
			} else if (line.contains("map")) {
				lastTitle = line;
			} else {
				List<Integer> lineNumbers = parseNumbers(line);
				ranges.add(new Range(lineNumbers.get(0), lineNumbers.get(1), lineNumbers.get(2)));
			}
		}

		maps.add(new Map(lastTitle, ranges));

		return maps;
	}

	private List<Integer> parseNumbers(String numberLine) {
		List<Integer> numbers = new ArrayList<>();
		Matcher matcher = Pattern.compile("(\\d+)").matcher(numberLine);
		while (matcher.find()) {
			numbers.add(Integer.parseInt(matcher.group()));
		}
		return numbers;
	}
}
