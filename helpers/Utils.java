package aoc22;

import java.nio.file.Files;
import java.nio.file.InvalidPathException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.*;

public class Utils {
    public static final String ANSI_RESET = "\u001B[0m";
    public static final String ANSI_RED = "\u001B[31m";
    public static final String ANSI_GREEN = "\u001B[32m";
    public static final String ANSI_YELLOW = "\u001B[33m";

    public static void log(String str) {
        System.out.println(str);
    }

    public static void error(String str) {
        System.out.println(ANSI_RED + str + ANSI_RESET);
    }

    public static List<String> fileToStringList(String filePath) {
        try {
            return Files.lines(Path.of(filePath)).toList();
        } catch (Exception e) {
            error("Failed to read file " + filePath);
            e.printStackTrace();
            return new ArrayList<>();
        }
    }

    public static List<String> splitString(String input, int partitions) {
        List<String> result = new ArrayList<>();
        int partitionLength = input.length() / partitions;
        for (int i = 0; i < input.length() - partitionLength; i += partitionLength) {
            result.add(input.substring(i, i + partitionLength));
        }

        return result;
    }

    public static String nullToBlank(Object obj) {
        return Objects.nonNull(obj) ? obj.toString() : "";
    }

    public static Integer nullToZero(Integer obj) {
        return Objects.nonNull(obj) ? obj : 0;
    }

    public static <T>void drawNodes(T[][] grid, List<int[]> visitedNodes) {
        String[][] strGrid = new String[grid.length][grid[0].length];
        for (String[] row : strGrid) {
            Arrays.fill(row, ".");
        }

        for (int[] coord : visitedNodes) {
            strGrid[coord[0]][coord[1]] = "o";
        }
        printGrid(strGrid);
    }

    public static class Pair<T, S> {
        public T first;
        public S second;

        public Pair(T first, S second) {
            this.first = first;
            this.second = second;
        }

        @Override
        public String toString() {
            return "[" + first + "," + second +']';
        }
    }

    public static void printGrid(int[][] grid) {
        System.out.print("\r");
        for (int[] i : grid) {
            for (int j : i) {
                System.out.printf("%3d ", j);
                System.out.print("\b");
            }
            System.out.print("\n");
        }
        System.out.print("\n");
    }

    public static void printGrid(String[][] grid) {
        System.out.print("\r");
        for (String[] i : grid) {
            for (String j : i) {
                System.out.printf("%2s ", j);
                System.out.print("\b");
            }
            System.out.print("\n");
        }
        System.out.print("\n");
    }

    public static <T>void drawPath(T[][] grid, List<String> path) {
        String[][] strGrid = new String[grid.length][grid[0].length];
        int pathLength = path.size();
        for (String[] row : strGrid) {
            Arrays.fill(row, ".");
        }

        for (int i = 0; i < pathLength - 1; i++) {
            int[] fromCoord = Arrays.stream(path.get(i).replaceAll("[\\[\\]]", "").split(",")).map(String::trim).mapToInt(Integer::parseInt).toArray();
            int[] toCoord = Arrays.stream(path.get(i+1).replaceAll("[\\[\\]]", "").split(",")).map(String::trim).mapToInt(Integer::parseInt).toArray();
            if (i == 0) {
                strGrid[fromCoord[0]][fromCoord[1]] = "O";
            } else if (i == pathLength - 2) {
                strGrid[fromCoord[0]][fromCoord[1]] = "E";
            } else {
                strGrid[fromCoord[0]][fromCoord[1]] = getCardinalDirection(fromCoord, toCoord);
            }
        }
        printGrid(strGrid);
    }

    /** Takes two x,y coordinates in the format "[x,y]" and returns a direction arrow based on the cardinal direction   of change*/
    private static String getCardinalDirection(int[] fromCoord, int[] toCoord) {
        int[] moveCoords = subtract(fromCoord, toCoord);
        return switch (moveCoords[0]) {
            case 0 -> moveCoords[1] < 0 ? "<" : ">";
            case 1,-1 -> moveCoords[0] < 0 ? "^" : "v";
            default -> ".";
        };
    }

    /** Subtract a FROM b */
    public static int[] subtract(int[] a, int[] b) {
        for (int i = 0; i < a.length; i++) {
            b[i]-=a[i];
        }
        return b;
    }

    public static String coordToString(int i, int j) {
        return String.format("[%d,%d]", i, j);
    }

}
