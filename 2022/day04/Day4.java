import aoc22.Utils;

import java.nio.file.Files;
import java.nio.file.InvalidPathException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Day4 {
    public static void main(String[] args) {
        String filePath = args[0];
        partOne(filePath);
        partTwo(filePath);
    }

    public static void partOne(String filePath) {
        List<String> totalAssignments = Utils.fileToStringList(filePath);
        int totalOverlappedShifts = 0;
        for (String assignments : totalAssignments) {
            String[] shifts = assignments.split(",");
            int[] firstStartEnd = Arrays.stream(shifts[0].split("-")).mapToInt(Integer::valueOf).toArray();
            int[] secondStartEnd = Arrays.stream(shifts[1].split("-")).mapToInt(Integer::valueOf).toArray();
            if (firstShiftInSecond(firstStartEnd, secondStartEnd) || secondShiftInFirst(firstStartEnd, secondStartEnd)) {
                totalOverlappedShifts++;
            }            
        }

        Utils.log("Total Encompassed Shifts: " + totalOverlappedShifts);
    }

    public static void partTwo(String filePath) {
        List<String> totalAssignments = Utils.fileToStringList(filePath);
        int totalOverlappedShifts = 0;
        for (String assignments : totalAssignments) {
            String[] shifts = assignments.split(",");
            int[] firstStartEnd = Arrays.stream(shifts[0].split("-")).mapToInt(Integer::valueOf).toArray();
            int[] secondStartEnd = Arrays.stream(shifts[1].split("-")).mapToInt(Integer::valueOf).toArray();
            if (shiftsOverlap(firstStartEnd, secondStartEnd)) {
                // System.out.println(shifts[0] + " overlaps " + shifts[1]);
                totalOverlappedShifts++;
            } else if (firstShiftInSecond(firstStartEnd, secondStartEnd) || secondShiftInFirst(firstStartEnd, secondStartEnd)) {
                totalOverlappedShifts++;
            } else {
                System.out.println(shifts[0] + " does NOT overlap " + shifts[1]);
            }
        }

        Utils.log("Total Overlapped Shifts: " + totalOverlappedShifts);
    }

    public static boolean firstShiftInSecond(int[] first, int[] second) {
        return first[0] >= second[0] && first[1] <= second[1];
    }

    public static boolean secondShiftInFirst(int[] first, int[] second) {
        return first[0] <= second[0] && first[1] >= second[1];
    }

    // 6-6,4-6
    public static boolean shiftsOverlap(int[] first, int[] second) {
        return (first[0] == second[0] || first[0] == second[1] || first[1] == second[0] || first[1] == second[1]) || 
               (first[0] < second[0] && first[1] > second[0]) ||
               (first[0] < second[1] && first[1] > second[1]);
    }
}


