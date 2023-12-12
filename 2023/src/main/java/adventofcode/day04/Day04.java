package adventofcode.day04;

import java.util.ArrayList;
import java.util.List;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

public class Day04 extends AbstractDay {
	record Card(int id, List<Integer> winningNumbers, List<Integer> yourNumbers) {
	}

	@Override
	public long solvePart1(String input) {
		return parseCards(input)
				.stream()
				.mapToInt(this::countMatchingNumbers)
				.filter(value -> value != 0)
				.map(this::calculateScore)
				.sum();
	}

	private int countMatchingNumbers(Card card) {
		return card.yourNumbers.stream()
				.filter(card.winningNumbers::contains)
				.toList().size();
	}

	private int calculateScore(int count) {
		return (int) Math.pow(2, count - 1.0);
	}

	@Override
	public long solvePart2(String input) {
		List<Card> originalCards = parseCards(input);
		return solveScratchCards((originalCards), 0);
	}

	private int solveScratchCards(final List<Card> cardsToScratch, final int count) {
		if (cardsToScratch.isEmpty()) {
			return count;
		}

		Card currentCard = cardsToScratch.getFirst();
		int currentCardAmount = (int) cardsToScratch.stream()
				.filter(card -> card.id() == currentCard.id())
				.count();
		cardsToScratch.removeIf(card -> card.id() == currentCard.id());

		for (int i = 1; i <= countMatchingNumbers(currentCard); i++) {
			int wonCardId = currentCard.id() + i;
			Card wonCard = cardsToScratch.stream()
					.filter(card -> card.id() == wonCardId)
					.findFirst()
					.orElseThrow();
			for (int j = 0; j < currentCardAmount; j++) {
				cardsToScratch.add(wonCard);
			}
		}
		return solveScratchCards(cardsToScratch, count + currentCardAmount);
	}


	private List<Card> parseCards(String input) {
		List<Card> cards = new ArrayList<>();
		int id = 1;
		for (String line : input.split("\n")) {
			String substringCard = line.substring(line.indexOf(":") + 2);
			String[] splitNumberLines = substringCard.split(" \\| ");
			cards.add(new Card(id++, ParseUtil.parseIntegerNumbers(splitNumberLines[0]), ParseUtil.parseIntegerNumbers(splitNumberLines[1])));
		}
		return cards;
	}
}
