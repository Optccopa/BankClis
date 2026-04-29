import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.println("Welcome to Java bank\nOptions:\n- 1. Deposit\n- 2. Withdraw\n- 3. Balance\n- 4. Exit");
        double balance = 0;
        while (true) {
            System.out.print("Enter input: ");
            String input = scanner.nextLine();
            switch (input) {
                case "1": { // deposit
                    System.out.print("Deposit amount: ");
                    double amount = Double.parseDouble(scanner.nextLine());
                    balance += amount;
                    System.out.println("Deposited: $" + amount);
                    break;
                }
                case "2": {
                    System.out.print("Withdraw amount: ");
                    double amount = Double.parseDouble(scanner.nextLine());
                    if (amount > balance) {
                        System.out.println("Cannot withdraw more than your balance.");
                        break;
                    }
                    balance -= amount;
                    System.out.println("Withdrew: $" + amount);
                    break;
                }
                case "3": {
                    System.out.println("Balance: $" + balance);
                    break;
                }
                case "4": {
                    scanner.close();
                    System.exit(0);
                }
                default: {
                    System.out.println("Invalid input");
                    break;
                }
            }
        }
    }
}