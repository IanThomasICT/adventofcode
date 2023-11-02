import aoc22.Utils;

import java.util.*;

public class Day13 {
    static String filePath = "src/input/day13-test.txt";
    // static String filePath = "src/input/day13.txt";
    public static void main(String[] args) throws InterruptedException {
        List<String> lines = Utils.fileToStringList(filePath);
        partOne(lines);
        // partTwo(lines);
    }

    public static void partOne(List<String> lines) throws InterruptedException {
        int rightPairs = 0;
        for (int i = 0; i < lines.size()-1; i+=3) {
            String left = lines.get(i).substring(1,lines.get(i).length()-1);    // trim brackets wrapping line
            String right = lines.get(i+1).substring(1,lines.get(i+1).length()-1);
            if (stringCompare(left, right)) {
                System.out.printf("Pair %d is a right pair\n", (i/3)+1);
                rightPairs++;
            }
        }
        System.out.println("Part 1: " + rightPairs + " right pairs");
    }


    /** conditions
     *  1. two integers, left !> right
     *  2.
    */
    public static boolean stringCompare(String leftStr, String rightStr) {
        boolean leftLonger = leftStr.length() < rightStr.length();
        int lenDiff = Math.abs(leftStr.length() - rightStr.length());
//        System.out.println("Left: " + leftStr.length() + ", Right: " + rightStr.length());
        for (int i = 0; i < Math.min(leftStr.length(), rightStr.length()); i++) {
//            System.out.println(i);
            char left = leftStr.charAt(i), right = rightStr.charAt(i);
            if (left == right) {
                continue; // [, same int, ],
            }

            if (isDigit(left) && isDigit(right) && left > right) {
                return false;
            }

            if (left == '[' && isDigit(right)) {
                String leftArr = getArray(leftStr.substring(i)); // "[...]", right = '#'
                int firstLeftValue = leftArr.chars().filter(Day13::isDigit).findFirst().orElse(58);
                return firstLeftValue > right;
            }

            // Last char of right string
            if (leftLonger && i == rightStr.length()-1) {
                return false; // Right ran out faster than left
            }
        }
        return true;
    }

    public static boolean isDigit(char a) {
        return a > 47 && a < 58;
    }

    public static boolean isDigit(int a) {
        return a > 47 && a < 58;
    }

    public static String getArray(String s) {
        int b = 0, i = 0;
        for (char c : s.toCharArray()) {
            switch(c) {
                case '[' -> b++;
                case ']' -> b--;
                default -> {}
            }
            if (b == 0) {
                return s.substring(0,i+1);
            }
            i++;
        }
        return s;
    }
}
