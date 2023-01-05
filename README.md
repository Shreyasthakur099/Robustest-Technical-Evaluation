# Robustest-Technical-Evaluation
This repository contains Golang based task as a part of Robustest Technical Evaluation round 2

-> The Details about task is in the pdf file `Backend Assignemnt.pdf` 
-> `account.json` file contains sample data to test the application.

<br><br><br>
## How to run the application

To start the application , open the directory in terminal and enter command `go run server.go`.

It will start the server at `PORT : 8080`

-> Now , Use Postman(or any other tool) to run APIs

-> API endpoints along with their Inputs/Outputs are shown below :

<br><br><br><br>
### To check Balance of a User Account as Admin , Enter `http://localhost:8080/admin/checkBalance/` 

* INPUT : 
![image](https://user-images.githubusercontent.com/71756168/210740406-dc375dce-ff09-4c69-8e48-f4f5c38e67f2.png)


OUTPUT : 

![image](https://user-images.githubusercontent.com/71756168/210740632-bbb62195-2c3f-450e-9563-5c5d17d3736a.png)



<br><br><br><br>
### To check Balance of a User Account as a User , Enter `http://localhost:8080/user/checkBalance/` 

* INPUT :
![image](https://user-images.githubusercontent.com/71756168/210741042-2181b2c5-e01b-4b65-9c62-99a103e58e4c.png)

OUTPUT : 

![image](https://user-images.githubusercontent.com/71756168/210741118-4a4438a4-ba91-40b1-a1c4-ab9fd118af23.png)




<br><br><br><br>
### To Send Money as a User , Enter `http://localhost:8080/user/sendMoney/` :


[NOTE : In real application scenario , there will be no need to enter `UserId` manually as it will be automatically retrieved after Login(Authentication)]

* INPUT : 

![image](https://user-images.githubusercontent.com/71756168/210741969-4e5f9784-925a-4268-9404-6d50e0cf1eb9.png)


OUTPUT : 

![image](https://user-images.githubusercontent.com/71756168/210742095-25b02936-b4c5-46ba-bc98-f5fedab0208a.png)




<br><br><br><br>
### To Send Money as Admin , Enter `http://localhost:8080/admin/sendMoney/` :

* INPUT : 

![image](https://user-images.githubusercontent.com/71756168/210742480-5f9c80db-3346-4eb4-ab43-a2f547a5b4f6.png)

OUTPUT : 

![image](https://user-images.githubusercontent.com/71756168/210742561-f9c323d2-a39d-45bf-a5a9-fd94ff6cb699.png)

If the Sender Doesnt have Sufficient Balance , then Output will be :

![image](https://user-images.githubusercontent.com/71756168/210742751-9b9bc64c-9e81-48ee-b429-fe13728b2a4b.png)



<br><br><br><br>
### Check Transactions as a User , Enter `http://localhost:8080/user/transactions/`

[NOTE : In real application scenario , there will be no need to enter `UserId` manually as it will be automatically retrieved after Login(Authentication)]

* INPUT : 

![image](https://user-images.githubusercontent.com/71756168/210743173-bbc9346f-81f4-4a94-9283-0a56a215b83a.png)


   OUTPUT : 
   
   ![image](https://user-images.githubusercontent.com/71756168/210744173-f4124d78-cda8-4c98-9071-b078ce8fe75e.png)


<br><br><br><br>
### Check Transactions as a User , Enter `http://localhost:8080/admin/transactions/`

* INTPUT : 
 
 ![image](https://user-images.githubusercontent.com/71756168/210744611-39839c2e-b4be-4a1b-95c5-c7e4e6dc0ad6.png)

  OUTPUT : 
  
  ![image](https://user-images.githubusercontent.com/71756168/210744743-b6e0c413-ffa6-42e2-80a6-6c9183934728.png)

