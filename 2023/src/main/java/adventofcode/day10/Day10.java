package adventofcode.day10;

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.stream.Stream;

import adventofcode.util.AbstractDay;

public class Day10 extends AbstractDay {

	record Position(int y, int x) {

		public Position add(int y, int x) {
			return new Position(this.y + y, this.x + x);
		}

		public Position add(Position position) {
			return add(position.y, position.x);
		}

		public Position opposite() {
			return new Position(-y, -x);
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

	enum Direction {
		TOP,
		RIGHT,
		BOTTOM,
		LEFT
	}

	@Override
	public long solvePart1(String input) {
		return calcDepthLoop(parseMap(input));
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
	public long solvePart2(String input) {
		char[][] map = parseMap(input);
		char newS = getReplacementForS(map);
		List<Position> loop = getLoop(map);

		int enclosedTiles = 0;
		for (int y = 0; y < map.length; y++) {
			boolean inside = false;
			boolean isLine = false;
			boolean fromTop = false;
			for (int x = 0; x < map[0].length; x++) {
				char currentChar = map[y][x];
				if (currentChar == 'S') {
					currentChar = newS;
				}
				if (!loop.contains(new Position(y, x))) {
					currentChar = '.';
				}

				switch (currentChar) {
				case 'L', 'F' -> {
					isLine = true;
					fromTop = currentChar == 'L';
				}
				case 'J', '7' -> {
					isLine = false;
					if ((fromTop && currentChar == '7') || (!fromTop && currentChar == 'J')) {
						inside = !inside;
					}
				}
				case '|' -> inside = !inside;
				case '.' -> {
					if (!isLine && inside) {
						enclosedTiles++;
					}
				}
				}
			}
		}
		return enclosedTiles;
	}

	private char getReplacementForS(char[][] map) {
		Position startPosition = findStartPosition(map);
		List<Direction> directions = getConnectedPositions(startPosition, map).stream()
				.map(p -> p.add(-startPosition.y, -startPosition.x))
				.map(this::getDirection)
				.toList();
		return Map.of(List.of(Direction.TOP, Direction.BOTTOM), '|',
						List.of(Direction.LEFT, Direction.RIGHT), '-',
						List.of(Direction.TOP, Direction.RIGHT), 'L',
						List.of(Direction.TOP, Direction.LEFT), 'J',
						List.of(Direction.LEFT, Direction.BOTTOM), '7',
						List.of(Direction.RIGHT, Direction.BOTTOM), 'F'
				).entrySet()
				.stream()
				.filter(entry -> new HashSet<>(entry.getKey()).containsAll(directions))
				.map(Map.Entry::getValue)
				.findFirst()
				.orElseThrow();
	}

	private Direction getDirection(Position position) {
		if (position.y == -1 && position.x == 0) {
			return Direction.TOP;
		}
		if (position.y == 0 && position.x == 1) {
			return Direction.RIGHT;
		}
		if (position.y == 1 && position.x == 0) {
			return Direction.BOTTOM;
		}
		if (position.y == 0 && position.x == -1) {
			return Direction.LEFT;
		}
		throw new IllegalStateException();
	}

	private List<Position> getLoop(char[][] map) {
		Position startPosition = findStartPosition(map);
		List<Position> loop = new ArrayList<>();
		List<Position> currentPositions = new ArrayList<>(List.of(startPosition));

		while (!currentPositions.isEmpty()) {
			loop.addAll(currentPositions);

			List<Position> nextPositions = new ArrayList<>();
			for (Position currentPosition : currentPositions) {
				List<Position> connectedPositions = getConnectedPositions(currentPosition, map).stream()
						.filter(p -> !p.equals(currentPosition))
						.filter(p -> p.isInBounds(map))
						.filter(p -> p.isPipe(map))
						.toList();
				for (Position position : connectedPositions) {
					if (!loop.contains(position)) {
						nextPositions.add(position);
					}
				}
			}
			currentPositions = nextPositions;
		}
		return loop;
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
