#include <iostream>
#include <map>

using namespace std;

double balance;

int main(){
    cout << "Welcome to Cpp Bank\n"
    << "Options:\n"
    << "- 1. Deposit\n"
    << "- 2. Withdraw\n"
    << "- 3. Balance\n"
    << "- 4. Exit" << endl;
    while (true){
        cout << "@CppBank> ";
        int input;
        cin >> input;

        double amount = 0;

        switch (input){
            case 1:{
                cout << "How much do you want to deposit? ";
                cin >> amount;
                balance += amount;
                cout << "Deposited: " << amount << endl;
                break;
            }
            case 2:{
                cout << "How much do you want to withdraw? ";
                cin >> amount;
                if (balance < amount){
                    cout << "Cannot withdraw more than your balance" << endl;
                    break;
                }
                balance -= amount;
                cout << "Withdrew: " << amount << endl;
                break;
            }
            case 3:{
                cout << "Balance: " << balance << endl;
                break;
            }
            case 4:{
                return 0;
            }
            default:{
                cout << "Invalid option" << endl;
                break;
            }
        }
    }
    return 0;
}