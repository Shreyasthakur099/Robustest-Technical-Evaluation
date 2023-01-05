package main

import (
	"encoding/json"
	"fmt"
	"io"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
    HOST = "localhost"
    PORT = "8080"
    TYPE = "tcp"
)


type Transaction struct{
	RecieverID	int `json :"ReceiverID"`	
	SenderID	int `json :"SenderID "`	
	UserID		int `json :"UserID , omitempty "`	
	Amount		int `json :"Amount"`	
}

type Account struct{
	Balance 		int	  			`json:"Balance"`
	UserID 			int	  			`json:"UserID"`
	Transactions 	[]Transaction 	`json:"Transactions"`
}

type allAccounts struct{
	AccountData []Account	`json:"accountData"`
}

func balanceCheckByAdmin(writer http.ResponseWriter, request *http.Request){
	// assuming Admin is authenticated and logged in
	
	writer.Header().Set("Content-Type", "application/json") 
	if request.Method == "GET" { 
	writer.WriteHeader(http.StatusOK) 
	}
	var inputData Account
	
	err := json.NewDecoder(request.Body).Decode(&inputData) 
	if err != nil { 
		log.Fatalln("There was an error decoding the request body" , err) 
	}else{
		fmt.Println("Checking balance for user id : " , inputData.UserID)
	}

	jsonFile , err := os.Open(("account.json"))
	if err != nil{
		log.Fatalln(err)
	}

	// parsing json file
	byteValue, _ := io.ReadAll(jsonFile)
	var allaccounts allAccounts
	
	err1 := json.Unmarshal([]byte(byteValue), &allaccounts)
	if err1 != nil{
		fmt.Println("Error in Parsing : " , err1)
	}

	
	// searching for the user
	for i:=0 ; i < len(allaccounts.AccountData) ; i++ {
		if allaccounts.AccountData[i].UserID == inputData.UserID{
			fmt.Println("Account Balance : " , allaccounts.AccountData[i].Balance)
			json.NewEncoder(writer).Encode(allaccounts.AccountData[i].Balance)
			return
		}
	}
}


func adminTransactions(writer http.ResponseWriter, request *http.Request){
	// assuming Admin is authenticated and logged in
	
	writer.Header().Set("Content-Type", "application/json") 
	if request.Method == "GET" { 
	writer.WriteHeader(http.StatusOK) 
	}
	var inputData Account
	
	err := json.NewDecoder(request.Body).Decode(&inputData) 
	if err != nil { 
		log.Fatalln("There was an error decoding the request body" , err) 
	}else{
		fmt.Println("Checking Transactions for user id : " , inputData.UserID)
	}

	jsonFile , err := os.Open(("account.json"))
	if err != nil{
		log.Fatalln(err)
	}

	// parsing json file
	byteValue, _ := io.ReadAll(jsonFile)
	var allaccounts allAccounts
	
	err1 := json.Unmarshal([]byte(byteValue), &allaccounts)
	if err1 != nil{
		fmt.Println("Error in Parsing : " , err1)
	}

	
	// searching for the user
	for i:=0 ; i < len(allaccounts.AccountData) ; i++ {
		if allaccounts.AccountData[i].UserID == inputData.UserID{
			fmt.Println("Account Transactions : " , allaccounts.AccountData[i].Transactions)
			json.NewEncoder(writer).Encode(allaccounts.AccountData[i].Transactions)
			return
		}
	}
}


