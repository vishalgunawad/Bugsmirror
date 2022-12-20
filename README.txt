
This is a Backend API made with Golang and Mongodb (used for User data storing)

to run this project go to the main directory insie the terminal and type "go run main.go"
it will start the server 


now use TunderClient OR Postman or any api calling software for  calling API

i am using TunderClient below is the examples for calling api in TunderClient


// CREATE USER

URL  = localhost:8080/api/v1/Users
MEATHOD = POST
Http Headers = "Content-Type: application/json" 
Body ={
    "id": 1,
    "First_Name": "Leanne Graham",
    "Last_Name": "Bret",
    "Date_Of_Birth" : "01/01/2000",
    "Email": "Sincere@april.biz",
    "Phone_number" : 123456789
  }






//GET ALL USER

URL  = localhost:8080/api/v1/Users
MEATHOD = GET
Http Headers = "Content-Type: application/json" 





//EDIT USER     {it will edit the user which have the id matching in the database}

URL  = localhost:8080/api/v1/Users
MEATHOD = PUT
Http Headers = "Content-Type: application/json" 
Body ={
    "id": 1,
    "First_Name": "Leanne Graham",
    "Last_Name": "Bret",
    "Date_Of_Birth" : "01/01/2000",
    "Email": "Sincere@april.biz",
    "Phone_number" : 123456789
  }





// DELETE USER

URL  = localhost:8080/api/v1/Users
MEATHOD = DELETE
Http Headers = "Content-Type: application/json" 
Body = { "id": 1}
