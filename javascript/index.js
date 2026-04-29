const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

let balance = 0;

console.log("Welcome to JsBank\nOptions:\n- 1. Deposit\n- 2. Withdraw\n- 3. Balance\n- 4. Quit");

function ask() {
    rl.question("@JsBank> ", (input) => {
        switch (input) {
            case "1":
                rl.question("Deposit Amount: $", (amount) => {
                    balance += parseFloat(amount);
                    console.log("Deposited $" + amount);
                    ask();
                });
                break;
            case "2":
                rl.question("Withdraw Amount: $", (amount) => {
                    amount = parseFloat(amount);
                    if (amount > balance) {
                        console.log("Cannot withdraw more than your balance");
                        ask();
                        return;
                    }
                    balance -= amount;
                    console.log("Withdrew $" + amount);
                    ask();
                });
                break;
            case "3":
                console.log("Balance: $" + balance);
                ask();
                break;
            case "4":
                rl.close();
                return;
            default:
                console.log("Invalid input");
                ask();
                break;
        }
    });
}

ask();