func adminSendMoney(writer http.ResponseWriter, request *http.Request){
	// assuming User is authenticated and logged in
		
	writer.Header().Set("Content-Type", "application/json") 
	if request.Method == "GET" { 
	writer.WriteHeader(http.StatusOK) 
	}
	
	var inputData Transaction 

	err := json.NewDecoder(request.Body).Decode(&inputData) 
	if err != nil { 
		log.Fatalln("There was an error decoding the request body" , err) 
	}

	jsonFile , err := os.Open(("account.json"))
	if err != nil{
		log.Fatalln(err)
	}

	// parsing json file
	byteValue, _ := io.ReadAll(jsonFile)
	var allaccounts allAccounts

	err1 := json.Unmarshal([]byte(byteValue), &allaccounts)
	if err1 != nil{
		fmt.Println("Error in Parsing : " , err1)
	}
	
	// saving transaction
	var newTransaction Transaction
	newTransaction.Amount = inputData.Amount

	var senderIndex , recieverIndex int
	// searching for the sender and reciever account
	for i:=0 ; i < len(allaccounts.AccountData) ; i++ {
		
		
		if allaccounts.AccountData[i].UserID == inputData.RecieverID{
			
			newTransaction.RecieverID = inputData.RecieverID
			recieverIndex = i;
		}
		
		if allaccounts.AccountData[i].UserID == inputData.SenderID{
			
			newTransaction.SenderID = allaccounts.AccountData[i].UserID
			senderIndex = i

			// checking for sufficient balance in Sender account
			if allaccounts.AccountData[i].Balance < inputData.Amount{
				fmt.Println("Insufficient Balance!")
				json.NewEncoder(writer).Encode("Insufficient Balance")
				return
			}
			
		}

		if  (newTransaction.SenderID != 0) && (newTransaction.RecieverID != 0){
			break	
		}

	}

	// sender and reciever object found ; updating data...
	allaccounts.AccountData[senderIndex].Transactions = append(allaccounts.AccountData[senderIndex].Transactions, newTransaction)
	allaccounts.AccountData[recieverIndex].Transactions = append(allaccounts.AccountData[recieverIndex].Transactions, newTransaction)
			
	// updating sender/reciever account balance
	allaccounts.AccountData[senderIndex].Balance = (allaccounts.AccountData[senderIndex].Balance - newTransaction.Amount)  
	allaccounts.AccountData[recieverIndex].Balance = (allaccounts.AccountData[recieverIndex].Balance + newTransaction.Amount)  
	
	contents, err := json.Marshal(allaccounts)
	if err != nil {
		fmt.Println(err)
	}
	
	error := os.WriteFile("account.json" , contents ,0644)
	if error != nil {

		fmt.Println("Error writing to file : " , err)

	}else{
	
		fmt.Println(" Transaction successfull for amount : " , newTransaction.Amount)
		json.NewEncoder(writer).Encode(" Transaction successfull !")
	}
	
	return
}


func userSendMoney(writer http.ResponseWriter, request *http.Request){

		// assuming User is authenticated and logged in
		
		writer.Header().Set("Content-Type", "application/json") 
		if request.Method == "GET" { 
		writer.WriteHeader(http.StatusOK) 
		}
		
		var inputData Transaction 
	
		err := json.NewDecoder(request.Body).Decode(&inputData) 
		if err != nil { 
			log.Fatalln("There was an error decoding the request body" , err) 
		}
	
		jsonFile , err := os.Open(("account.json"))
		if err != nil{
			log.Fatalln(err)
		}
	
		// parsing json file
		byteValue, _ := io.ReadAll(jsonFile)
		var allaccounts allAccounts
	
		err1 := json.Unmarshal([]byte(byteValue), &allaccounts)
		if err1 != nil{
			fmt.Println("Error in Parsing : " , err1)
		}
		
		// making new transaction
		var newTransaction Transaction
		newTransaction.Amount = inputData.Amount
		newTransaction.RecieverID = inputData.RecieverID
		newTransaction.SenderID = inputData.UserID
	
		var senderIndex , recieverIndex int
		// searching for the sender and reciever account
		for i:=0 ; i < len(allaccounts.AccountData) ; i++ {
			
			
			if allaccounts.AccountData[i].UserID == inputData.RecieverID{
				
				recieverIndex = i;
			}
			
			if allaccounts.AccountData[i].UserID == inputData.SenderID{
				
				senderIndex = i

				// checking for sufficient balance in Sender account
				if allaccounts.AccountData[i].Balance < inputData.Amount{
					fmt.Println("Insufficient Balance!")
					json.NewEncoder(writer).Encode("Insufficient Balance")
					return
				}
				
			}
	
			if  (newTransaction.SenderID != 0) && (newTransaction.RecieverID != 0){
				break	
			}
	
		}
	
		// sender and reciever object found ; updating data...
		allaccounts.AccountData[senderIndex].Transactions = append(allaccounts.AccountData[senderIndex].Transactions, newTransaction)
		allaccounts.AccountData[recieverIndex].Transactions = append(allaccounts.AccountData[recieverIndex].Transactions, newTransaction)
				
		// updating sender/reciever account balance
		allaccounts.AccountData[senderIndex].Balance = (allaccounts.AccountData[senderIndex].Balance - newTransaction.Amount)  
		allaccounts.AccountData[recieverIndex].Balance = (allaccounts.AccountData[recieverIndex].Balance + newTransaction.Amount)  
		
		contents, err := json.Marshal(allaccounts)
		if err != nil {
			fmt.Println(err)
		}
		
		error := os.WriteFile("account.json" , contents ,0644)
		if error != nil {
	
			fmt.Println("Error writing to file : " , err)
	
		}else{
		
			fmt.Println(" Transaction successfull for amount : " , newTransaction.Amount)
			json.NewEncoder(writer).Encode(" Transaction successfull !")
		}
		
		return	
}


