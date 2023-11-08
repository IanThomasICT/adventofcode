import aoc22.Utils;

import java.nio.file.Files;
import java.nio.file.InvalidPathException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Day6 {
        public static void main(String[] args) {
        String filePath = args[0];
        String input = Utils.fileToStringList(filePath).get(0);
        int i = 0;
        Set<String> letters = new HashSet<>(Arrays.asList(input.substring(i, i+=4).split("")));
        while (letters.size() != 4 && i < input.length() - 4) {
            letters = new HashSet<>(Arrays.asList(input.substring(i, i+=4).split("")));
        }
        System.out.println(input.substring(i-4, i) + ": " + (i-1));
            
        // partOne(lines, i);
        // partTwo(lines, i);
    }

    public static void partOne(List<String> lines, int startLine) {
        
    }

    public static void partTwo(List<String> lines, int startLine) {
        
    }
}


