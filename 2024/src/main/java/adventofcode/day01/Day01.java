package adventofcode.day01;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

import adventofcode.util.AbstractDay;

public class Day01 extends AbstractDay {


	public long solvePart1(String input) {
		Tuple tuple = parseInput(input);
		long distances = 0;
		for (int i = 0; i < tuple.list1.size(); i++) {
			distances += Math.abs(tuple.list1.stream().sorted().toList().get(i) - tuple.list2.stream().sorted().toList().get(i));
		}
		return distances;
	}

	@Override
	public long solvePart2(String input) {
		long similarityScore = 0;
		Tuple tuple = parseInput(input);
		Map<Integer, List<Integer>> collected = tuple.list2.stream().collect(Collectors.groupingBy(integer -> integer));
		for (int i = 0; i < tuple.list1.size(); i++) {
			int number = tuple.list1.get(i);
			similarityScore+= (long) number * collected.getOrDefault(number, List.of()).size();
		}
		return similarityScore;
	}

	private Tuple parseInput(String input) {
		ArrayList<Integer> listLeft = new ArrayList<>();
		ArrayList<Integer> listRight = new ArrayList<>();
		for (String line : input.split("\n")) {
			String[] split = line.split("   ");
			listLeft.add(Integer.parseInt(split[0]));
			listRight.add(Integer.parseInt(split[1]));
		}
		return new Tuple(listLeft, listRight);
	}

	private record Tuple(List<Integer> list1, List<Integer> list2) {
	}
}