func checkBalanceByUser(writer http.ResponseWriter, request *http.Request){
	// assuming User is authenticated and logged in
	
	writer.Header().Set("Content-Type", "application/json") 
	if request.Method == "GET" { 
	writer.WriteHeader(http.StatusOK) 
	}
	var inputData Account
	
	err := json.NewDecoder(request.Body).Decode(&inputData) 
	if err != nil { 
		log.Fatalln("There was an error decoding the request body" , err) 
	}else{
		fmt.Println("Checking balance for user id : " , inputData.UserID)
	}

	jsonFile , err := os.Open(("account.json"))
	if err != nil{
		log.Fatalln(err)
	}

	// parsing json file
	byteValue, _ := io.ReadAll(jsonFile)
	var allaccounts allAccounts
	
	err1 := json.Unmarshal([]byte(byteValue), &allaccounts)
	if err1 != nil{
		fmt.Println("Error in Parsing : " , err1)
	}

	
	// searching for the user
	for i:=0 ; i < len(allaccounts.AccountData) ; i++ {
		if allaccounts.AccountData[i].UserID == inputData.UserID{
			fmt.Println("Account Balance : " , allaccounts.AccountData[i].Balance)
			json.NewEncoder(writer).Encode(allaccounts.AccountData[i].Balance)
			return
		}
	}
}


func userTransactions(writer http.ResponseWriter, request *http.Request){
	// assuming User is authenticated and logged in
	
	writer.Header().Set("Content-Type", "application/json") 
	if request.Method == "GET" { 
	writer.WriteHeader(http.StatusOK) 
	}
	var inputData Account
	
	err := json.NewDecoder(request.Body).Decode(&inputData) 
	if err != nil { 
		log.Fatalln("There was an error decoding the request body" , err) 
	}else{
		fmt.Println("Checking balance for user id : " , inputData.UserID)
	}

	jsonFile , err := os.Open(("account.json"))
	if err != nil{
		log.Fatalln(err)
	}

	// parsing json file
	byteValue, _ := io.ReadAll(jsonFile)
	var allaccounts allAccounts
	
	err1 := json.Unmarshal([]byte(byteValue), &allaccounts)
	if err1 != nil{
		fmt.Println("Error in Parsing : " , err1)
	}

	
	// searching for the user
	for i:=0 ; i < len(allaccounts.AccountData) ; i++ {
		if allaccounts.AccountData[i].UserID == inputData.UserID{
			fmt.Println("Account Transactions : " , allaccounts.AccountData[i].Transactions)
			json.NewEncoder(writer).Encode(allaccounts.AccountData[i].Transactions)
			return
		}
	}
}

func main(){

	http.HandleFunc("/admin/checkBalance/" , balanceCheckByAdmin )
	http.HandleFunc("/admin/transactions/" , adminTransactions )
	http.HandleFunc("/admin/sendMoney/" , adminSendMoney )
	http.HandleFunc("/user/sendMoney/" , userSendMoney )
	http.HandleFunc("/user/checkBalance/" , checkBalanceByUser )
	http.HandleFunc("/user/transactions/" , userTransactions )

	err := http.ListenAndServe(":8080" , nil)
	if err != nil{
		log.Fatalln("Error in server : " , err)
	}
}

