package adventofcode.day10;

import adventofcode.util.AbstractDay;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.stream.Stream;

public class Day10 extends AbstractDay {

    record Position(int y, int x) {
        public Position add(int y, int x) {
            return new Position(this.y + y, this.x + x);
        }

        public boolean isPipe(char[][] map) {
            return getTile(map) != '.';
        }

        public char getTile(char[][] map) {
            return map[y][x];
        }

        public boolean isInBounds(char[][] map) {
            return 0 <= y && y < map.length && 0 <= x && x < map[0].length;
        }
    }

    @Override
    public int solvePart1(String input) {
        char[][] map = parseMap(input);

        Position startPosition = findStartPosition(map);

        // return calcDepth2(map, Collections.emptyList(), new ArrayList<>(List.of(startPosition)), 0);
        return calcDepthLoop(map);
    }


    private int calcDepthLoop(char[][] map) {
        Position startPosition = findStartPosition(map);
        List<Position> lastPositions = new ArrayList<>();
        List<Position> currentPositions = new ArrayList<>(List.of(startPosition));
        int depth = 0;

        while (!currentPositions.isEmpty()) {
            List<Position> nextPositions = new ArrayList<>();
            for (Position currentPosition : currentPositions) {
                List<Position> connectedPositions = getConnectedPositions(currentPosition, map).stream()
                        .filter(p -> !p.equals(currentPosition))
                        .filter(p -> p.isInBounds(map))
                        .filter(p -> p.isPipe(map))
                        .toList();
                for (Position position : connectedPositions) {
                    if (!lastPositions.contains(position)) {
                        nextPositions.add(position);
                    }
                }
            }
            lastPositions = currentPositions;
            currentPositions = nextPositions;
            depth++;
        }
        return depth - 1;
    }


    private int calcDepthLast(char[][] map, List<Position> lastPositions, List<Position> currentPositions, int depth) {
        if (currentPositions.isEmpty()) {
            return depth - 1;
        }
        List<Position> nextPositions = new ArrayList<>();
        for (Position currentPosition : currentPositions) {
            List<Position> connectedPositions = getConnectedPositions(currentPosition, map).stream()
                    .filter(p -> !p.equals(currentPosition))
                    .filter(p -> p.isInBounds(map))
                    .filter(p -> p.isPipe(map))
                    .toList();
            for (Position position : connectedPositions) {
                if (!lastPositions.contains(position)) {
                    nextPositions.add(position);
                }
            }
        }
        return calcDepthLast(map, currentPositions, nextPositions, depth + 1);
    }

    private int calcDepthPipe(char[][] map, List<Position> pipe, List<Position> currentPositions, int depth) {
        if (currentPositions.isEmpty()) {
            return depth - 1;
        }
        List<Position> nextPositions = new ArrayList<>();
        for (Position currentPosition : currentPositions) {
            List<Position> connectedPositions = getConnectedPositions(currentPosition, map).stream()
                    .filter(p -> !p.equals(currentPosition))
                    .filter(p -> p.isInBounds(map))
                    .filter(p -> p.isPipe(map))
                    .toList();
            for (Position position : connectedPositions) {
                if (!pipe.contains(position)) {
                    pipe.add(position);
                    nextPositions.add(position);
                }
            }
        }
        return calcDepthPipe(map, pipe, nextPositions, depth + 1);
    }


    private List<Position> getConnectedPositions(Position current, char[][] map) {
        return switch (current.getTile(map)) {
            case '|' -> List.of(current.add(-1, 0), current.add(1, 0));
            case '-' -> List.of(current.add(0, -1), current.add(0, 1));
            case 'L' -> List.of(current.add(-1, 0), current.add(0, 1));
            case 'J' -> List.of(current.add(-1, 0), current.add(0, -1));
            case '7' -> List.of(current.add(0, -1), current.add(1, 0));
            case 'F' -> List.of(current.add(0, 1), current.add(1, 0));
            case 'S' -> Stream.of(current.add(-1, 0),
                            current.add(0, 1),
                            current.add(1, 0),
                            current.add(0, -1))
                    .filter(position -> position.isInBounds(map))
                    .filter(p -> getConnectedPositions(p, map).contains(current)).toList();
            default -> Collections.emptyList(); // '.'
        };
    }

    private boolean isInBounds(String[] matrix, int y, int x) {
        return 0 <= y && y < matrix.length && 0 <= x && x < matrix[0].length();
    }

    private Position findStartPosition(char[][] map) {
        for (int y = 0; y < map.length; y++) {
            for (int x = 0; x < map[y].length; x++) {
                if (map[y][x] == 'S') {
                    return new Position(y, x);
                }
            }
        }
        throw new IllegalStateException();
    }

    @Override
    public int solvePart2(String input) {
        return 0;
    }

    private char[][] parseMap(String input) {
        String[] lines = input.split("\n");
        char[][] map = new char[lines.length][lines[0].length()];
        for (int y = 0; y < lines.length; y++) {
            String currentLine = lines[y];
            for (int x = 0; x < currentLine.length(); x++) {
                map[y][x] = lines[y].charAt(x);
            }
        }
        return map;
    }
}
