Youtube Showcash simple : https://www.youtube.com/watch?v=cgZAhC8mWos

# command !help

### **OWNER**
```
**Command: !buy <item> <quantity>**
- Buy items from the store.
	Example: !buy cid 5

**Command: !addbal <id client> <value>**
- Add a balance to a client's inventory.
	Example: !addbal 1234567890 10

**Command: !usebal <id client> <value>**
- Use a balance from a client's inventory.
	Example: !usebal 1234567890 5

**Command: !register**
- register your account discord
	Example: !register

**Command: !bal**
- check balance you acocunt
	Example: !bal

**Note:**
- Make sure to use the correct format for the commands.
- Replace <item> and <quantity> with the specific item and quantity you want to buy.
- Replace <id client> with the client ID, and <value> with the desired value when using !addbal or !usebal.
- Use "!addbal" to add a balance to a client's inventory. For example, "!addbal 1234567890 10" adds 10 to the balance for the client with ID 1234567890.
- Use "!usebal" to deduct a balance from a client's inventory. For example, "!usebal 1234567890 5" deducts 5 from the balance for the client with ID 1234567890.
- Use "!register" to register account at store
- Use "!bal" to check balance
```
### **Client**
```
**Command: !buy <item> <quantity>**
- Buy items from the store.
	Example: !buy cid 5
	
**Command: !register**
- register your account discord
	Example: !register

**Command: !bal**
- check balance you acocunt
	Example: !bal

**Note:**
- Make sure to use the correct format for the commands.
- Replace <item> and <quantity> with the specific item and quantity you want to buy.
- Use "!register" to register account at store
- Use "!bal" to check balance
```

# command !register

### **Rest API**

- ***Sucess***
```
Metode 	: POST
Body	: {id:integer}
Status	: 201
Response: {name:"Success register account"}
```
- ***Failed***
```
Metode 	: POST
Body	: {id:integer}
Status	: 401
Response: {name:"Account already register"}
```

### **Chat Response**

- ***Sucess***
```
Message:"Success register account"
```
- ***Failed***
```
Message:"Account already register"
```

# command !buy cid quantity

### **Rest API**
- ***Sucess***
```
Metode 	: PUT
Header	: {id:integer}
Body	: {quantity:integer}
Status	: 200
Response: {name:"uniqcode\nuniqcode"}
```

- ***Failed***
```
Metode 	: GET
Header	: {id:integer}
Body	: {quantity:integer}
Status	: 400
Response: {name:"account not register"}
Reason	: Account not register
```
```
Metode 	: PUT
Header	: {id:integer}
Body	: {quantity:integer}
Status	: 400
Response: {name:"stock availabel in 1"}
Reason	: Stock in database not availabel
```
```
Metode 	: PUT
Header	: {id:integer}
Body	: {quantity:integer}
Status	: 400
Response: {name:"items are sold-out"}
Reason	: Stock in database 0
```
```
Metode 	: PUT
Header	: {id:integer}
Body	: {quantity:integer}
Status	: 400
Response: {name:"balance is low"}
Reason	: Balance is low
```
```
Metode 	: PUT
Header	: {id:integer}
Body	: {quantity:integer}
Status	: 400
Response: 
{name:"quantity cannot be null"}
{name:"quantity is integer"}
{name:"Invalid quantity. Please provide a quantity greater than or equal to 1."}
{name:"id cannot be null"}
Reason	: Wrong Format 
```
### **Chat Response**

- ***Sucess***
```
Message:"CID items at here <file.txt>"
```
- ***Failed***
```
Message:"account not register"
		"stock availabel in 1"
		"items are sold-out"
		"balance is low"
		"quantity is integer"
		"Invalid quantity. Please provide a quantity greater than or equal to 1."
		"id cannot be null"
		"Invalid input. Please use the format "!buy cid 1" and provide a valid numeric value."
```

# command !addbal idClient value

## OWNER
### **Rest API**
- ***Sucess***
```
Metode 	: PUT
Header	: {id:integer}
Body	: {value:integer}
Status	: 200
Response: {name:"Success add balance"}
```

- ***Failed***
```
Metode 	: PUT
Header	: {id:integer}
Body	: {value:integer}
Status	: 400
Response: {name:"account not register"}
Reason	: Account not register
```
```
Metode 	: PUT
Header	: {id:integer}
Body	: {value:integer}
Status	: 400
Response: 
{name:"value cannot be null"}
{name:"value is integer"}
{name:"Invalid value. Please provide a value greater than or equal to 1."}
{name:"id cannot be null"}
Reason	: Wrong Format 
```

### **Chat Response**

- ***Sucess***
```
Message:"Success add balance"
```
- ***Failed***
```
Message:"account not register"
		"value is integer"
		"Invalid value. Please provide a value greater than or equal to 1."
		"id cannot be null"
```

## CLIENT
- ***Failed***
### **Rest API**
```
Metode 	: PUT
Header	: {id:integer}
Body	: {quantity:integer}
Status	: 400
Response: {name:"Authorization"}
```

### **Chat Response**
```
Message:"Authorization"
```

# command !delbal idClient value

### **Rest API**
- ***Sucess***
```
Metode 	: PUT
Header	: {id:integer}
Body	: {value:integer}
Status	: 200
Response: {name:"Success deduct balance"}
```

- ***Failed***
```
Metode 	: PUT
Header	: {id:integer}
Body	: {value:integer}
Status	: 400
Response: {name:"account not register"}
Reason	: Account not register
```
```
Metode 	: PUT
Header	: {id:integer}
Body	: {value:integer}
Status	: 400
Response: {name:"balance now 1 you cannot deduct in 2"}
Reason	: balance is low from deduct
```
```
Metode 	: PUT
Header	: {id:integer}
Body	: {value:integer}
Status	: 400
Response: 
{name:"value cannot be null"}
{name:"value is integer"}
{name:"Invalid value. Please provide a value greater than or equal to 1."}
{name:"id cannot be null"}
Reason	: Wrong Format 
```

### **Chat Response**

- ***Sucess***
```
Message:"Success deduct balance"
```
- ***Failed***
```
Message:"account not register"
		"balance now 1 you cannot deduct in 2"
		"value is integer"
		"Invalid value. Please provide a value greater than or equal to 1."
		"id cannot be null"
```

## CLIENT
- ***Failed***
### **Rest API**
```
Metode 	: PUT
Header	: {id:integer}
Body	: {quantity:integer}
Status	: 400
Response: {name:"Authorization"}
```

### **Chat Response**
```
Message:"Authorization"
```

# command !bal

### **Rest API**
- ***Sucess***
```
Metode 	: GET
Header	: {id:integer}
Status	: 200
Response: {name:"you balace : 1"}
```

- ***Failed***
```
Metode 	: GET
Header	: {id:integer}
Status	: 400
Response: {name:"account not register"}
Reason	: Account not register
```
### **Chat Response**
- ***Sucess***
```
Message:"you balace : 1"
```
- ***Failed***
```
Message:"account not register"
```
