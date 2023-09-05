## Example of coinbbase-like transaction with CLI interface

Run: 
~ ```go run coinbase_transaction/cmd/main.go createblockchain -address "name"``` to create a blockchain with a specified address;
~ ```go run coinbase_transaction/cmd/main.go printchain``` to print the chain;
~ ```go run coinbase_transaction/cmd/main.go getbalance -address "name"``` to get actual balance of coins in the specified address;
~ ```go run coinbase_transaction/cmd/main.go send -from "name" -to "name" -amount 123``` to send the specified amount of coins from specified address to specified address.