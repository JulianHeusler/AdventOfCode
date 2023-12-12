package adventofcode.day11;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.IntStream;

import adventofcode.util.AbstractDay;

public class Day11 extends AbstractDay {

	record Position(int y, int x) {
	}

	@Override
	public int solvePart1(String input) {
		List<Position> galaxies = getGalaxiesPositions(input);
		List<Position> galaxiesCopy = new ArrayList<>(galaxies);

		int lengths = 0;
		for (Position galaxy1 : galaxies) {
			galaxiesCopy.remove(galaxy1);
			for (Position galaxy2 : galaxiesCopy) {
				lengths += calculateDistance(galaxy1, galaxy2, galaxies);
			}
		}
		return lengths;
	}

	private int calculateDistance(Position a, Position b, List<Position> galaxies) {
		int dx = Math.abs(a.x - b.x);
		int dy = Math.abs(a.y - b.y);

		long emptyColumns = IntStream.range(Math.min(a.x, b.x), Math.max(a.x, b.x))
				.filter(i -> isEmptyColumn(i, galaxies))
				.count();

		long emptyRows = IntStream.range(Math.min(a.y, b.y), Math.max(a.y, b.y))
				.filter(i -> isEmptyRow(i, galaxies))
				.count();

		int result = dx + dy + (int) emptyColumns + (int) emptyRows;
		return result;
	}

	@Override
	public int solvePart2(String input) {
		return 0;
	}


	private static boolean isEmptyRow(int rowNumber, List<Position> galaxies) {
		return galaxies.stream()
				.map(Position::y)
				.noneMatch(y -> y == rowNumber);
	}

	private static boolean isEmptyColumn(int columnNumber, List<Position> galaxies) {
		return galaxies.stream()
				.map(Position::x)
				.noneMatch(x -> x == columnNumber);
	}

	private List<Position> getGalaxiesPositions(String input) {
		List<Position> galaxies = new ArrayList<>();
		String[] split = input.split("\n");
		for (int y = 0; y < split.length; y++) {
			for (int x = 0; x < split[0].length(); x++) {
				char current = split[y].charAt(x);
				if (current == '#') {
					galaxies.add(new Position(y, x));
				}
			}
		}
		return galaxies;
	}
}
