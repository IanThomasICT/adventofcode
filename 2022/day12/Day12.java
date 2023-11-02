//import aoc22.Utils;
//
//import java.util.*;
//import java.util.stream.Collectors;
//
//public class Day12 {
//
//    static String filePath = "input/day12-test.txt";
//    // static String filePath = "input/day12.txt";
//    public static void main(String[] args) throws InterruptedException {
//        List<String> lines = Utils.fileToStringList(filePath);
//        partOne(lines);
//        // partTwo(lines);
//    }
//
//    public static void partOne(List<String> lines) {
//        for (int i = 0; i < lines.size(); i++) {
//            for (int j = 0; j < lines.get(0).length(); j++) {
//
//            }
//        }
//    }
//
//
//
//
//
//
//
//
//
//
//    private static final int MAX_DISTANCE = 10000000;
//
//    public static class Node {
//        int[] coord;
//        int value;
//        Node parent;
//        Node child;
//        int distanceFromStart = MAX_DISTANCE;
//        Node(int[] coords, int val) {
//            coord = coords;
//            value = val;
//        }
//        void setDistance(int distance, Node parent) {
//            distanceFromStart = Math.min(distance, distanceFromStart);
//            this.parent = parent;
//        }
//
//        Node(Node n) {
//            coord = new int[]{n.coord[0], n.coord[1]};
//            value = n.value;
//            distanceFromStart = n.distanceFromStart;
//        }
//
//        int distanceFromStart() {
//            return distanceFromStart;
//        }
//
//        @Override
//        public String toString() {
//            return "Node{" +
//                    "coord=" + Arrays.toString(coord) +
//                    ", value=" + value +
//                    ", distance=" + distanceFromStart +
//                    '}';
//        }
//    }
//
//    public static class Grid {
//        Node[][] grid;
//        int length, height;
//        Node start, end;
//        PriorityQueue<Node> unvisited = new PriorityQueue<>(Comparator.comparing(Node::distanceFromStart));
//        PriorityQueue<Node> visited = new PriorityQueue<>(Comparator.comparing(Node::distanceFromStart));
//
//        public Grid(int height, int length) {
//            this.length = length;
//            this.height = height;
//            this.grid = new Node[height][length];
//        }
//
//        public Grid(int[][] grid) {
//            this.length = grid[0].length;
//            this.height = grid.length;
//            this.grid = new Node[height][length];
//            for (int i = 0; i < grid.length; i++) {
//                for (int j = 0; j < grid[0].length; j++) {
//                    Node node = new Node(new int[]{i,j}, grid[i][j]);
//                    this.grid[i][j] = node;
//                    unvisited.add(node);
//
//                    switch(grid[i][j]) {
//                        case S -> {
//                            start = node;
//                            node.distanceFromStart = 0;
//                            node.value = 0;
//                        }
//                        case E -> {
//                            end = node;
//                            node.distanceFromStart = 10000005;
//                            node.value = 27;
//                        }
//                        default -> {}
//                    }
//                }
//            }
//        }
//
//        void importGrid(int[][] grid) {
//            for (int i = 0; i < grid.length; i++) {
//                for (int j = 0; j < grid[0].length; j++) {
//                    Node node = new Node(new int[]{i,j}, grid[i][j]);
//                    this.grid[i][j] = node;
//                    unvisited.add(node);
//
//                    switch(grid[i][j]) {
//                        case S -> {
//                            start = node;
//                            node.distanceFromStart = 0;
//                            node.value = 0;
//                        }
//                        case E -> {
//                            end = node;
//                            node.distanceFromStart = 10000005;
//                            node.value = 27;
//                        }
//                        default -> {}
//                    }
//                }
//            }
//        }
//
//        int[][] toIntGrid() {
//            return Arrays.stream(grid)
//                    .map(nArr -> Arrays.stream(nArr)
//                            .map(n -> n.value)
//                            .mapToInt(Integer::intValue)
//                            .toArray())
//                    .toArray(int[][]::new);
//        }
//
//        void printGrid() {
//            int[][] nodeGrid = Arrays.stream(grid)
//                    .map(nArr -> Arrays.stream(nArr)
//                            .map(n -> n.value)
//                            .mapToInt(Integer::intValue)
//                            .toArray())
//                    .toArray(int[][]::new);
//            Utils.printGrid(nodeGrid);
//        }
//
//        public boolean isUnvisited(Node node) {
//            return unvisited.contains(node);
//        }
//
//        public void visit(Node node) {
//            unvisited.remove(node);
//            visited.add(node);
//        }
//
//        public Node getNextNode() {
//            return unvisited.poll();
//        }
//
//        public void refreshQueue(Node move) {
//            if (unvisited.contains(move)) {
//                unvisited.remove(move);
//                unvisited.add(move);
//            }
//        }
//
//        public LinkedList<String> getShortestPathTo(Node startNode) {
//            Node node = startNode;
//            LinkedList<String> pathFromEnd = new LinkedList<>();
//            while (node != null) {
//                pathFromEnd.push(Arrays.toString(node.coord));
//                node = node.parent;
//            }
//            return pathFromEnd;
//        }
//
//        public Grid copyGrid() {
//            return new Grid(toIntGrid());
//        }
//    }
//
//
//    final static int E = ('E' - 'a' + 1);
//    final static int S = ('S' - 'a' + 1);
//    public static void partOne(List<String> lines) throws InterruptedException {
//        int[][] grid = convertToIntGrid(lines);
//        Grid graph = new Grid(lines.size(), lines.get(0).length());
//        graph.importGrid(grid);
//
//        traverseGraph(graph);
//        System.out.println("Part 1: shortest path is " + graph.end.distanceFromStart + " moves");
//    }
//
//    public static void partTwo(List<String> lines) throws InterruptedException {
//        int[][] grid = convertToIntGrid(lines);
//        Grid graph = new Grid(lines.size(), lines.get(0).length());
//        graph.importGrid(grid);
//
//        TreeMap<Integer, List<String>> pathsReachingEnd = new TreeMap<>();
//        // Get nodes that have possible ascending moves
//        LinkedList<Node> groundLevelNodes = Arrays.stream(graph.grid)
//                .flatMap(Arrays::stream)
//                .filter(n -> n.value == 1)
//                .filter(n -> getPossibleMoves(n, graph).stream().anyMatch(m -> m.value == 2))
//                .collect(Collectors.toCollection(LinkedList::new));
//
//        int min = MAX_DISTANCE;
//        while (!groundLevelNodes.isEmpty()) {
//            int[][] tmpGrid = convertToIntGrid(lines);
//            Grid tmpGraph = new Grid(lines.size(), lines.get(0).length());
//            tmpGraph.importGrid(tmpGrid);
//
//            Node node = groundLevelNodes.poll();
//            min = Math.min(min, traverseGraphFromNode(node, tmpGraph));
//        }
//
//        System.out.println("Part 2: shortest path is " + min + " moves");
//        // Utils.drawPath(graph.grid, pathsReachingEnd.firstEntry().getValue());
//
//        // Answers:
//        // Too high: 510, 509
//        // Correct: 508
//    }
//
//    public static int totalPossibleMovesToEnd(Node node, Node end) {
//        return Math.abs(end.coord[0] - node.coord[0]) + Math.abs(end.coord[1] - node.coord[1]);
//    }
//
//    private static int[][] convertToIntGrid(List<String> lines) {
//        int height = lines.size(), width = lines.get(0).length();
//        int[][] grid = new int[height][width];
//        for (int i = 0; i < height; i++) {
//            String row = lines.get(i);
//            for (int j = 0; j < width; j++) {
//                int value = row.charAt(j) - 'a' + 1;
//                grid[i][j] = value;
//            }
//        }
//        return grid;
//    }
//
//    public static PriorityQueue<Node> getPossibleMoves(Node node, Grid graph) {
//        int i = node.coord[0], j = node.coord[1], value = graph.grid[i][j].value;
//        int iMax = graph.height - 1, jMax = graph.length - 1;
//        PriorityQueue<Node> cardinalDirections = new PriorityQueue<>(Comparator.comparingDouble(n -> getDistanceScore(n, graph.end)));
//
//        if (i != 0      && (graph.grid[i-1][j].value - value <= 1 || graph.grid[i-1][j].value == E) && graph.isUnvisited(graph.grid[i-1][j]))  cardinalDirections.add(graph.grid[i-1][j]);
//        if (i != iMax   && (graph.grid[i+1][j].value - value <= 1 || graph.grid[i+1][j].value == E) && graph.isUnvisited(graph.grid[i+1][j]))  cardinalDirections.add(graph.grid[i+1][j]);
//        if (j != 0      && (graph.grid[i][j-1].value - value <= 1 || graph.grid[i][j-1].value == E) && graph.isUnvisited(graph.grid[i][j-1]))  cardinalDirections.add(graph.grid[i][j-1]);
//        if (j != jMax   && (graph.grid[i][j+1].value - value <= 1 || graph.grid[i][j+1].value == E) && graph.isUnvisited(graph.grid[i][j+1]))  cardinalDirections.add(graph.grid[i][j+1]);
//
//        // System.out.printf(node + "| Available moves: " + cardinalDirections.stream().map(n -> Arrays.toString(n.coord)).toList() + "\n");
//        return cardinalDirections;
//    }
//
//    public static PriorityQueue<Node> getPossibleMoves(Node node, Grid graph, PriorityQueue<Node> unvisited) {
//        int i = node.coord[0], j = node.coord[1], value = graph.grid[i][j].value;
//        int iMax = graph.height - 1, jMax = graph.length - 1;
//        PriorityQueue<Node> cardinalDirections = new PriorityQueue<>(Comparator.comparingDouble(n -> getDistanceScore(n, graph.end)));
//
//        if (i != 0      && graph.grid[i-1][j].value - value <= 1 && (unvisited.contains(graph.grid[i-1][j]) || node.distanceFromStart + 1 <= graph.grid[i-1][j].distanceFromStart))  cardinalDirections.add(graph.grid[i-1][j]);
//        if (i != iMax   && graph.grid[i+1][j].value - value <= 1 && (unvisited.contains(graph.grid[i+1][j]) || node.distanceFromStart + 1 <= graph.grid[i+1][j].distanceFromStart))  cardinalDirections.add(graph.grid[i+1][j]);
//        if (j != 0      && graph.grid[i][j-1].value - value <= 1 && (unvisited.contains(graph.grid[i][j-1]) || node.distanceFromStart + 1 <= graph.grid[i][j-1].distanceFromStart))  cardinalDirections.add(graph.grid[i][j-1]);
//        if (j != jMax   && graph.grid[i][j+1].value - value <= 1 && (unvisited.contains(graph.grid[i][j+1]) || node.distanceFromStart + 1 <= graph.grid[i][j+1].distanceFromStart))  cardinalDirections.add(graph.grid[i][j+1]);
//
//        System.out.printf(node + "| Available moves: " + cardinalDirections.stream().map(n -> Arrays.toString(n.coord)).toList() + "\n");
//        return cardinalDirections;
//    }
//
//    private static Double getDistanceScore(Node node1, Node node2) {
//        return Math.sqrt(Math.pow(node1.coord[1] - node2.coord[1],2) + Math.pow(node1.coord[0] - node2.coord[0],2));
//    }
//
//
//    public static void traverseGraph(Grid graph) {
//        Node node = graph.start;
//        while (graph.end.parent == null) {
//            PriorityQueue<Node> possibleMoves = getPossibleMoves(node, graph);
//            for (Node move: possibleMoves) {
//                move.setDistance(node.distanceFromStart + 1, node);
//                node.child = move;
//                graph.refreshQueue(move);
//            }
//            node = graph.getNextNode();
//        }
//    }
//
//    public static Integer traverseGraphFromNode(Node start, Grid graph) {
//        Node node = start;
//        node.distanceFromStart = 0;
//        graph.visit(start);
//        while (graph.end.parent == null) {
//            PriorityQueue<Node> possibleMoves = getPossibleMoves(node, graph);
//            for (Node move: possibleMoves) {
//                move.setDistance(node.distanceFromStart + 1, node);
//                node.child = move;
//                graph.refreshQueue(move);
//            }
//
//            if (graph.unvisited.peek() == null || graph.unvisited.peek().distanceFromStart == MAX_DISTANCE) {
//                return MAX_DISTANCE;
//            }
//
//            node = graph.getNextNode();
//        }
//
//        // Utils.drawPath(graph.grid,graph.getShortestPathTo(node));
//
//        return graph.end.distanceFromStart;
//
//    }
//
//}
//
