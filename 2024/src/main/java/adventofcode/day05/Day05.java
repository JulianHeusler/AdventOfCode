package adventofcode.day05;

import java.util.ArrayList;
import java.util.List;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

public class Day05 extends AbstractDay {

	@Override
	public long solvePart1(String input) {
		long sum = 0;
		String[] split = input.split("\n\n");
		List<Rule> rules = parseRules(split[0]);
		List<Order> orders = parseOrders(split[1]);

		for (Order order : orders) {
			if (isValidOrder(order, rules)) {
				sum += order.getMiddlePageNumber();
			}
		}
		return sum;
	}

	private boolean isValidOrder(Order order, List<Rule> rules) {
		List<Rule> relevantRules = rules.stream().filter(rule -> order.pages().contains(rule.pageNumber) &&
				order.pages().contains(rule.beforePageNumber)).toList();

		for (Rule rule : relevantRules) {
			if (!rule.satisfiesRules(order.pages)) {
				return false;
			}
		}

		return true;
	}

	@Override
	public long solvePart2(String input) {
		return 0;
	}

	private List<Order> parseOrders(String splitInput) {
		List<Order> orders = new ArrayList<>();
		for (String line : splitInput.split("\n")) {
			orders.add(new Order(ParseUtil.parseIntegerNumbers(line)));
		}
		return orders;
	}

	private record Order(List<Integer> pages) {
		int getMiddlePageNumber() {
			int middle = pages.size() / 2;
			return pages.get(middle);
		}
	}

	private List<Rule> parseRules(String splitInput) {
		List<Rule> rules = new ArrayList<>();
		for (String line : splitInput.split("\n")) {
			String[] split = line.split("\\|");
			rules.add(new Rule(Integer.parseInt(split[0]), Integer.parseInt(split[1])));
		}
		return rules;
	}

	private record Rule(int pageNumber, int beforePageNumber) {
		boolean satisfiesRules(List<Integer> pages) {
			return pages.indexOf(pageNumber) <= pages.indexOf(beforePageNumber);
		}
	}
}
