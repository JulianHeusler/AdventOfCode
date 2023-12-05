package adventofcode.day05;

import adventofcode.util.AbstractDay;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day05 extends AbstractDay {


    record Map(String title, List<Range> ranges) {
    }

    record Range(Long destinationRangeStart, Long sourceRangeStart, Long length) {
        boolean contains(Long l) {
            return sourceRangeStart <= l && l <= sourceRangeStart + length;
        }

//        boolean overlaps(Range other) {
//            return IntStream.range(sourceRangeStart, sourceRangeStart + length).anyMatch(other::contains);
//        }
    }

    @Override
    public int solvePart1(String input) {
        return (int) parseInput(input).stream().mapToLong(Long::longValue).min().orElseThrow();
    }

    @Override
    public int solvePart2(String input) {
        return 0;
    }


    private List<Long> parseInput(String input) {
        String[] split = input.split("\n");
        List<Long> seeds = parseNumbers(split[0]);

        List<Map> maps = getMaps(Arrays.copyOfRange(split, 2, split.length));

        List<Long> result = new ArrayList<>();

        for (Long seed : seeds) {
            System.out.println("seed:" + seed);
            result.add(getLocation(seed, maps));
        }

        return result;
    }

    private Long getLocation(Long seed, List<Map> maps) {
        Long next = seed;
        for (Map map : maps) {
            Long current = next;
            next = map.ranges.stream().filter(range -> range.contains(current))
                    .map(range -> resolve(range, current)).findFirst().orElse(current);
            System.out.println(next);
        }
        return next;
    }

    private Long resolve(Range range, Long current) {
        Long offset = range.destinationRangeStart - range.sourceRangeStart;
        return current + offset;
    }


    private List<Map> getMaps(String[] lines) {
        List<Map> maps = new ArrayList<>();

        List<Range> ranges = new ArrayList<>();
        String lastTitle = "";
        for (String line : lines) {
            if (line.isEmpty()) {
                maps.add(new Map(lastTitle, ranges));
                ranges = new ArrayList<>();
            } else if (line.contains("map")) {
                lastTitle = line;
            } else {
                List<Long> lineNumbers = parseNumbers(line);
                ranges.add(new Range(lineNumbers.get(0), lineNumbers.get(1), lineNumbers.get(2)));
            }
        }

        maps.add(new Map(lastTitle, ranges));
        return maps;
    }

    private List<Long> parseNumbers(String numberLine) {
        List<Long> numbers = new ArrayList<>();
        Matcher matcher = Pattern.compile("(\\d+)").matcher(numberLine);
        while (matcher.find()) {
            numbers.add(Long.parseLong(matcher.group()));
        }
        return numbers;
    }
}
