package adventofcode.day07;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import adventofcode.util.AbstractDay;

public class Day07 extends AbstractDay {
	private static boolean isPartTwo = false;

	enum Type {
		FIVE_OF_KIND,
		FOUR_OF_KIND,
		FULL_HOUSE,
		THREE_OF_KIND,
		TWO_PAIR,
		ONE_PAIR,
		HIGH_CARD;


		int getRank() {
			return Type.values().length - this.ordinal();
		}

	}

	record Hand(List<Character> cards, int bid) {

		public int compareTo(Hand other) {
			int rankSelf = resolveType(cards).getRank();
			int rankOther = resolveType(other.cards).getRank();
			if (rankSelf == rankOther) {
				return compareCardByCard(other);
			}
			return Integer.compare(rankSelf, rankOther);
		}

		@Override
		public String toString() {
			return String.format("Cards: %s, Bid: %d, Type: %s", cards, bid, resolveType(cards));
		}

		private Type resolveType(List<Character> cards) {
			int mostCommonCardCount = mostCommonCharCount(cards);
			if (mostCommonCardCount == 5) {
				return Type.FIVE_OF_KIND;
			}
			if (mostCommonCardCount == 4) {
				return Type.FOUR_OF_KIND;
			}
			if (isFullHouse(cards)) {
				return Type.FULL_HOUSE;
			}
			if (mostCommonCardCount == 3) {
				return Type.THREE_OF_KIND;
			}
			if (isTwoPair(cards)) {
				return Type.TWO_PAIR;
			}
			if (mostCommonCardCount == 2) {
				return Type.ONE_PAIR;
			}
			return Type.HIGH_CARD;
		}

		private int compareCardByCard(Hand other) {
			for (int i = 0; i < cards.size(); i++) {
				int compareSingleCards = getCardValue(this.cards().get(i)).compareTo(getCardValue(other.cards.get(i)));
				if (compareSingleCards == 0) {
					continue;
				}
				return compareSingleCards;
			}
			throw new IllegalStateException();
		}

		private Integer getCardValue(Character c) {
			return switch (c) {
				case 'A' -> 14;
				case 'K' -> 13;
				case 'Q' -> 12;
				case 'J' -> isPartTwo ? 1 : 11;
				case 'T' -> 10;
				default -> Integer.parseInt(String.valueOf(c));
			};
		}

		private boolean isFullHouse(List<Character> cards) {
			if (mostCommonCharCount(cards) > 4) {
				return false;
			}

			return cards.stream()
					.filter(c -> isPartTwo ? c != 'J' : true)
					.collect(Collectors.groupingBy(c -> c, Collectors.counting()))
					.size() == 2;
		}

		private boolean isTwoPair(List<Character> cards) {
			return cards.stream().collect(Collectors.groupingBy(c -> c, Collectors.counting())).size() == 3;
		}

		private int mostCommonCharCount(List<Character> cards) {
			if (isPartTwo) {
				int jokerCount = (int) cards.stream().filter(character -> character == 'J').count();
				return jokerCount + cards.stream()
						.filter(c -> c != 'J')
						.collect(Collectors.groupingBy(c -> c, Collectors.counting()))
						.values().stream()
						.mapToInt(Math::toIntExact)
						.max().orElse(-1);
			}

			return cards.stream()
					.collect(Collectors.groupingBy(c -> c, Collectors.counting()))
					.values().stream()
					.mapToInt(Math::toIntExact)
					.max().orElse(-1);
		}
	}

	@Override
	public long solvePart1(String input) {
		isPartTwo = false;
		return calculateTotalWinnings(input);
	}

	@Override
	public long solvePart2(String input) {
		isPartTwo = true;
		return calculateTotalWinnings(input);
	}

	private int calculateTotalWinnings(String input) {
		List<Hand> hands = parseHands(input);
		List<Hand> sortedHands = hands.stream().sorted(Hand::compareTo).toList();
		sortedHands.forEach(System.out::println);

		int result = 0;
		for (int rank = 1; rank <= sortedHands.size(); rank++) {
			result += rank * sortedHands.get(rank - 1).bid;
		}
		return result;
	}

	private List<Hand> parseHands(String input) {
		List<Hand> hands = new ArrayList<>();
		for (String line : input.split("\n")) {
			List<Character> cards = new ArrayList<>();
			String[] split = line.split(" ");
			split[0].chars().forEach(card -> cards.add((char) card));
			hands.add(new Hand(cards, Integer.parseInt(split[1])));
		}
		return hands;
	}
}
