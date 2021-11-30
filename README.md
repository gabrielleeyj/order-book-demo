# order-book-demo
A simulation to create an order book. From a known dataset.

## Instruction
1. Clone repo.
2. Run `go mod tidy` at the root of the directory.
3. Run `go build order-book-demo` to run the program.
4. Run `go test` to run the test suite.

### Operations
The order book will require the following operations

**Add** - Submits and Order  
**Cancel** - Cancels the Order based on order_id

#### Sample Inputs

Where,

**A = Add Order**    
**X = Cancel Order**

**S** - *Sell*  
**B** - *Buy*
```
A,100000,S,1,1075
A,100001,B,9,1000 
A,100002,B,30,975 
A,100003,S,10,1050 
A,100004,B,10,950 
A,100005,S,2,1025 
A,100006,B,1,1000 
X,100004,B,10,950 
A,100007,S,5,1025 
A,100008,B,3,1050 
X,100008,B,3,1050 
X,100005,S,2,1025
```
