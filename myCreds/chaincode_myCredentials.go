
package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// Executes when each peer deploys its instance of the chaincode. It starts the chaincode and registers it with the peer.
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

//==============================================================================================================================

// Init resets all the things
// GJS: THIS is where we will create the user ledger
// MVP UC 1
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//We expect 3 arguments:  Social Security Number,  Full name, Date of birth
	if len(args) != 3 {
        	return nil, errors.New("Incorrect number of arguments. Expecting 1")
    	}
	
	//GJS  social security number
    	err1 := stub.PutState("social_security_number", []byte(args[0]))
    	if err1 != nil {
        	return nil, err1
    	}
	//GJS  full name
    	err2 := stub.PutState("full_name", []byte(args[1]))
    	if err2 != nil {
        	return nil, err2
    	}
	//GJS  date of birth, treat as string for now
    	err3 := stub.PutState("date_of_birth", []byte(args[2]))
    	if err3 != nil {
        	return nil, err3
    	}

    	return nil, nil
}

//==============================================================================================================================

// Invoke is our entry point to invoke a chaincode function
// Invoke functions are captured as transactions, which get grouped into blocks for writing to the ledger. 
// Updating the ledger is achieved by invoking the chaincode.
// MVP UC 2
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	fmt.Println("invoke myCredentials is running " + function)

    	// Handle different functions
    	if function == "init" {
        	return t.Init(stub, "init", args)
    	} else if function == "write" {
        	return t.write(stub, args)
    	}
    	fmt.Println("invoke did not find func: " + function)

    	return nil, errors.New("Received unknown function invocation")

}

//WRITE
// [GJS] Will focus only on adding degree certificates for the time being.
// This requires a number of details
//  - Certification Type (degree, diploma)
//  - Institution Name (MIT, Harvard)
//  - Degree Name (MBA, .Eng)
//  - DateStart
//  - DateEnd
//  - Other Details (comma delimited) Specialization, Cum Laude etc  
// MVP UC 2a
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var name, value string
    	var err error
    	fmt.Println("running write() for myCredentials")

    	if len(args) != 12 {
        	return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
    	}

	//we have 6 key value pairs that (for now at least) capture everything we want to know about the degree
	//GJS use every set of args (0 & 1, 2 & 3, etc) to store as key value pairs against the 6 variables we wish to track.  
	//For each of the 6 instances, one can create quality validation code, but out of scope for this MVP
	//NOTE: OPEN QUESTION - If I want to add a second degree, will this code below overwite the data of the degree that was added first
	//So, the question is... should one really have created a structure to pack the metadata in and create a unique name value for each university?
	//in that case one would perhaps be forced to write mor sophisticated query code.
	//NOTE: SmartContracts are considered for this MVP

	//  - Certification Type (degree, diploma)
    	name = args[0]                            
    	value = args[1]
    	err = stub.PutState(name, []byte(value))  //write the variable into the chaincode state
    	if err != nil {
        	return nil, err
    	}

	//  - Institution Name (MIT, Harvard)
    	name = args[2]                            
    	value = args[3]
    	err = stub.PutState(name, []byte(value))  //write the variable into the chaincode state
    	if err != nil {
        	return nil, err
    	}

	//  - Degree Name (MBA, .Eng)
    	name = args[4]                            
    	value = args[5]
    	err = stub.PutState(name, []byte(value))  //write the variable into the chaincode state
    	if err != nil {
        	return nil, err
    	}

	//  - DateStart
    	name = args[6]                            
    	value = args[7]
    	err = stub.PutState(name, []byte(value))  //write the variable into the chaincode state
    	if err != nil {
        	return nil, err
    	}

	//    - DateEnd
    	name = args[8]                            
    	value = args[9]
    	err = stub.PutState(name, []byte(value))  //write the variable into the chaincode state
    	if err != nil {
        	return nil, err
    	}

	//    - Other Details (comma delimited) Specialization, Cum Laude etc  
    	name = args[10]                            
    	value = args[11]
    	err = stub.PutState(name, []byte(value))  //write the variable into the chaincode state
    	if err != nil {
        	return nil, err
    	}

    	return nil, nil
}



//==============================================================================================================================

// Query is our entry point for queries
// MVP UC 3
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	fmt.Println("query is running " + function)

    	// Handle different functions
    	if function == "read" {                            
		//read a variable
        	return t.read(stub, args)
    	}
    	fmt.Println("query did not find func: " + function)

    	return nil, errors.New("Received unknown function query")

}

//READ
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    	var name, jsonResp string
    	var err error
	//GJS - the question that I have is that if I have several degrees, should the 6 variable names have a unique qualifyer that
	//points to the particular implementation so that I do not overwrite the first degree's data when I add the second
    	if len(args) != 1 {
        	return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
    	}

    	name = args[0]
    	valAsbytes, err := stub.GetState(name)
    	if err != nil {
        	jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
        	return nil, errors.New(jsonResp)
    	}

    	return valAsbytes, nil
}
