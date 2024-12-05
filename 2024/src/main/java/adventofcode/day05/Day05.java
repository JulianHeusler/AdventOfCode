package adventofcode.day05;

import java.util.ArrayList;
import java.util.LinkedList;
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
		String[] split = input.split("\n\n");
		List<Rule> rules = parseRules(split[0]);
		List<Order> orders = parseOrders(split[1]);

		return orders.stream()
				.filter(order -> !isValidOrder(order, rules))
				.map(order -> sortOrder(order, rules).getMiddlePageNumber())
				.mapToInt(Integer::intValue)
				.sum();
	}

	private Order sortOrder(Order order, List<Rule> rules) {
		List<Integer> sortedPages = new ArrayList<>();
		sortedPages.addFirst(order.pages.getFirst());
		for (int i = 1; i < order.pages().size(); i++) {
			Integer nextPage = order.pages().get(i);
			sortedPages = validRuleSort(rules, sortedPages, nextPage);
		}
		return new Order(sortedPages);
	}

	private List<Integer> validRuleSort(List<Rule> rules, List<Integer> pages, int page) {
		for (int i = 0; i < pages.size() + 1; i++) {
			List<Integer> candidate = insertNumberAt(pages, page, i);
			if (isValidOrder(new Order(candidate), rules)) {
				return candidate;
			}
		}
		throw new IllegalStateException();
	}

	private List<Integer> insertNumberAt(List<Integer> integers, int number, int index) {
		LinkedList<Integer> linkedList = new LinkedList<>(integers);
		linkedList.add(index, number);
		return linkedList;
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
