package adventofcode.day07;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import adventofcode.util.AbstractDay;

public class Day07 extends AbstractDay {

	enum Type {
		FiveOfKind,
		FourOfKind,
		FullHouse,
		ThreeOfKind,
		TwoPair,
		OnePair,
		HighCard;

		int getRank() {
			return Type.values().length - this.ordinal();
		}
	}

	record Hand(List<Character> cards, int bid) {

		public int compareTo(Hand other) {
			int compareTo = isHigher(other);
			return compareTo;
		}

		private int isHigher(Hand other) {
			Type typeSelf = resolveType(cards);
			int rankSelf = typeSelf.getRank();
			Type typeOther = resolveType(other.cards);
			int rankOther = typeOther.getRank();

			if (rankSelf == rankOther) {
//				if (getHighestValue(cards) > getHighestValue(other.cards)) {
//					return 1;
//				}
//				if (getHighestValue(cards) < getHighestValue(other.cards)) {
//					return -1;
//				}
				return isHandHigherSingle(other);
			}
			return Integer.compare(rankSelf, rankOther);
		}

		@Override
		public String toString() {
			return String.format("Cards: %s, Bid: %d, Type: %s", cards, bid, resolveType(cards));
		}

		private int extracted(Hand other) {
			if (isFullHouse(cards) && isFullHouse(other.cards)) {
				int compareTo = getHighestValue(cards).compareTo(getHighestValue(other.cards));
				if (compareTo == 0) {
					return isHandHigherSingle(other);
				}
				return compareTo;
			}
			if (isFullHouse(cards) && (sameChars(other.cards) < 4)) {
				return 1;
			}
			if ((sameChars(cards) < 4) && isFullHouse(other.cards)) {
				return -1;
			}

			if (sameChars(cards) == sameChars(other.cards) && sameChars(cards) > 1) {
				int compareTo = getHighestValue(cards).compareTo(getHighestValue(other.cards));
				if (compareTo == 0) {
					return isHandHigherSingle(other);
				}
				return compareTo;
			}

			if (sameChars(cards) > sameChars(other.cards)) {
				return 1;
			}
			if (sameChars(cards) < sameChars(other.cards)) {
				return -1;
			}
			return isHandHigherSingle(other);
		}


		private Integer getHighestValue(List<Character> cards) {
			List<Character> maxCombo = cards.stream()
					.collect(Collectors.groupingBy(c -> c)).values().stream().max((o1, o2) ->
							{
								if (o1.size() == o2.size()) {
									return compareSingleCards(o1.getFirst(), o2.getFirst());
								}
								return Integer.compare(o1.size(), o2.size());
							}
					).orElseThrow();
			return getCardValue(maxCombo.getFirst());
		}

		private int isHandHigherSingle(Hand other) {
			for (int i = 0; i < cards.size(); i++) {
				Integer cardValue = getCardValue(cards.get(i));
				Integer cardValueOther = getCardValue(other.cards.get(i));
				if (cardValue.equals(cardValueOther)) {
					continue;
				}
				return cardValue.compareTo(cardValueOther);
			}
			throw new IllegalStateException();
		}

		private int compareSingleCards(char a, char b) {
			return getCardValue(a).compareTo(getCardValue(b));
		}

		private Integer getCardValue(Character c) {
			return switch (c) {
				case 'A' -> 14;
				case 'K' -> 13;
				case 'Q' -> 12;
				case 'J' -> 11;
				case 'T' -> 10;
				default -> Integer.parseInt(String.valueOf(c));
			};
		}


	}

	private static Type resolveType(List<Character> cards) {
		int sameChars = sameChars(cards);
		if (sameChars == 5) {
			return Type.FiveOfKind;
		}
		if (sameChars == 4) {
			return Type.FourOfKind;
		}
		if (isFullHouse2(cards)) {
			return Type.FullHouse;
		}
		if (sameChars == 3) {
			return Type.ThreeOfKind;
		}
		if (isTwoPair(cards)) {
			return Type.TwoPair;
		}
		if (sameChars == 2) {
			return Type.OnePair;
		}
		return Type.HighCard;
	}

	private static boolean isTwoPair(List<Character> cards) {
		return cards.stream().collect(Collectors.groupingBy(c -> c, Collectors.counting())).size() == 3;
	}

	private static boolean isFullHouse(List<Character> cards) {
		Object[] sortedHand = cards.stream().sorted().toArray();
		return (sortedHand[0] == sortedHand[1] && sortedHand[1] == sortedHand[2] && sortedHand[3] == sortedHand[4])
				|| (sortedHand[0] == sortedHand[1] && sortedHand[2] == sortedHand[3] && sortedHand[3] == sortedHand[4]);
	}

	private static boolean isFullHouse2(List<Character> cards) {
		return cards.stream().collect(Collectors.groupingBy(c -> c, Collectors.counting())).size() == 2;
	}

	private static int sameChars(List<Character> cards) {
		return cards.stream()
				.collect(Collectors.groupingBy(c -> c, Collectors.counting()))
				.values().stream()
				.mapToInt(Math::toIntExact)
				.max().orElse(0);
	}

	@Override
	public int solvePart1(String input) {
		List<Hand> hands = parseHands(input);

		List<Hand> sorted = hands.stream().sorted(Hand::compareTo).toList();
		sorted.forEach(hand -> System.out.println(hand));
		long result = 0;
		for (int rank = 1; rank <= sorted.size(); rank++) {
			result += (long) rank * sorted.get(rank - 1).bid;
		}

		System.out.println(result);

		return (int) result;
	}

	@Override
	public int solvePart2(String input) {
		return 0;
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
