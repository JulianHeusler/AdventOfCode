package adventofcode.day01;

import adventofcode.util.AbstractDay;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day01 extends AbstractDay {


    public long solvePart1(String input) {
        Input input1 = parseInput(input);
        long distances = 0;
        for (int i = 0; i < input1.list1.size(); i++) {
            distances+= Math.abs(input1.list1.get(i) - input1.list2.get(i));
        }
        return distances;
    }

    @Override
    public long solvePart2(String input) {
        return 0;
    }

    private Input parseInput(String input){
        ArrayList<Integer> listLeft = new ArrayList<>();
        ArrayList<Integer> listRight = new ArrayList<>();
        for (String line : input.split("\n")) {
            String[] split = line.split("   ");
            listLeft.add(Integer.parseInt(split[0]));
            listRight.add(Integer.parseInt(split[1]));
        }
        return new Input(listLeft.stream().sorted().toList(), listRight.stream().sorted().toList());
    }

    private record Input(List<Integer> list1, List<Integer>list2) {}
}