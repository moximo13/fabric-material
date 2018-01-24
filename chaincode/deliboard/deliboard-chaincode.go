// VERSION 1.0.0 BY AU 23/1/2018
// STUDY FROM: fabcar AND tuna
// BASIC FUNCTION INCLUDE: BUILD LEDGER, CREATE REQUEST, AND QUERRY ALL REQUEST

// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
 */

package main

/* Imports  
* 4 utility libraries for handling bytes, reading and writing JSON, 
formatting, and string manipulation  
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts  
*/ 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Tuna structure, with 4 properties.  
Structure tags are used by encoding/json library
*/
type request struct {							// IN DEVELOPMENT...
	Timestamp	string	`json:"timestamp"`
	SenderName 	string	`json:"sender_name"`
	SenderAddr	string	`json:"sender_address"`
	ReceiveName	string	`json:"receive_name"`
	ReceiveAddr	string	`json:"receive_address"`
	DelivererName	string	`json:"deliverer_name"`
	Price		string	`json:"price"`	
	Status		string	`json:"status"`	
	Code		string	`json:"code"`	
		
}

/*
 * The Init method *
 called when the Smart Contract "tuna-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function 
 -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "tuna-chaincode"
 The app also specifies the specific smart contract function to call with args
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "buildLedger" {
		return s.buildLedger(APIstub)
	} else if function == "buildRequest" {
		return s.buildRequest(APIstub, args)
	} else if function == "queryAllRequest" {
		return s.queryAllRequest(APIstub)
	} else if function == "queryRequest" {
		return s.queryAllRequest(APIstub)
	}

	return shim.Error("Unable to identify Smart Contract function name.")
}

/*
 * The queryTuna method *
Used to view the records of one particular tuna
It takes one argument -- the key for the tuna in question
 */
func (s *SmartContract) queryRequest(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	requestAsBytes, _ := APIstub.GetState(args[0])
	if requestAsBytes == nil {
		return shim.Error("Could not locate request")
	}
	return shim.Success(requestAsBytes)
}

/*
 * The initLedger method *
Will add test data (10 tuna catches)to our network
 */
func (s *SmartContract) buildLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	Requests := []request{
		request{Timestamp: "2145376543", SenderName: "923F", SenderAddr: "L150HA", ReceiveName: "Alex Forkford", ReceiveAddr: "Bangkok", DelivererName: "de wea", Price: "10", Status: "under way", Code: "98sad56w4d56"},
	}

	i := 0
	for i < len(Requests) {
		fmt.Print("i is ", i)
		requestAsBytes, _ := json.Marshal(Requests[i])
		APIstub.PutState(strconv.Itoa(i+1), requestAsBytes)
		fmt.Println("Added", Requests[i])
		i = i + 1
	}
	return shim.Success(nil)
}

/*
 * The recordTuna method *
Fisherman like Sarah would use to record each of her tuna catches. 
This method takes in five arguments (attributes to be saved in the ledger). 
 */
func (s *SmartContract) buildRequest(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect input detected: need only key, time, sender name, sen addr, receive_name, receive_address")
	}

	var newRequest = request{Timestamp: args[1], SenderName: args[2], SenderAddr: args[3], ReceiveName: args[4], ReceiveAddr: args[5], DelivererName: args[6], Price: args[7], Status: args[8], Code: args[9]}
	

	requestAsBytes, _ := json.Marshal(newRequest)
	err := APIstub.PutState(args[0], requestAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record new request: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The queryAllTuna method *
allows for assessing all the records added to the ledger(all tuna catches)
This method does not take any arguments. Returns JSON string containing results. 
 */
func (s *SmartContract) queryAllRequest(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer

	// HERE GO NOTHING!!!!
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}\n")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllRequest:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/*
 * The changeTunaHolder method *
The data in the world state can be updated with who has possession. 
This function takes in 2 arguments, tuna id and new holder name. 
 */
// func (s *SmartContract) changeTunaHolder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

// 	if len(args) != 2 {
// 		return shim.Error("Incorrect number of arguments. Expecting 2")
// 	}

// 	tunaAsBytes, _ := APIstub.GetState(args[0])
// 	if tunaAsBytes == nil {
// 		return shim.Error("Could not locate tuna")
// 	}
// 	tuna := Tuna{}

// 	json.Unmarshal(tunaAsBytes, &tuna)
// 	// Normally check that the specified argument is a valid holder of tuna
// 	// we are skipping this check for this example
// 	tuna.Holder = args[1]

// 	tunaAsBytes, _ = json.Marshal(tuna)
// 	err := APIstub.PutState(args[0], tunaAsBytes)
// 	if err != nil {
// 		return shim.Error(fmt.Sprintf("Failed to change tuna holder: %s", args[0]))
// 	}

// 	return shim.Success(nil)
// }

/*
 * main function *
calls the Start function 
The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Something wrong while creating new Smart Contract: %s", err)
	}
}
