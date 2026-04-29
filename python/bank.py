print("""Weclome to PyBank
Options:
- 1. Deposit
- 2. Withdraw
- 3. Balance
- 4. Quit""")

while True:
    userInput = input("@PyBank> ")
    if userInput == "1":
        amount = int(input("Deposit amount: "))
        balance += amount
        print(f"Deposited ${amount}")
        continue
    elif userInput == "2":
        amount = int(input("Deposit amount: "))
        if amount > balance:
            print(f"Amount cannot be over balance, Balance: ${balance}")
            continue
        balance -= amount
        print(f"Withdrew ${amount}")
        continue
    elif userInput == "3":
        print(f"Balance: ${balance}")
        continue
    elif userInput == "4":
        break
    else:
        print("Invalid input")
        continue