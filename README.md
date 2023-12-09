# command !help

### **OWNER**
```
**Command: !buy <item> <quantity>**
- Buy items from the store.
	Example: !buy cid 5

**Command: !addball <id client> <value>**
- Add a balance to a client's inventory.
	Example: !addball 1234567890 10

**Command: !useball <id client> <value>**
- Use a balance from a client's inventory.
	Example: !useball 1234567890 5

**Command: !register**
- register your account discord
	Example: !register

**Note:**
- Make sure to use the correct format for the commands.
- Replace <item> and <quantity> with the specific item and quantity you want to buy.
- Replace <id client> with the client ID, and <value> with the desired value when using !addball or !useball.
- Use "!addball" to add a balance to a client's inventory. For example, "!addball 1234567890 10" adds 10 to the balance for the client with ID 1234567890.
- Use "!useball" to deduct a balance from a client's inventory. For example, "!useball 1234567890 5" deducts 5 from the balance for the client with ID 1234567890.
- Use "!register" to register account at store
```
### **Client**
```
**Command: !buy <item> <quantity>**
- Buy items from the store.
	Example: !buy cid 5
	
**Command: !register**
- register your account discord
	Example: !register

**Note:**
- Make sure to use the correct format for the commands.
- Replace <item> and <quantity> with the specific item and quantity you want to buy.
- Use "!register" to register account at store
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