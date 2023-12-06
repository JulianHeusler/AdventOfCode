package adventofcode.day06;

import adventofcode.util.AbstractDay;
import adventofcode.util.ParseUtil;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public class Day06 extends AbstractDay {
    record Race(int time, int distance) {}

    record Tuple(double a, double b) {}

    @Override
    public int solvePart1(String input) {
        return parseRaces(input)
                .stream()
                .map(race -> getWinningChargeTimes(race).size())
                .reduce(1, (a, b) -> a * b);
    }

    private List<Integer> getWinningChargeTimes(Race race) {
        List<Integer> winningChargeTimes = new ArrayList<>();
        for (int i = 0; i < race.time; i++) {
            if (calculateDistance(race.time, i) > race.distance) {
                winningChargeTimes.add(i);
            }
        }
        return winningChargeTimes;
    }

    private int calculateDistance(int raceTime, int chargeTime) {
        int timeLeft = raceTime - chargeTime;
        return timeLeft * chargeTime;
    }

    @Override
    public int solvePart2(String input) {
        String[] split = input.split("\n");
        long time = parseWholeNumber(split[0]);
        long distance = parseWholeNumber(split[1]);
        Tuple tuple = quadraticFormula(-1, time, -distance);
        return (int) (Math.ceil(tuple.b) - Math.ceil(tuple.a));
    }

    private Tuple quadraticFormula(double a, double b, double c) {
        double discriminant = b * b - 4 * a * c;

        if (discriminant > 0) {
            double root1 = (-b + Math.sqrt(discriminant)) / (2 * a);
            double root2 = (-b - Math.sqrt(discriminant)) / (2 * a);
            return new Tuple(root1, root2);
        } else if (discriminant == 0) {
            double root = -b / (2 * a);
            return new Tuple(root, root);
        } else {
            throw new IllegalStateException("imaginary");
        }
    }

    private Long parseWholeNumber(String line) {
        return Long.parseLong(ParseUtil.parseIntegerNumbers(line)
                .stream()
                .map(Object::toString)
                .collect(Collectors.joining()));
    }

    private List<Race> parseRaces(String input) {
        List<Race> races = new ArrayList<>();
        String[] split = input.split("\n");
        List<Integer> times = ParseUtil.parseIntegerNumbers(split[0]);
        List<Integer> distances = ParseUtil.parseIntegerNumbers(split[1]);
        for (int i = 0; i < times.size(); i++) {
            races.add(new Race(times.get(i), distances.get(i)));
        }
        return races;
    }
}
