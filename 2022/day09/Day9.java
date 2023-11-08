import aoc22.Utils;

import java.util.*;
import java.util.stream.IntStream;

import static java.util.Map.Entry.comparingByValue;

public class Day9 {
    public static class Rope {
        Knot head;
        LinkedList<Knot> knots = new LinkedList<>();

        public Rope(int size) {
            this.head = new Knot(0,0);
            knots.add(new Knot(0,0,head));
            IntStream.range(1,size).forEach(i -> knots.add(new Knot(0,0, knots.get(i-1)))); // Create linked list
        }
        
        public Knot getKnot(int index) {
            return knots.get(index);
        }

        public Knot getHead() {
            return head;
        }

        public Knot getTail() {
            return knots.getLast();
        }

        public void moveHead(String direction) {
            Knot head = getHead();
            head.move(direction);
        }

        @Override
        public String toString() {
            return knots.stream().toList().toString();
        }

        public void moveKnots() {
            for (Knot knot : knots) {
                knot.moveToClosestPositionToParent();
            }
        }
    }

    public static class Knot {
        int x, y;
        Knot parent;

        public Knot(int x, int y) {
            this.x = x;
            this.y = y;
        }

        public Knot(int x, int y, Knot parent) {
            this.x = x;
            this.y = y;
            this.parent = parent;
        }

        public List<Integer> getCoords() {
            return List.of(x,y);
        }

        public void moveToClosestPositionToParent() {
            List<Integer> closestPositionToParent = findClosestPosition();
            this.x = closestPositionToParent.get(0);
            this.y = closestPositionToParent.get(1);
        }

        public List<Integer> findClosestPosition() {
            if (Math.abs(parent.x - x) >= 2 || Math.abs(parent.y - y) >= 2) {
                // get all possible moves
                int[][] possibleMoves = new int[][]{
                        new int[]{x, y + 1},
                        new int[]{x, y - 1},
                        new int[]{x + 1, y},
                        new int[]{x - 1, y},
                        new int[]{x + 1, y + 1},
                        new int[]{x + 1, y - 1},
                        new int[]{x - 1, y - 1},
                        new int[]{x - 1, y + 1}
                };

                // Get closest position to parent by using Distance formula
                List<Integer> closestMove = Arrays.stream(possibleMoves)
                        .flatMap(arr -> Map.of(List.of(arr[0], arr[1]), getDistanceFromParent(arr[0], arr[1])).entrySet().stream())
                        .min(comparingByValue())
                        .get().getKey();

                return closestMove;
            }
            return getCoords();
        }

        public double getDistanceFromParent(int x, int y) {
            return Math.sqrt(Math.pow(parent.x - x, 2) + Math.pow(parent.y - y,2));
        }

        public void move(String direction) {
            switch(direction) {
                case "U" -> this.x++;
                case "D" -> this.x--;
                case "L" -> this.y--;
                case "R" -> this.y++;
            }
        }
    }

    public static void main(String[] args) {
        String filePath = args[0];
        List<String> instructions = Utils.fileToStringList(filePath);
        partOne(instructions);
        partTwo(instructions);
    }

    public static void partOne(List<String> instructions) {
        int steps = 0;
        String direction;
        Knot head = new Knot(0,0);
        Knot tail = new Knot(0,0, head);
        TreeSet<String> path = new TreeSet<>();
        path.add(tail.getCoords().toString());

        for (String line : instructions) {
            String[] parts = line.split(" ");
            direction = parts[0];
            steps = Integer.parseInt(parts[1]);

            for (int i = 0; i < steps; i++) {
                head.move(direction);
                tail.moveToClosestPositionToParent();
                path.add(tail.getCoords().toString());
            }
        }
        System.out.println("Part 1: " + path.size());
    }

    public static void partTwo(List<String> instructions) {
        int steps = 0, lineNum = 0;
        String direction;
        Rope rope = new Rope(9);
        TreeSet<String> path = new TreeSet<>();
        path.add(rope.getHead().getCoords().toString());
        for (String line : instructions) {
            lineNum++;
            String[] parts = line.split(" ");
            direction = parts[0];
            steps = Integer.parseInt(parts[1]);

            for (int i = 0; i < steps; i++) {
                rope.moveHead(direction);
                rope.moveKnots();
                path.add(rope.getTail().getCoords().toString());
            }
        }
        System.out.println("Part 2: " + path.size());
    }
    /** Answers:
     * 6087 - PART ONE
     * 2493 - PART TWO
    */
}


