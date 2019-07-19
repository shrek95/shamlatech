/*
3. Property monetary attributes for renting

* Property ID
* Monthly rent
* Deposit
* Monthly maintenance
* Additional Expenses

4. Property monetary attributes for selling

* Property ID
* Price
* Registery charges
* Maintenance expences
* Additional Expenses
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type propertyMonetaryVals struct {
}

type rentMonetaryVals struct {
	pID             string
	uID             string
	pType           string
	monthRent       uint64
	deposit         uint64
	maintenanceType string // monthly/yearly/onetime
	maintenanceVal  uint64
	addOnExpenses   uint64
}

type sellMonetaryVals struct {
	pID                 string
	uID                 string
	pType               string
	price               uint64
	govCharges          uint64
	maintenanceExpences uint64
	addOnExpenses       uint64
}

func (p *propertyMonetaryVals) Init(stub.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Incorrect arguments. Expecting a key and a value")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("ERROR: Failed to create asset: %s", args[0]))
	}

	return shim.Success(nil)
}

func (p *propertyMonetaryVals) createRentEntry(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 7 {
		return shim.Error("ERROR: Incorrect number of arguments. Expecting 7 \n")
	}

	rentValAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("ERROR: Property Get operation failed ")
	}

	if rentValAsBytes == nil {
		return shim.Error("ERROR: Property already exist with id " + args[0])
	}

	monthRent, err := strconv.ParseUint(args[2], 10, 64)
	if err != nil {
		return shim.Error("ERROR: Monthly rent parsing failed\n")
	}
	deposit, err := strconv.ParseUint(args[3], 10, 64)
	if err != nil {
		return shim.Error("ERROR: Deposit parsing failed\n")
	}
	maintenanceVal, err := strconv.ParseUint(args[5], 10, 64)
	if err != nil {
		return shim.Error("ERROR: Maintenance value parsing failed\n")
	}
	addOnExpenses, err := strconv.ParseUint(args[6], 10, 64)
	if err != nil {
		return shim.Error("ERROR: Add On Expenses parsing failed\n")
	}

	var rent = rentMonetaryVals{
		pID:             args[0],
		uID:             args[1],
		pType:           "rent",
		monthRent:       args[2],
		deposit:         args[3],
		maintenanceType: args[4],
		maintenanceVal:  args[5],
		addOnExpenses:   args[6],
	}

	rentValAsBytes, err := json.Marshal(rent)
	if err != nil {
		return shim.Error("ERROR: Marshaling unsuccessful")
	}

	err = stub.PutState(args[0], rentValAsBytes)
	if err != nil {
		return shim.Error("ERROR: " + err.Error())
	}

	cKey1, err := stub.CreateCompositeKey("pID~uID~pType", []string{
		rent.pID,
		rent.uID,
		rent.pType,
	})

	value := []byte{0x00}
	err = stub.PutState(cKey1, value)
	if err != nil {
		return shim.Error("ERROR: " + err.Error())
	}

	return shim.Success(nil)

	/*
		pID
		uID
		pType
		monthRent
		deposit
		maintenanceType
		maintenanceVal
		addOnExpences
	*/
}

func (p *propertyMonetaryVals) createSellEntry(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 6 {
		return shim.Error("ERROR: Incorrect number of arguments. Expecting 6 \n")
	}

	sellValAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("ERROR: Property Get operation failed ")
	}

	if sellValAsBytes == nil {
		return shim.Error("ERROR: Property already exist with id " + args[0])
	}

	price, err := strconv.ParseUint(args[2], 10, 64)
	if err != nil {
		return shim.Error("ERROR: Monthly rent parsing failed\n")
	}
	govCharges, err := strconv.ParseUint(args[3], 10, 64)
	if err != nil {
		return shim.Error("ERROR: Deposit parsing failed\n")
	}
	maintenanceExp, err := strconv.ParseUint(args[4], 10, 64)
	if err != nil {
		return shim.Error("ERROR: Maintenance value parsing failed\n")
	}
	addOnExpenses, err := strconv.ParseUint(args[5], 10, 64)
	if err != nil {
		return shim.Error("ERROR: Add On Expenses parsing failed\n")
	}
	/*
	   pID
	   uID
	   pType
	   price
	   govCharges
	   maintenanceExpences
	   addOnExpenses

	*/
	var sell = sellMonetaryVals{
		pID:            args[0],
		uID:            args[1],
		pType:          "sell",
		price:          price,
		govCharges:     govCharges,
		maintenanceExp: maintenanceExp,
		addOnExpenses:  addOnExpenses,
	}

	sellValsAsBytes, err := json.Marshal(sell)
	if err != nil {
		return shim.Error("ERROR: Marshaling unsuccessful")
	}

	err = stub.PutState(args[0], sell)
	if err != nil {
		return shim.Error("ERROR: " + err.Error())
	}

	cKey1, err := stub.CreateCompositeKey("pID~uID~pType", []string{
		sell.pID,
		sell.uID,
		sell.pType,
	})

	value := []byte{0x00}
	err = stub.PutState(cKey1, value)
	if err != nil {
		return shim.Error("ERROR: " + err.Error())
	}

	return shim.Success(nil)

}

func (p *propertyMonetaryVals) getAllOnRentProperties(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

func (p *propertyMonetaryVals) getAllOnSellProperties(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

func (p *propertyMonetaryVals) Invoke(stub.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "createRentEntry" {
		return p.createRentEntry(stub, args)
	}

	if function == "createSellEntry" {
		return p.createSellEntry(stub, args)
	}

	if function == "getAllOnRentProperties" {
		return p.getAllOnRentProperties(stub, args)
	}

	if function == "getAllOnSellProperties" {
		return p.getAllOnSellProperties(stub, args)
	}

	return shim.Error("ERROR: Received unknown function invocatio: " + function)

}
