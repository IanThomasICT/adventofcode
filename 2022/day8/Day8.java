import aoc22.Utils;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.InvalidPathException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.*;
import java.util.stream.IntStream;

import static aoc22.Utils.*;

public class Day8 {
    public static record Tree (int x, int y, int height) {}

    public static void main(String[] args) throws IOException {
        String filePath = args[0];
        List<String> instructions = fileToStringList(filePath);
        partOne(instructions);
        partTwo(instructions);
    }

    public static void partOne(List<String> instructions) {
        int width = instructions.get(0).length();
        int height = instructions.size();
        int visibleTrees = (width - 2)*2 + (height - 2)*2  + 4;
        Tree[][] grid = new Tree[height][width];
        fillArray(instructions, width, grid);

        Tree tree;
        for (int i = 1; i < width - 1; i++) {
            for (int j = 1; j < height - 1 ; j++) {
                tree = grid[i][j];
                boolean visible = false;
                // if isVisible, add to list
                if (visibleFromLeft(grid, tree))
                    visible = true;
                else if (visibleFromRight(grid, tree))
                    visible = true;
                else if (visibleFromTop(grid, tree))
                    visible = true;
                else if (visibleFromBottom(grid, tree))
                    visible = true;

                if (visible)
                    visibleTrees++;
            }
        }
        System.out.println("Part 1: " + visibleTrees + " trees are visible from the edge");
    }

    public static void partTwo(List<String> instructions) {
        int width = instructions.get(0).length();
        int height = instructions.size();
        Tree[][] grid = new Tree[height][width];
        fillArray(instructions, width, grid);

        // find tree with the most scenic score
        Tree tree = new Tree(-1, -1, 0);
        int maxScenicScore = 0;
        Tree maxTree = tree;
        for (int i = 1; i < width - 1; i++) {
            for (int j = 0; j < height - 1; j++) {
                tree = grid[i][j];
                int scenicScore = getScenicScore(grid, tree);
                if (scenicScore > maxScenicScore) {
                    maxScenicScore = scenicScore;
                    maxTree = tree;
                }
            }
        }
        System.out.println("Part 2: " + maxTree + " has the highest scenic score of " + maxScenicScore);
    }

    private static int getScenicScore(Tree[][] grid, Tree tree) {
        List<Integer> multipliers = new ArrayList<>();
        for (int x = tree.x - 1; x >= 0; x--) {
            if (grid[tree.y][x].height >= tree.height || x == 0) {
                multipliers.add(Math.abs(tree.x - x));
                break;
            }
        }
        for (int x = tree.x + 1; x < grid[0].length; x++) {
            if (grid[tree.y][x].height >= tree.height || x == grid[0].length - 1) {
                multipliers.add(Math.abs(tree.x - x));
                break;
            }
        }
        for (int y = tree.y - 1; y >= 0; y--) {
            if (grid[y][tree.x].height >= tree.height || y == 0) {
                multipliers.add(Math.abs(tree.y - y));
                break;
            }
        }
        for (int y = tree.y + 1; y < grid.length; y++) {
            if (grid[y][tree.x].height >= tree.height || y == grid.length - 1) {
                multipliers.add(Math.abs(tree.y - y));
                break;
            }
        }
        return multipliers.stream().mapToInt(Integer::intValue).reduce(1, (s1, s2) -> s1 * s2);
    }

    private static boolean visibleFromLeft(Tree[][] grid, Tree tree) {
        for (int x = 0; x < tree.x; x++) {
            if (grid[tree.y][x].height >= tree.height) {
                return false;
            }
        }
        return true;
    }

    private static boolean visibleFromRight(Tree[][] grid, Tree tree) {
        for (int x = grid[0].length - 1; x > tree.x; x--) {
            if (grid[tree.y][x].height >= tree.height) {
                return false;
            }
        }
        return true;
    }

    private static boolean visibleFromTop(Tree[][] grid, Tree tree) {
        for (int y = 0; y < tree.y; y++) {
            if (grid[y][tree.x].height >= tree.height) {
                return false;
            }
        }
        return true;
    }

    private static boolean visibleFromBottom(Tree[][] grid, Tree tree) {
        for (int y = grid.length - 1; y > tree.y; y--) {
            if (grid[y][tree.x].height >= tree.height) {
                return false;
            }
        }
        return true;
    }

    private static void fillArray(List<String> instructions, int width, Tree[][] array) {
        int i = 0;
        for (String line : instructions) {
            int[] row = line.chars().map(s -> s - 48).toArray();
            for (int j = 0; j < width; j++) {
                array[i][j] = new Tree(j,i,row[j]);
            }
            i++;
        }
    }



}
