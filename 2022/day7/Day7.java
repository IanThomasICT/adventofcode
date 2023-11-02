import aoc22.Utils;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.InvalidPathException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.*;

public class Day7 {
    public static class Dir {
        String name = "";
        Dir parentDir;
        List<Dir> children = new ArrayList<>();
        Integer sumOfFiles = 0;
        Integer totalSize = 0;

        public Dir(String name) {
            this.name = name;
        }

        public Dir(String name, Dir parentDir) {
            this.name = name;
            this.parentDir = parentDir;
        }

        public void addDir(Dir dir) {
            this.children.add(dir);
            dir.parentDir = this;
        }

        public void addFile(Integer fileSize) {
            sumOfFiles += fileSize;
        }

        public Integer getSumOfFiles() {
            return this.sumOfFiles;
        }

        public void clearSize() {
            sumOfFiles = 0;
        }

        @Override
        public String toString() {
            return "Dir{" +
                    "Size=" + getSumOfFiles() +
                    ", Dirs=" + children +
                    ", parent='" + parentDir + '\'' +
                    '}';
        }

        public boolean childExists(String childName) {
            return children.stream()
                    .anyMatch(c -> c.name.equals(childName));
        }


        public Dir getChild(String childName) {
            return children.stream()
                    .filter(c -> c.name.equals(childName))
                    .findFirst().orElse(null);
        }

        public Integer getTotalSize() {
            return totalSize;
        }
    }

    public static class FileSystem {
        Dir root;

        public Dir getRoot() {
            return root;
        }

        public FileSystem(Dir root) {
            this.root = root;
        }

        public List<Dir> getAllDirectoriesAndTotalSize() {
            List<Dir> directories = new ArrayList<>();
            Dir root = getRoot();
            Integer totalSize = addChildDirectoriesAndSetTotalSize(root, directories);
            root.totalSize = totalSize;
            return directories;
        }

        public Integer addChildDirectoriesAndSetTotalSize(Dir node, List<Dir> list) {
            Integer totalSize = node.getSumOfFiles();
            if (!node.children.isEmpty()) {
                for (Dir child: node.children) {
                    totalSize += addChildDirectoriesAndSetTotalSize(child, list);
                }
            }
            list.add(node);
            node.totalSize = totalSize;
            return totalSize;
        }
    }

    private static FileSystem mapFileSystem(List<String> instructions) {
        FileSystem fileSystem = new FileSystem(new Dir("/", new Dir("")));
        Dir currentDir = fileSystem.getRoot();
        for (String line : instructions) {
            try {
                String[] parts = line.split(" ");
                switch (parts[0]) {
                    case "$" -> {
                        switch (parts[1]) {
                            case "cd" -> {
                                switch (parts[2]) {
                                    case ".." -> currentDir = currentDir.parentDir;
                                    case "/" -> currentDir = fileSystem.getRoot();
                                    default ->
                                            currentDir = currentDir.getChild(parts[2]); // Assuming changing to an existing directory
                                }
                            }
                            case "ls" -> currentDir.clearSize();
                            default -> throw new RuntimeException("Unexpected command in line: " + parts[1]);
                        }
                    }
                    case "dir" -> {
                        if (!currentDir.childExists(parts[1])) {
                            currentDir.addDir(new Dir(parts[1]));
                        }
                    }
                    default -> currentDir.addFile(Integer.valueOf(parts[0]));
                }
            } catch (Exception e) {
                Utils.error("An error occurred on line: " + line);
                e.getStackTrace();
                throw e;
            }
        }
        return fileSystem;
    }


    public static void main(String[] args) throws IOException {
        String filePath = args[0];
        List<String> instructions = Utils.fileToStringList(filePath);
        partOne(instructions);
        partTwo(instructions);
    }

    public static void partOne(List<String> instructions) {
        instructions = instructions.subList(2, instructions.size()); // Skip root cd
        FileSystem fileSystem = mapFileSystem(instructions);
        List<Dir> directories = fileSystem.getAllDirectoriesAndTotalSize();
        System.out.println("Part 1: " + directories.stream().filter(dir -> dir.getTotalSize() < 100_000).mapToInt(Dir::getTotalSize).sum());

        // System.out.println(fileSystem.getSumOfTotalsUnder100k());
    }

    public static void partTwo(List<String> instructions) {
        instructions = instructions.subList(2, instructions.size()); // Skip root cd
        FileSystem fileSystem = mapFileSystem(instructions);
        List<Dir> directories =  fileSystem.getAllDirectoriesAndTotalSize();
        Integer usedSpace = fileSystem.getRoot().getTotalSize();
        Integer requiredSpaceForUpdate = 30_000_000 - (70_000_000 - usedSpace);
        System.out.println("Part 2: " + directories.stream().filter(d -> d.getTotalSize() > requiredSpaceForUpdate).mapToInt(Dir::getTotalSize).min().getAsInt());
    }

}


