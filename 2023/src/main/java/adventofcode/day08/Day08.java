package adventofcode.day08;

import adventofcode.util.AbstractDay;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day08 extends AbstractDay {

    record Element(String name, String left, String right) {}

    @Override
    public int solvePart1(String input) {
        String directions = input.substring(0, input.indexOf("\n"));
        Map<String, Element> elements = parseElements(input);

        int steps = 0;
        Element current = elements.get("AAA");
        while (!current.name.equals("ZZZ")) {
            int dir = directions.chars().findFirst().orElseThrow();
            directions = shiftStringLeft(directions);
            steps++;

            if (dir == 'R') {
                current = elements.get(current.right());
            } else {
                current = elements.get(current.left());
            }
        }

        return steps;
    }

    @Override
    public int solvePart2(String input) {
        String directions = input.substring(0, input.indexOf("\n"));
        Map<String, Element> elements = parseElements(input);

        int steps = 0;
        List<Element> currentElements = getStartingElements(elements);
        while (!currentElements.stream().map(Element::name).allMatch(s -> s.endsWith("Z"))) {
            int direction = directions.chars().findFirst().orElseThrow();
            String dir = Character.toString(direction);
            directions = shiftStringLeft(directions);
            steps++;

            currentElements = simulateMove(elements, direction, currentElements);
        }

        return steps;
    }

    private List<Element> getStartingElements(Map<String, Element> elements) {
        return elements.entrySet()
                .stream()
                .filter(entry -> entry.getKey().endsWith("A"))
                .map(Map.Entry::getValue)
                .toList();
    }

    private List<Element> simulateMove(final Map<String, Element> elementMap, int direction, List<Element> currentElements) {
        List<Element> resultingElements = new ArrayList<>();
        for (Element current : currentElements) {
            if (direction == 'R') {
                resultingElements.add(elementMap.get(current.right()));
            } else {
                resultingElements.add(elementMap.get(current.left()));
            }
        }
        return resultingElements;
    }

    private String shiftStringLeft(String input) {
        return input.substring(1) + input.charAt(0);
    }

    private Map<String, Element> parseElements(String input) {
        Map<String, Element> elements = new HashMap<>();
        Pattern pattern = Pattern.compile("^(\\w\\w\\w) = \\((\\w\\w\\w), (\\w\\w\\w)\\)$");

        String[] split = input.split("\n");
        for (int i = 2; i < split.length; i++) {
            Matcher matcher = pattern.matcher(split[i]);
            if (matcher.matches()) {
                Element element = new Element(matcher.group(1), matcher.group(2), matcher.group(3));
                elements.put(element.name(), element);
            }
        }
        return elements;
    }
}
