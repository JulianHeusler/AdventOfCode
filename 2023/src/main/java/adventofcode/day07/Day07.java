package adventofcode.day07;

import adventofcode.util.AbstractDay;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;

public class Day07 extends AbstractDay {

    record Hand(List<Character> cards, int bid) {

        private int isHigher(Hand other) {
            if (isFullHouse(cards) && isFullHouse(other.cards)) {
                return isHandHigherSingle(other);
            }
            if (isFullHouse(cards) && (sameChars(other.cards) < 4)) {
                return 1;
            }
            if ((sameChars(cards) < 4) && isFullHouse(other.cards)) {
                return -1;
            }

            if (sameChars(cards) > sameChars(other.cards)) {
                return 1;
            }
            if (sameChars(cards) < sameChars(other.cards)) {
                return -1;
            }
            return isHandHigherSingle(other);
        }

        private int isHandHigherSingle(Hand other) {
            for (int i = 0; i < cards.size(); i++) {
                if (Objects.equals(cards.get(i), other.cards.get(i))) {
                    continue;
                }
                return getCardValue(cards.get(i)).compareTo(getCardValue(other.cards.get(i)));
            }
            throw new IllegalStateException();
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

        private boolean isFullHouse(List<Character> cards) {
            Object[] sortedHand = cards.stream().sorted().toArray();
            return (sortedHand[0] == sortedHand[1] && sortedHand[1] == sortedHand[2] && sortedHand[3] == sortedHand[4])
                    || (sortedHand[0] == sortedHand[1] && sortedHand[2] == sortedHand[3] && sortedHand[3] == sortedHand[4]);
        }

        private int sameChars(List<Character> cards) {
            return cards.stream()
                    .collect(Collectors.groupingBy(c -> c, Collectors.counting()))
                    .values().stream()
                    .mapToInt(Math::toIntExact)
                    .max().orElse(0);
        }
    }

    @Override
    public int solvePart1(String input) {
        List<Hand> hands = parseHands(input);

        List<Hand> sorted = hands.stream().sorted(Hand::isHigher).toList();
        int result = 0;
        for (int rank = 1; rank <= sorted.size(); rank++) {
            result += rank * sorted.get(rank - 1).bid;
        }

        return result;
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
