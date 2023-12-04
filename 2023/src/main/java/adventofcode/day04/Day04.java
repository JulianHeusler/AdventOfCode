package adventofcode.day04;

import adventofcode.util.AbstractDay;

import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day04 extends AbstractDay {
    record Card(int id, List<Integer> winningNumbers, List<Integer> yourNumbers) {}

    @Override
    public int solvePart1(String input) {
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
    public int solvePart2(String input) {
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
            cards.add(new Card(id++, parseNumbers(splitNumberLines[0]), parseNumbers(splitNumberLines[1])));
        }
        return cards;
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
