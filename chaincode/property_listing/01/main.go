/*
   * Property ID
   * Owner User ID
   * Property status(available/unavailable)
   * Property For(Renting/Acquisation)
   * Property address
   * Property aesthetic attributes
   * type(flat/duplex/villa/triplex)
   * carpet area
   * state of property
       * furnisher(t/f)
       * unfurnished(t/f)
   * Property age
   * rooms
   * balcony
   * parking(Bike/car)
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type propertyListing struct {
}

type propertyDetails struct {
	ptype      string //	(1bhk,2bhk,3bhk,4bhk,5bhk)
	buildArea  uint64 //	(sqft)
	buildType  string //  	(flat/duplex/villa/triplex)
	furnishing string //  	(furnished/semi-furnished/fullyfurnished)
	pAge       uint32
	carPark    bool
	bikePark   bool
}

type property struct {
	pID      string
	uID      string
	pStatus  string //  (available/unavailable)
	pUse     string //  (Renting/Sell)
	address  string
	pDetails propertyDetails // struct
}

func (t *propertyListing) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Incorrect arguments. Expecting a key and a value ")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}

	return shim.Success(nil)
}

func (t *propertyListing) createProperty(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 12 {
		return shim.Error("ERROR: Incorrect number of arguments. Expecting 12 arguments \n")
	}

	propertyAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("ERROR: Property Get operation failed ")
	}

	if propertyAsBytes == nil {
		return shim.Error("ERROR: Property already exists with id" + args[1])
	}

	buildArea, err := strconv.ParseUint(args[6], 10, 64)
	if err != nil {
		return shim.Error("ERROR: Build Area parsing failed\n")
	}

	pAge, err := strconv.ParseUint(args[9], 10, 32)
	if err != nil {
		return shim.Error("ERROR: Property Age parsing failed\n")
	}

	carPark, err := strconv.ParseBool(args[10])
	if err != nil {
		return shim.Error("ERROR: Cark Park Status parsing failed\n")
	}

	bikePark, err := strconv.ParseBool(args[11])
	if err != nil {
		return shim.Error("ERROR: Bike Park Status parsing failed\n")
	}
	var property = property{
		pID:     args[0],
		uID:     args[1],
		pStatus: args[2],
		pUse:    args[3],
		address: args[4],
		pDetails: propertyDetails{
			ptype:      args[5],
			buildArea:  buildArea,
			buildType:  args[7],
			furnishing: args[8],
			pAge:       uint32(pAge),
			carPark:    carPark,
			bikePark:   bikePark,
		},
	}

	propertyAsBytes, err = json.Marshal(property)
	if err != nil {
		return shim.Error("ERROR: Marshalling unsuccessful ")
	}

	err = stub.PutState(args[1], propertyAsBytes)
	if err != nil {
		return shim.Error("ERROR: " + err.Error())
	}

	cKey1, err := stub.CreateCompositeKey("pID~uID~pStatus~pUse", []string{
		property.pID,
		property.uID,
		property.pStatus,
		property.pUse,
	})

	value := []byte{0x00}
	err = stub.PutState(cKey1, value)
	if err != nil {
		shim.Error("ERROR: " + err.Error())
	}

	return shim.Success(nil)
}

func (t *propertyListing) getPropertyByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("ERROR: Incorrect number of arguments. Expecting 1 ")
	}

	propertyAsBytes, err := stub.GetState(strings.ToLower(args[0]))
	if err != nil {
		return shim.Error("ERROR: Failed to get the state of ID " + args[0])
	}

	if propertyAsBytes == nil {
		return shim.Error("ERROR: No data is available for property Id " + args[0])
	}

	return shim.Success(propertyAsBytes)
}

func (t *propertyListing) updateProperty(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 12 {
		return shim.Error("ERROR: Incorrect number of arguments. Expecting 12 ")
	}

	propertyAsBytes, err := stub.GetState(strings.ToLower(args[0]))
	if err != nil {
		shim.Error("ERROR: Failed to get the state of ID: " + args[0])
	}

	if propertyAsBytes == nil {
		shim.Error("ERROR: No data available for property ID: " + args[0])
	}

	property := property{}

	json.Unmarshal(propertyAsBytes, &property)

	buildArea, err := strconv.ParseUint(args[6], 10, 64)
	if err != nil {
		return shim.Error("ERROR: Build Area parsing failed\n")
	}

	pAge, err := strconv.ParseUint(args[9], 10, 32)
	if err != nil {
		return shim.Error("ERROR: Property Age parsing failed\n")
	}

	carPark, err := strconv.ParseBool(args[10])
	if err != nil {
		return shim.Error("ERROR: Cark Park Status parsing failed\n")
	}

	bikePark, err := strconv.ParseBool(args[11])
	if err != nil {
		return shim.Error("ERROR: Bike Park Status parsing failed\n")
	}

	property.pID = args[0]
	property.uID = args[1]
	property.pStatus = args[2]
	property.pUse = args[3]
	property.address = args[4]
	property.pDetails.ptype = args[5]
	property.pDetails.buildArea = buildArea
	property.pDetails.buildType = args[7]
	property.pDetails.furnishing = args[8]
	property.pDetails.pAge = uint32(pAge)
	property.pDetails.carPark = carPark
	property.pDetails.bikePark = bikePark

	propertyAsBytes, err = json.Marshal(property)
	if err != nil {
		return shim.Error("ERROR: Marshalling unsuccessfull \n")
	}

	err = stub.PutState(args[0], propertyAsBytes)
	if err != nil {
		return shim.Error("ERROR: No update made \n")
	}

	cKeyUpdate, err := stub.CreateCompositeKey("pID~uID~pStatus~pUser", []string{
		property.pID,
		property.uID,
		property.pStatus,
		property.pUse,
	})

	value := []byte{0x00}
	err = stub.PutState(cKeyUpdate, value)
	if err != nil {
		shim.Error("ERROR: " + err.Error())
	}

	return shim.Success(nil)
}

func (t *propertyListing) deleteProperty(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("ERROR: Incorrect number of arguments. Expecting 1 ")
	}

	var property = property{}
	pID := args[0]

	propertyAsBytes, err := stub.GetState(pID)
	if err != nil {
		shim.Error("ERROR: " + err.Error())
	}

	if propertyAsBytes == nil {
		return shim.Error(fmt.Sprintf("ERROR: No data is available for userID %s ", args[0]))
	}

	err = json.Unmarshal(propertyAsBytes, &property)
	if err != nil {
		return shim.Error("ERROR: Unmarshal unsuccessfull ")
	}

	err = stub.DelState(pID)
	if err != nil {
		return shim.Error(fmt.Sprintf("ERROR: Failed to delete state for state %s\n", args[0]))
	}

	cKeyDel, err := stub.CreateCompositeKey("pID~uID~pStatus~pUser", []string{
		property.pID,
		property.uID,
		property.pStatus,
		property.pUse,
	})

	err = stub.DelState(cKeyDel)
	if err != nil {
		shim.Error("ERROR: " + err.Error())
	}
	return shim.Success(nil)
}

func (t *propertyListing) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "createProperty" {
		return t.createProperty(stub, args)
	}

	if function == "getPropertyByID" {
		return t.getPropertyByID(stub, args)
	}

	if function == "updateProperty" {
		return t.updateProperty(stub, args)
	}

	if function == "deleteProperty" {
		return t.deleteProperty(stub, args)
	}

	return shim.Error("ERROR: Received unknown function invocation: " + function)
}

func main() {
	err := shim.Start(new(propertyListing))
	if err != nil {
		fmt.Printf("ERROR: Failed creating new Smart Contract: ")
	}
}
