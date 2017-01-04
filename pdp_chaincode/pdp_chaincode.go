package main
import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func main() {
	fmt.Println("pdp ... main is running ")
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
//func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error)
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("pdp ->... init is running " + function)
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3: Nome,Professione,Disoccupato")
	}

	err := stub.PutState("Nome", []byte(args[0]))
	if err != nil {
		return nil, err
	}
	err1 := stub.PutState("Professione", []byte(args[1]))
	if err1 != nil {
		return nil, err1
	}
	err2 := stub.PutState("Dicoccupato", []byte(args[2]))
	if err2 != nil {
		return nil, err2
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("pdp ... invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("pdp ... query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key1,key2,key3, value1,value2,value3 string
	var err error
	fmt.Println("running write()")

	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6. name of the keys (nome,professione,disoccupato and value to set")
	}

	 
	key1 = args[0]  
	value1 = args[1]
	err = stub.PutState(key1, []byte(value1)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	key2 = args[2]  
	value2 = args[3]
	err = stub.PutState(key2, []byte(value2)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	key3 = args[4]  
	value3 = args[5]
	err = stub.PutState(key3, []byte(value3)) //write the variable into the chaincode state
	
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	fmt.Printf("Array di valori letti %s\n", valAsbytes)
	return valAsbytes, nil
}


