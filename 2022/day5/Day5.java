import aoc22.Utils;

import java.nio.file.Files;
import java.nio.file.InvalidPathException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.*;
import java.util.stream.Collectors;

public class Day5 {
    //     [V] [G]             [H]
    // [Z] [H] [Z]         [T] [S]
    // [P] [D] [F]         [B] [V] [Q]
    // [B] [M] [V] [N]     [F] [D] [N]
    // [Q] [Q] [D] [F]     [Z] [Z] [P] [M]
    // [M] [Z] [R] [D] [Q] [V] [T] [F] [R]
    // [D] [L] [H] [G] [F] [Q] [M] [G] [W]
    // [N] [C] [Q] [H] [N] [D] [Q] [M] [B]
    // 1   2   3   4   5   6   7   8   9
    public static Map<String, Deque<String>> containerStart = Map.of(
            "1", new ArrayDeque<String>(List.of("Z","P","B","Q","M","D","N")),
            "2", new ArrayDeque<String>(List.of("V", "H", "D", "M", "Q", "Z", "L", "C")),
            "3", new ArrayDeque<String>(List.of("G","Z","F","V","D","R","H","Q")),
            "4", new ArrayDeque<String>(List.of("N","F","D","G","H")),
            "5", new ArrayDeque<String>(List.of("Q","F","N")),
            "6", new ArrayDeque<String>(List.of("T","B","F","Z","V","Q","D")),
            "7", new ArrayDeque<String>(List.of("H","S","V","D","Z","T","M","Q")),
            "8", new ArrayDeque<String>(List.of("Q","N","P","F","G","M")),
            "9", new ArrayDeque<String>(List.of("M","R","W","B"))
    );

    public static void main(String[] args) {
        String filePath = args[0];
        List<String> lines = Utils.fileToStringList(filePath);
        int i = 0;
        while (!lines.get(i).contains("move")) {
            i++;
        }
        // partOne(lines, i);
        partTwo(lines, i);
    }

    public static void partOne(List<String> lines, int startLine) {
        Map<String, Deque<String>> containers = new HashMap<>(containerStart);
        System.out.println(containers);
        for (String line : lines.subList(startLine,lines.size())) {
            String[] instructions = line.split(" ");
            int quantity = Integer.parseInt(instructions[1]);
            String fromIdx = instructions[3];
            String toIdx = instructions[5];
            moveContainersIndividually(quantity, containers.get(fromIdx), containers.get(toIdx));
        }
        System.out.println("Part One: " + containers.entrySet().stream().sorted(Comparator.comparing(Map.Entry::getKey)).map(s -> s.getValue().getFirst()).collect(Collectors.joining("")));
    }

    public static void partTwo(List<String> lines, int startLine) {
        Map<String, Deque<String>> containers = new HashMap<>();
        containers.putAll(containerStart);
        System.out.println(containers);
        for (String line : lines.subList(startLine,lines.size())) {
            String[] instructions = line.split(" ");
            int quantity = Integer.parseInt(instructions[1]);
            String fromIdx = instructions[3];
            String toIdx = instructions[5];
            moveContainersByGroup(quantity, containers.get(fromIdx), containers.get(toIdx));
            // renderContainers(containers);
        }
        System.out.println("Part Two: " + containers.entrySet().stream().sorted(Comparator.comparing(Map.Entry::getKey)).map(s -> s.getValue().getFirst()).collect(Collectors.joining("")));
    }

    public static void moveContainersIndividually(int quantity, Deque<String> fromStack, Deque<String> toStack) {
        try {
            for (int i = 0; i < quantity; i++) {
                toStack.push(fromStack.pop());
            }
        } catch (NoSuchElementException e) {
            Utils.error(String.format("Error occurred moving %s from %s to %s", fromStack.getFirst(), fromStack, toStack));
            e.printStackTrace();
        }
    }

    public static void moveContainersByGroup(int quantity, Deque<String> fromStack, Deque<String> toStack) {
        try {
            Deque<String> tmp = new ArrayDeque<>();
            for (int i = 0; i < quantity; i++) {
                tmp.push(fromStack.pop());
            }
            while (!tmp.isEmpty()) {
                toStack.push(tmp.pop());
            }
        } catch (NoSuchElementException e) {
            Utils.error(String.format("Error occurred moving %s from %s to %s", fromStack.getFirst(), fromStack, toStack));
            e.printStackTrace();
        }
    }

    public static void renderContainers(Map<String, Deque<String>> containers) {
        Map<String, Deque<String>> containerState = new HashMap<String, Deque<String>>();
        containerState.putAll(containers);
        while (oneIsNotEmpty(containerState)) {
            System.out.printf("\r [%s]", containerState.get(1).isEmpty() ? "    " : containerState.get(1).pop());
            System.out.printf(" [%s]", containerState.get(2).isEmpty() ? "    " : containerState.get(2).pop());
            System.out.printf(" [%s]", containerState.get(3).isEmpty() ? "    " : containerState.get(3).pop());
            System.out.printf(" [%s]", containerState.get(4).isEmpty() ? "    " : containerState.get(4).pop());
            System.out.printf(" [%s]", containerState.get(5).isEmpty() ? "    " : containerState.get(5).pop());
            System.out.printf(" [%s]", containerState.get(6).isEmpty() ? "    " : containerState.get(6).pop());
            System.out.printf(" [%s]", containerState.get(7).isEmpty() ? "    " : containerState.get(7).pop());
            System.out.printf(" [%s]", containerState.get(8).isEmpty() ? "    " : containerState.get(8).pop());
            System.out.printf(" [%s]", containerState.get(9).isEmpty() ? "    " : containerState.get(9).pop());
            System.out.print("\n");
        }
        System.out.println(" [1] [2] [3] [4] [5] [6] [7] [8] [9]");

    }

    private static boolean oneIsNotEmpty(Map<String, Deque<String>> map) {
        return map.values().stream().anyMatch(d -> !d.isEmpty());
    }
}


