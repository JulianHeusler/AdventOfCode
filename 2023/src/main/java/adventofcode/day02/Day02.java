package adventofcode.day02;

import adventofcode.util.AbstractDay;

import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day02 extends AbstractDay {

    record Game(int id, List<Round> rounds) {
        boolean isPossible() {
            return rounds.stream().allMatch(round -> round.red <= 12 && round.green <= 13 && round.blue <= 14);
        }
    }

    record Round(int red, int green, int blue) {
    }

    @Override
    public int solvePart1(String input) {
        return parseGames(input).stream().filter(Game::isPossible).mapToInt(Game::id).sum();
    }

    @Override
    public int solvePart2(String input) {
        return parseGames(input).stream().mapToInt(this::calculateGameScore).sum();
    }

    private int calculateGameScore(Game game) {
        return game.rounds.stream().mapToInt(Round::red).max().orElse(1)
                * game.rounds.stream().mapToInt(Round::green).max().orElse(1)
                * game.rounds.stream().mapToInt(Round::blue).max().orElse(1);
    }

    private List<Game> parseGames(String input) {
        List<Game> gameList = new ArrayList<>();
        int id = 1;
        for (String line : input.split("\n")) {
            gameList.add(new Game(id++, parseRounds(line)));
        }
        return gameList;
    }

    private List<Round> parseRounds(String game) {
        List<Round> rounds = new ArrayList<>();
        for (String round : game.split(";")) {
            rounds.add(new Round(getCount("red", round), getCount("green", round), getCount("blue", round)));
        }
        return rounds;
    }

    private int getCount(String color, String round) {
        Matcher matcher = Pattern.compile("(\\d+) " + color).matcher(round);
        if (matcher.find()) {
            return Integer.parseInt(matcher.group(1));
        }
        return 0;
    }
}
