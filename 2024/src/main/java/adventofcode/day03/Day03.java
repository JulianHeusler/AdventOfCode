package adventofcode.day03;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

import adventofcode.util.AbstractDay;

public class Day03 extends AbstractDay {

	@Override
	public long solvePart1(String input) {
		long sum = 0;
		Matcher matcher = Pattern.compile("mul\\((\\d{1,3}),(\\d{1,3})\\)").matcher(input);
		while (matcher.find()) {
			sum += getProductFromMatch(matcher);
		}
		return sum;
	}

	@Override
	public long solvePart2(String input) {
		long sum = 0;
		boolean enabled = true;
		Pattern pattern = Pattern.compile("^mul\\((\\d{1,3}),(\\d{1,3})\\)");

		for (String line : input.split("\n")) {
			int i = 0;
			while (i < line.length()) {
				String cursor = line.substring(i);
				if (cursor.startsWith("do()")) {
					enabled = true;
					i += 4;
					continue;
				}
				if (cursor.startsWith("don't()")) {
					enabled = false;
					i += 7;
					continue;
				}
				if (enabled) {
					Matcher matcher = pattern.matcher(cursor);
					if (matcher.find()) {
						sum += getProductFromMatch(matcher);
					}
				}
				i++;
			}
		}

		return sum;
	}

	private long getProductFromMatch(Matcher matcher) {
		int firstDigit = Integer.parseInt(matcher.group(1));
		int secondDigit = Integer.parseInt(matcher.group(1 + 1));
		return (long) firstDigit * secondDigit;
	}
}
