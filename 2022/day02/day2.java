import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.util.Map;
import java.util.Map.Entry;

class day2 {
    public static final int MAX_ITERATIONS = 3000;
    public static final Map<String, String> decision = Map.of(
        "Y", "A", 
        "Z", "B", 
        "X", "C"
    );
    public static final Map<String, Integer> scores = Map.of(
        "A",1,
        "B",2,
        "C",3,
        "X",1,  
        "Y",2,  
        "Z",3
    );
    public static void main(String[] args) {
        try (BufferedReader reader = new BufferedReader(new FileReader(new File(args[0])))) {
            String line = reader.readLine();
            String[] match;
            String opponent, choice;
            int total = 0, iterations = 0;
            boolean ends = false;
            // X - LOSE
            // Y - DRAW
            // Z - WIN 
            while (line != null && iterations < MAX_ITERATIONS) {
                ends = (iterations < 500 || iterations > 2000);
                match = line.split(" ");
                opponent = match[0];
                choice = match[1];
                switch(choice) {
                    case "X" -> { 
                        if (ends)
                            System.out.println("LOSE against " + opponent + " | " + total + " + " + ((scores.get(opponent) + 2) % 3 + 1));
                        total += (scores.get(opponent) + 2) % 3 + 1;
                    }        // LOSE
                        
                    case "Y" -> { 
                        if (ends)
                            System.out.println("TIE against " + opponent + " | " + total + " + " + scores.get(opponent) + " + " + 3);
                        total += scores.get(opponent) + 3;
                    }                  // TIE
                                
                    case "Z" -> { 
                        if (ends)
                            System.out.println("WIN against " + opponent + " | " + total + " + " + ((scores.get(opponent) + 1) % 3 + 1) + " + " + 6);
                        total += (scores.get(opponent) + 1) % 3 + 1 + 6;
                    }    // WIN
                                
                    }
    
                line = reader.readLine();
                iterations++;
            }
            System.out.println("Total: " + total);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}