import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

public class Day3 { 
    public static void main(String[] args) {
        // Iterate through list
        try {
            String inputFile = args[0];
            partOne(inputFile);
            
            // Part 2
            partTwo(inputFile);
        } catch (Exception e) {
            System.out.println("An error occurred");
        }
    }

    private static void partTwo(String inputFile) throws IOException {
        List<String> rucksacks = Files.lines(Path.of(inputFile)).collect(Collectors.toList());
        int total = 0;
        System.out.println("Part 2: Checking rucksacks in groups of threes for common badges");
        for (int i = 0; i < rucksacks.size() - 2; i += 3) {
            String[] containers = {rucksacks.get(i), rucksacks.get(i+1), rucksacks.get(i+2) };
            Set<String> firstRucksack = new HashSet<>(Arrays.asList(containers[0].split(""))); 
            Set<String> secondRucksack = new HashSet<>(Arrays.asList(containers[1].split(""))); 
            Set<String> thirdRucksack = new HashSet<>(Arrays.asList(containers[2].split(""))); 
            
            for (String item : firstRucksack) {
                if (secondRucksack.contains(item) && thirdRucksack.contains(item)) {
                    int priority = item.compareTo("a") < 0 ? item.codePointAt(0) - 38 : item.codePointAt(0) - 96;
                    total += priority;
                }
            }
        }
        System.out.println(total);
            
}

    private static void partOne(String inputFile) throws IOException {
            List<String> rucksacks = Files.lines(Path.of(inputFile)).collect(Collectors.toList());
            int total = 0;
            System.out.println("Part 1: Checking rucksacks for common items");
            for (String rucksack : rucksacks) {
                String[] containers = {rucksack.substring(0, rucksack.length()/2), rucksack.substring(rucksack.length()/2)};
                Set<String> firstContainerContents = new HashSet<>(Arrays.asList(containers[0].split(""))); 
                Set<String> secondContainerContents = new HashSet<>(Arrays.asList(containers[1].split(""))); 
                
                for (String item : firstContainerContents) 
                    if (secondContainerContents.contains(item)) {
                        int priority = item.compareTo("a") < 0 ? item.codePointAt(0) - 38 : item.codePointAt(0) - 96;
                        total += priority;
                    }
                }
            System.out.println(total);
    }
}