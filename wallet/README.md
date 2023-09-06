## Example of wallet module for blockchain

///////Run: 
~ ```go run wallet/cmd/main.go createblockchain -address "name"``` to create a blockchain with a specified address;
~ ```go run wallet/cmd/main.go printchain``` to print the chain;
~ ```go run wallet/cmd/main.go getbalance -address "name"``` to get actual balance of coins in the specified address;
~ ```go run wallet/cmd/main.go send -from "name" -to "name" -amount 123``` to send the specified amount of coins from specified address to specified address;
~ ```go run wallet/cmd/main.go createwallet``` to create new Wallet;
~ ```go run wallet/cmd/main.go listaddresses``` to list existing wallets;