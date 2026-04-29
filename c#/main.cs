//dotnet run c#/main.cs

double balance = 0;

Console.WriteLine("Welcome to C# bank\nOptions:\n- 1. Deposit\n- 2. Withdraw\n- 3. Balance\n- 4. Exit");

while (true)
{
    Console.Write("@C#Bank> ");
    string option = Console.ReadLine();

    switch (option)
    {
        case "1": // Deposit
            {
                Console.Write("Deposit amount: ");
                double amount = double.Parse(Console.ReadLine());
                balance += amount;
                Console.WriteLine("Deposited $" + amount);
                break;
            }
        case "2": // Withdraw
            {
                Console.Write("Withdraw amount: ");
                double amount = double.Parse(Console.ReadLine());
                if (amount > balance)
                {
                    Console.WriteLine("Cannot withdraw more than your balance");
                    break;
                }
                balance -= amount;
                Console.WriteLine("Withdrew $" + amount);
                break;
            }
        case "3": // Balance
            {
                Console.WriteLine("Balance: $" + balance);
                break;
            }
        case "4": // Exit
            {
                return;
            }
        default:
            {
                Console.WriteLine("Invalid input");
                break;
            }
    }
}