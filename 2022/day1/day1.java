import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.io.Reader;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.List;
import java.util.TreeMap;
import java.util.Map.Entry;
import java.util.stream.Stream;

class day1 {
    public static final String ANSI_RESET = "\u001B[0m";
    public static final String ANSI_RED = "\u001B[31m";
    public static final String ANSI_GREEN = "\u001B[32m";
    public static final String ANSI_YELLOW = "\u001B[33m";
    public static final int MAX_ITERATIONS = 2500;

    /**
     * Day 1 of Advent of Code 2022:
     *   - Problem 1: Read through a file calculating groups of integers "calories" delimited by blank lines   
     * 
     **/
    public static void main(String[] args) {
        String fileName = "";
        TreeMap<Integer, Integer> elves = new TreeMap<>();
        try {
            fileName = args[0];
        } catch (Exception e) {
            error("java day1.java [file] \nInvalid args: Program must include a file");
            System.exit(1);
        }
        
        if (!fileName.contains(".txt")) {
            error("java day1.java [file] \nInvalid file type: File must be a text file [*.txt]");
            System.exit(1);
        }
        
        try(BufferedReader reader = new BufferedReader(new FileReader(new File(fileName)));) {
            log("Calculating total calories...");
            String line = reader.readLine();
            int max = 0, count = 0, elfNum = 1, iterations = 0;
            while (line != null && iterations < MAX_ITERATIONS) {
                if (line.isBlank()) {
                    if (max < count) {
                        max = count;
                    }
                    if (count > 0) {
                        elves.put(count, elfNum++);
                    }
                    count = 0;
                } else {
                    count += Integer.parseInt(line);
                }
                line = reader.readLine();
                iterations++;
            }
            if (count != 0) {
                elves.put(count, elfNum);
            }
            List<Entry<Integer, Integer>> top3 = new ArrayList<>();
            for (int i = 0; i < 3; i++) {
                top3.add(elves.pollLastEntry());
            } 
            log("""
                    Top #3 Elves with the most calories:
                      - Elf #%d: %d
                      - Elf #%d: %d
                      - Elf #%d: %d
                    """.formatted(
                        top3.get(0).getValue(),top3.get(0).getKey(),
                        top3.get(1).getValue(),top3.get(1).getKey(),
                        top3.get(2).getValue(),top3.get(2).getKey()
                        ));
            log(String.format("The Top 3 Elves have a total of %d calories.", top3.stream().mapToInt(Entry::getKey).sum()));
            System.exit(0);
        } catch (IOException e) {
            error("Failed to load the file " + fileName);
            e.printStackTrace();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static void log(String str) {
        System.out.println(str);
    }

    public static void error(String str) {
        System.out.println(ANSI_RED + str + ANSI_RESET);
    }
}
