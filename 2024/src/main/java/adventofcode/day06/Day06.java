package adventofcode.day06;

import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

import adventofcode.util.AbstractDay;

public class Day06 extends AbstractDay {

	private record Vector(Position position, Direction direction) {
		@Override
		public String toString() {
			return "Vector pos: %s dir: %s".formatted(position, direction);
		}
	}

	private enum Direction {
		UP,
		RIGHT,
		DOWN,
		LEFT
	}

	private record Position(int x, int y) {
		private Position add(Position position) {
			return new Position(x + position.x(), y + position.y());
		}

		@Override
		public String toString() {
			return "x:%s y:%s".formatted(x, y);
		}
	}

	@Override
	public long solvePart1(String input) {
		char[][] map = parseMap(input);
		Vector startVector = new Vector(findStartPosition(map), Direction.UP);

		try {
			return getVisitedPositions(startVector, map).size();
		} catch (CircleException e) {
			return 0;
		}
	}

	private static Set<Vector> guardTurnPositions;

	private Set<Position> getVisitedPositions(Vector currentVector, char[][] map) throws CircleException {
		guardTurnPositions = new HashSet<>();
		Set<Position> visitedPositions = new HashSet<>();
		while (inBounds(currentVector.position(), map)) {
			visitedPositions.add(currentVector.position());
			currentVector = simulateMove(currentVector.position(), currentVector.direction(), map);
		}
		return visitedPositions;
	}

	private Vector simulateMove(Position currentPosition, Direction currentDirection, char[][] map)
			throws CircleException {
		Position nextPosition = move(currentPosition, currentDirection);
		char symbol = getSymbol(nextPosition, map);
		if (symbol == '#' || symbol == 'O') {
			if (!guardTurnPositions.add(new Vector(currentPosition, currentDirection))) {
				throw new CircleException();
			}
			return simulateMove(currentPosition, rotate90(currentDirection), map);
		}
		return new Vector(nextPosition, currentDirection);
	}

	@Override
	public long solvePart2(String input) {
		long circles = 0;

		char[][] map = parseMap(input);
		Vector startVector = new Vector(findStartPosition(map), Direction.UP);

		Set<Position> visitedPositions = new HashSet<>();
		try {
			visitedPositions = getVisitedPositions(startVector, map);
		} catch (CircleException ignored) {
		}

		for (Position visitedPosition : visitedPositions) {
			try {
				getVisitedPositions(startVector, placeObstacle(visitedPosition, map));
			} catch (Exception exception) {
				circles++;
			}
		}

		return circles;
	}

	private char[][] placeObstacle(Position position, char[][] map) {
		char[][] modifiedMap = Arrays.stream(map).map(char[]::clone).toArray(char[][]::new);
		modifiedMap[position.y()][position.x()] = 'O';
		System.out.println("Placed at: " + position);
		return modifiedMap;
	}

	private Position findStartPosition(char[][] map) {
		for (int y = 0; y < map.length; y++) {
			for (int x = 0; x < map[y].length; x++) {
				if (map[y][x] == '^') {
					return new Position(x, y);
				}
			}
		}
		throw new IllegalStateException();
	}

	private char getSymbol(Position position, char[][] map) {
		if (!inBounds(position, map)) {
			return ' ';
		}
		return map[position.y()][position.x()];
	}

	private Direction rotate90(Direction direction) {
		return switch (direction) {
			case UP -> Direction.RIGHT;
			case RIGHT -> Direction.DOWN;
			case DOWN -> Direction.LEFT;
			case LEFT -> Direction.UP;
		};
	}

	private Position move(Position currentPos, Direction direction) {
		return currentPos.add(switch (direction) {
			case UP -> new Position(0, -1);
			case RIGHT -> new Position(1, 0);
			case DOWN -> new Position(0, 1);
			case LEFT -> new Position(-1, 0);
		});
	}

	private char[][] parseMap(String input) {
		String[] lines = input.split("\n");
		char[][] map = new char[lines.length][];
		for (int i = 0; i < lines.length; i++) {
			map[i] = lines[i].toCharArray();
		}
		return map;
	}

	private boolean inBounds(Position position, char[][] map) {
		return inBounds(position.x(), position.y(), map);
	}

	private boolean inBounds(int x, int y, char[][] map) {
		return 0 <= x && x < map[0].length && 0 <= y && y < map.length;
	}
}
