package adventofcode.day13;

import adventofcode.util.AbstractDay;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

import static org.junit.jupiter.api.Assumptions.assumeTrue;

public class Day13 extends AbstractDay {

    @Override
    public long solvePart1(String input) {
        List<List<String>> notes = getNotes(input);
        long result = 0;

        for (List<String> note : notes) {
            assumeTrue(note.size() % 2 != 0);
            assumeTrue(note.getFirst().length() % 2 != 0);

            Optional<Integer> columnIndex = getColumnIndex(note);
            if (columnIndex.isPresent()) {
                System.out.println("col += " + (columnIndex.get() + 1));
                result += (columnIndex.get() + 1);
            }

            Optional<Integer> rowIndex = getRowIndex(note);
            if (rowIndex.isPresent()) {
                System.out.println("row += " + (rowIndex.get() + 1));
                result += (100L * (rowIndex.get() + 1));
            }
        }
        return result;
    }

    private Optional<Integer> getRowIndex(List<String> note) {
        int rowCount = note.size();
//        if (note.get(0).equals(note.get(rowCount - 2))) {
//            return Optional.of((rowCount / 2) - 1);
//        }
//        if (note.get(1).equals(note.get(rowCount - 1))) {
//            return Optional.of(rowCount / 2);
//        }
        if (isMirrored(note.subList(0, rowCount - 1))) {
            return Optional.of((rowCount / 2) - 1);
        }
        if (isMirrored(note.subList(1, rowCount))) {
            return Optional.of(rowCount / 2);
        }
        return Optional.empty();
    }

    private boolean isMirrored(List<String> note) {
        assumeTrue(note.size() % 2 == 0);
        int middle = note.size() / 2;
        List<String> mirror = note.subList(middle, note.size()).reversed();
        for (int i = 0; i < middle; i++) {
            if (!note.get(i).equals(mirror.get(i))) {
                return false;
            }
        }
        return true;
    }

    private Optional<Integer> getColumnIndex(List<String> note) {
        int columnCount = note.getFirst().length();
        List<String> columns = new ArrayList<>();
        for (int i = 0; i < columnCount; i++) {
            columns.add(getColumn(i, note));
        }

        if (isMirrored(columns.subList(0, columns.size() - 1))) {
            return Optional.of((columnCount / 2) - 1);
        }
        if (isMirrored(columns.subList(1, columns.size()))) {
            return Optional.of(columnCount / 2);
        }

//        if (getColumn(0, note).equals(getColumn(columnCount - 2, note))) {
//            return Optional.of((columnCount / 2) - 1);
//        }
//        if (getColumn(1, note).equals(getColumn(columnCount - 1, note))) {
//            return Optional.of(columnCount / 2);
//        }
        return Optional.empty();
    }

    private String getColumn(int columnIndex, List<String> note) {
        return note.stream().map(line -> line.charAt(columnIndex))
                .map(Object::toString)
                .collect(Collectors.joining());
    }

    private List<Character> getColumnChars(int columnIndex, List<String> note) {
        return note.stream().map(line -> line.charAt(columnIndex)).toList();
    }


    public long solvePart1Old(String input) {
        List<List<String>> notes = getNotes(input);
        long result = 0;

        for (List<String> note : notes) {
            if (note.isEmpty()) {
                continue;
            }

            int lengthHorizonzal = 0;
            Optional<Integer> horizontal = findHorizontalReflection(note);
            if (horizontal.isPresent()) {
                lengthHorizonzal = lengthHorizonzal(horizontal.get(), note);
            }
            int lengthVertical = 0;
            Optional<Integer> vertical = findVerticalReflection(note);
            if (vertical.isPresent()) {
                lengthVertical = lengthVertical(vertical.get(), note);
            }

            if (lengthHorizonzal == 0 && lengthVertical == 0) {
                //continue;
            }

            if (lengthHorizonzal > lengthVertical) {
                // if (isPerfectReflectionHorizontal(lengthHorizonzal, note)) {
                System.out.println("horizontal: " + (lengthHorizonzal + 1));
                result += (100L * (lengthHorizonzal + 1));
            } else {
                // if (isPerfectReflectionColumn(lengthVertical, note)) {
                System.out.println("vertical: " + (lengthVertical + 1));
                result += (lengthVertical + 1);
            }
        }

        return result;
    }

    @Override
    public long solvePart2(String input) {
        return 0;
    }


    private boolean isPerfectReflectionColumn(int columnIndex, List<String> note) {
        int columnsCount = note.getFirst().length();
        assumeTrue(columnsCount % 2 == 1);
        return (columnIndex) * 2 == columnsCount - 1
                || (columnIndex + 1) * 2 == columnsCount - 1;
    }

    private boolean isPerfectReflectionHorizontal(int rowIndex, List<String> note) {
        int rowCount = note.size();
        assumeTrue(rowCount % 2 == 1);
        return (rowIndex) * 2 == rowCount - 1
                || (rowIndex + 1) * 2 == rowCount - 1;
    }

    public Optional<Integer> findVerticalReflection(List<String> note) {
        for (int i = 0; i < note.getFirst().length() - 1; i++) {
            int columnIndex = i;
            List<Character> column1 = note.stream().map(line -> line.charAt(columnIndex)).toList();
            List<Character> column2 = note.stream().map(line -> line.charAt(columnIndex + 1)).toList();
            if (equalColumns(column1, column2)) {
                return Optional.of(i);
            }
        }
        return Optional.empty();
    }

    private boolean equalColumns(List<Character> col1, List<Character> col2) {
        if (col1.size() != col2.size()) {
            return false;
        }
        for (int x = 0; x < col1.size(); x++) {
            if (!col1.get(x).equals(col2.get(x))) {
                return false;
            }
        }
        return true;
    }

    public Optional<Integer> findHorizontalReflection(List<String> note) {
        for (int i = 0; i < note.size() - 1; i++) {
            if (note.get(i).equals(note.get(i + 1))) {
                return Optional.of(i);
            }
        }
        return Optional.empty();
    }

    public int lengthVertical(int i, List<String> note) {
        int j = i + 1;
        int result = 0;
        while (i > 0 && j < note.getFirst().length()) {
            int columnIndex = i;
            List<Character> list1 = note.stream().map(s -> s.charAt(columnIndex)).toList();
            int finalJ = j;
            List<Character> list2 = note.stream().map(s -> s.charAt(finalJ)).toList();
            if (!list1.equals(list2)) {
                return result;
            }
            result++;
            i--;
            j++;
        }
        return result;
    }

    public int lengthHorizonzal(int i, List<String> note) {
        int j = i + 1;
        int result = 0;
        while (i > 0 && j < note.size()) {
            if (!note.get(i).equals(note.get(j))) {
                return result;
            }
            result++;
            i--;
            j++;
        }
        return result;
    }

    public List<List<String>> getNotes(String input) {
        List<List<String>> notes = new ArrayList<>();
        List<String> note = new ArrayList<>();

        for (String line : input.split("\n")) {
            if (line.isBlank()) {
                notes.add(note);
                note = new ArrayList<>();
            } else {
                note.add(line);
            }
        }
        if (!note.isEmpty()) {
            notes.add(note);
        }
        return notes;
    }
}
