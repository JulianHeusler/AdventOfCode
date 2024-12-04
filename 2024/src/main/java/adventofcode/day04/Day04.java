package adventofcode.day04;

import adventofcode.util.AbstractDay;

public class Day04 extends AbstractDay {

	@Override
	public long solvePart1(String input) {
		long sum = 0;
		String[] lines = input.split("\n");
		char[][] map = new char[lines.length][];
		for (int y = 0; y < lines.length; y++) {
			map[y] = lines[y].toCharArray();
		}

		for (int y = 0; y < map.length; y++) {
			for (int x = 0; x < map[y].length; x++) {
				if (verticalToRight(x, y, map)) {
					System.out.printf("x:%s y:%s%n", x, y);
					sum++;
				}
				if (horizontalDown(x, y, map)) {
					System.out.printf("x:%s y:%s%n", x, y);
					sum++;
				}
				if (diagonalDownRight(x, y, map)) {
					System.out.printf("x:%s y:%s%n", x, y);
					sum++;
				}
				if (diagonalDownLeft(x, y, map)) {
					System.out.printf("x:%s y:%s%n", x, y);
					sum++;
				}
			}
		}

		return sum;
	}

	private boolean horizontalDown(int x, int y, char[][] map) {
		if (!inBounds(x, y + 3, map)) {
			return false;
		}

		return map[y][x] == 'X' && map[y + 1][x] == 'M' && map[y + 2][x] == 'A' && map[y + 3][x] == 'S'
				|| map[y][x] == 'S' && map[y + 1][x] == 'A' && map[y + 2][x] == 'M' && map[y + 3][x] == 'X';
	}

	private boolean verticalToRight(int x, int y, char[][] map) {
		if (!inBounds(x + 3, y, map)) {
			return false;
		}

		return map[y][x] == 'X' && map[y][x + 1] == 'M' && map[y][x + 2] == 'A' && map[y][x + 3] == 'S'
				|| map[y][x] == 'S' && map[y][x + 1] == 'A' && map[y][x + 2] == 'M' && map[y][x + 3] == 'X';
	}

	private boolean diagonalDownRight(int x, int y, char[][] map) {
		if (!inBounds(x, y, map)) {
			return false;
		}
		if (!inBounds(x + 1, y + 1, map)) {
			return false;
		}
		if (!inBounds(x + 2, y + 2, map)) {
			return false;
		}
		if (!inBounds(x + 3, y + 3, map)) {
			return false;
		}

		return map[y][x] == 'X' && map[y + 1][x + 1] == 'M' && map[y + 2][x + 2] == 'A' && map[y + 3][x + 3] == 'S'
				|| map[y][x] == 'S' && map[y + 1][x + 1] == 'A' && map[y + 2][x + 2] == 'M' && map[y + 3][x + 3] == 'X';
	}

	private boolean diagonalDownLeft(int x, int y, char[][] map) {
		if (!inBounds(x, y, map)) {
			return false;
		}
		if (!inBounds(x - 1, y + 1, map)) {
			return false;
		}
		if (!inBounds(x - 2, y + 2, map)) {
			return false;
		}
		if (!inBounds(x - 3, y + 3, map)) {
			return false;
		}

		return map[y][x] == 'X' && map[y + 1][x - 1] == 'M' && map[y + 2][x - 2] == 'A' && map[y + 3][x - 3] == 'S'
				|| map[y][x] == 'S' && map[y + 1][x - 1] == 'A' && map[y + 2][x - 2] == 'M' && map[y + 3][x - 3] == 'X';
	}

	private boolean inBounds(int x, int y, char[][] map) {
		return 0 <= x && x < map[0].length && 0 <= y && y < map.length;
	}

	@Override
	public long solvePart2(String input) {
		return 0;
	}
}
