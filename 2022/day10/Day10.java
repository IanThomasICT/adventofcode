import java.util.ArrayList;
import java.util.List;
import java.util.stream.IntStream;

import static aoc22.Utils.fileToStringList;

public class Day10 {

    // static String filePath = "input/day10-test.txt";
    static String filePath = "input/day10.txt";
    public static void main(String[] args) {
        System.out.println(filePath);
        List<String> instructions = fileToStringList(filePath);
        List<Integer> cycleStartValues = partOne(instructions);
        partTwo(cycleStartValues);
    }

    public static List<Integer> partOne(List<String> instructions) {
        List<Integer> cycleStartValues = new ArrayList<>(List.of(1,1)); // padded to start indexing at 1, and first cycle starts at 1
        processCycleInstructions(instructions, cycleStartValues);
        Integer sumOfCycles = IntStream.iterate(20, i -> i <= 220, i -> i + 40)
                // .peek(i -> System.out.printf("Cycle %d: %d\n", i, cycleStartValues.get(i)))
                .map(i -> i * cycleStartValues.get(i))
                .sum();

        System.out.println("Part 1: " + sumOfCycles);
        return cycleStartValues;
        // Answer: 14820
    }

    public static void partTwo(List<Integer> cycleStartValues) {
        int window = 40;
        System.out.print("0\t");
        for (int i = 0; i < window; i++) {
            Integer cycleValue = cycleStartValues.get(i+1);
            printPixel(i % 40, cycleValue);
            if (i == window - 1 && window < 240) {
                System.out.printf(" %d\n%d\t", i, i+1);
                window += 40;
            } else if (i == window - 1) {
                System.out.printf(" %d", i, i+1);
            }
        }
        // Answer: RZEKEFHA
    }

    private static void printPixel(Integer cycle, Integer cycleValue) {
        System.out.print(Math.abs(cycle - cycleValue) <= 1 ?  "#" : ".");
    }

    private static void processCycleInstructions(List<String> instructions, List<Integer> cycleValues) {
        Integer x = 1;
        for (String line : instructions) {
            String[] parts = line.split(" ");
            String instruction = parts[0];
            switch (instruction) {
                case "addx" -> cycleValues.addAll(List.of(x,x += Integer.parseInt(parts[1])));
                case "noop" -> cycleValues.add(x);
            }
        }
    }
}
