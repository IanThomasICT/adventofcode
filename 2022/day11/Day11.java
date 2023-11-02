import aoc22.Utils;

import java.util.Arrays;
import java.util.Comparator;
import java.util.List;
import java.util.Stack;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.Function;
import java.util.stream.Collectors;

public class Day11 {
    static class Monkey {
        Stack<Long> startingItems = new Stack<>();
        Function<Long, Long> operation;
        String operationString = "";
        long numOfInspections;
        int divisibleBy;
        int throwToIfTrue;
        int throwToIfFalse;

        Monkey(Stack<Long> items, Function<Long, Long> operation, Integer divisibleBy, Integer ifTrue, Integer ifFalse) {
            this.startingItems = items;
            this.operation = operation;
            this.divisibleBy = divisibleBy;
            this.throwToIfTrue = ifTrue;
            this.throwToIfFalse = ifFalse;
        }

        Monkey() {}

        @Override
        public String toString() {
            return "Monkey{" +
                    "startingItems=" + startingItems +
                    ", operation=" + operation +
                    ", divisibleBy=" + divisibleBy +
                    ", throwToIfTrue=" + throwToIfTrue +
                    ", throwToIfFalse=" + throwToIfFalse +
                    '}';
        }
    }

    // static String filePath = "input/day11-test.txt";
    static String filePath = "input/day11.txt";
    public static void main(String[] args) {
        // System.out.println(filePath);
        List<String> instructions = Utils.fileToStringList(filePath);
        partOne(instructions);
        partTwo(instructions);
    }

    public static void partOne(List<String> instructions) {
        List<Monkey> monkeys = new Stack<>();
        populateMonkeys(instructions, monkeys);
        int rounds = 20;
        for (int i = 0; i < rounds; i++) {
            processInspections(monkeys, 3, null);
        }

        Long levelOfMonkeyBusiness = multiplyTopTwoInspectors(monkeys);
        System.out.println("Part 1: " + levelOfMonkeyBusiness);
    }

    public static void partTwo(List<String> instructions) {
        List<Monkey> monkeys = new Stack<>();
        populateMonkeys(instructions, monkeys);
        int maxModulus = monkeys.stream()
                .mapToInt(m -> m.divisibleBy)
                .boxed()
                .reduce(1, (i,j) -> i * j);
        int rounds = 10000;
        for (int i = 0; i < rounds; i++) {
            processInspections(monkeys, 1, maxModulus);
        }

        Long levelOfMonkeyBusiness = multiplyTopTwoInspectors(monkeys);
        System.out.println("Part 2: " + levelOfMonkeyBusiness);
    }

    private static void processInspections(List<Monkey> monkeys, int worryDivisor, Integer maxModulus) {
        for (Monkey monkey : monkeys) {
            while (!monkey.startingItems.isEmpty()) {
                Long item = monkey.startingItems.pop();
                monkey.numOfInspections++;
                item = monkey.operation.apply(item);
                item /= worryDivisor;
                if (maxModulus != null) {
                    item %= maxModulus;
                }
                if (item % monkey.divisibleBy == 0) {
                    monkeys.get(monkey.throwToIfTrue).startingItems.push(item);
                } else {
                    monkeys.get(monkey.throwToIfFalse).startingItems.push(item);
                }
            }
        }
    }

    private static Long multiplyTopTwoInspectors(List<Monkey> monkeys) {
        return monkeys.stream()
                .map(m -> m.numOfInspections)
                .sorted(Comparator.reverseOrder())
                .limit(2)
                .reduce(1L,(i1, i2) -> i1 * i2);
    }

    private static void populateMonkeys(List<String> instructions, List<Monkey> monkeys) {
        for (int i = 0; i < instructions.size()-1; i+=7) {
            List<String> monkeyParams = instructions.subList(i, i + 6);
            Monkey monkey = new Monkey();
            monkey.startingItems = Arrays.stream(monkeyParams.get(1).split(":")[1].trim().split(","))
                    .map(String::trim)
                    .map(Long::valueOf)
                    .collect(Collectors.toCollection(Stack::new));
            addOperation(monkeyParams.get(2).split(":")[1].trim(), monkey);
            monkey.divisibleBy = Integer.parseInt(monkeyParams.get(3).split("divisible by")[1].trim());
            monkey.throwToIfTrue = Integer.parseInt(monkeyParams.get(4).split("monkey ")[1]);
            monkey.throwToIfFalse = Integer.parseInt(monkeyParams.get(5).split("monkey ")[1]);
            monkeys.add(monkey);
        }
    }

    private static void addOperation(String operationInput, Monkey monkey) {
        String[] operationOnOld = operationInput.split("old ")[1].split(" ");
        String operator = operationOnOld[0];
        int value = !operationOnOld[1].contains("old") ? Integer.parseInt(operationOnOld[1]) : -1; // Distinguish when operation includes old
        switch(operator) {
            case "+" -> {
                monkey.operation = i -> i + (value == -1 ? i : value);
                monkey.operationString = "increased by " + (value == -1 ? "itself" : value);
            }
            case "*" -> {
                monkey.operation = i -> i * (value == -1 ? i : value);
                monkey.operationString = "multiplied by " + (value == -1 ? "itself" : value);
            }
        };
    }


}
