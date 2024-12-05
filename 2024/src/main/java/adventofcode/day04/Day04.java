package adventofcode.day04;

import java.util.List;

import adventofcode.util.AbstractDay;

public class Day04 extends AbstractDay {

	private final List<char[][]> allowedConfigurations = List.of(
			new char[][]{
					{'M', '.', 'S'},
					{'.', 'A', '.'},
					{'M', '.', 'S'}
			},
			new char[][]{
					{'M', '.', 'M'},
					{'.', 'A', '.'},
					{'S', '.', 'S'}
			}, new char[][]{
					{'S', '.', 'M'},
					{'.', 'A', '.'},
					{'S', '.', 'M'}
			}, new char[][]{
					{'S', '.', 'S'},
					{'.', 'A', '.'},
					{'M', '.', 'M'}
			}
	);

	@Override
	public long solvePart1(String input) {
		long sum = 0;
		char[][] map = getMap(input);

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

	private static char[][] getMap(String input) {
		String[] lines = input.split("\n");
		char[][] map = new char[lines.length][];
		for (int y = 0; y < lines.length; y++) {
			map[y] = lines[y].toCharArray();
		}
		return map;
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
		long sum = 0;
		char[][] map = getMap(input);
		for (int y = 0; y < map.length; y++) {
			for (int x = 0; x < map[y].length; x++) {
				char current = map[y][x];
				if (current == 'A') {
					continue;
				}

				if (isValidXMasCross(x, y, map)) {
					sum++;
				}
			}
		}

		return sum;
	}

	private boolean isValidXMasCross(int x, int y, char[][] map) {
		if (!inBounds(x + 2, y + 2, map)) {
			return false;
		}

		char[][] currentConfiguration = new char[3][3];
		for (int i = 0; i < 3; i++) {
			currentConfiguration[i][0] = map[y + i][x];
			currentConfiguration[i][1] = map[y + i][x + 1];
			currentConfiguration[i][2] = map[y + i][x + 2];
		}

		for (char[][] allowedConfiguration : allowedConfigurations) {
			if (matches(currentConfiguration, allowedConfiguration)) {
				// System.out.println(Arrays.deepToString(currentConfiguration));
				return true;
			}
		}
		return false;
	}

	private boolean matches(char[][] currentConfiguration, char[][] allowedConfiguration) {
		for (int y = 0; y < currentConfiguration.length; y++) {
			for (int x = 0; x < currentConfiguration[y].length; x++) {
				if (allowedConfiguration[y][x] == '.') {
					continue;
				}
				if (currentConfiguration[y][x] != allowedConfiguration[y][x]) {
					return false;
				}
			}
		}
		return true;
	}
}
