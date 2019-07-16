//#### User Registering
//1. Register User: data_fields:
//     User Type(Owner/tenant)
//     First Name
//     Surname
//     DOB
//     Contact Number
//     Email ID
//     Contact Address--------

// CHAICODE FOR USER REGISTRATION

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// STRUCT registrationManager-----------

type registrationManager struct {
}

// STRUCT userReg----------

type userReg struct {
	uType     string // (Owner/tenant)
	uID       string // userID
	fName     string // First Name
	lName     string // Last Name
	dob       uint32 // Date of Birth
	contactNo uint64 // Contact Number
	email     string // Email ID
	address   string // Address
}

func (t *registrationManager) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Incorrect arguments. Expecting a key and a value")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}

	return shim.Success(nil)
}

func (t *registrationManager) createUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 7 {
		return shim.Error("ERROR: Incorrect number of arguments. Expecting 8 arguments \n")
	}

	userAsBytes, err := stub.GetState(args[1])
	if err != nil {
		return shim.Error("ERROR: User Get operation failed ")
	}

	if userAsBytes == nil {
		return shim.Error("ERROR: USER already exist with id" + args[1])
	}

	dob, err := strconv.ParseUint(args[4], 10, 32)
	if err != nil {
		return shim.Error("ERROR: DOB parsing failed ")
	}
	contactNo, err := strconv.ParseUint(args[5], 10, 64)
	if err != nil {
		return shim.Error("ERROR: contactNo parsing failed ")
	}

	var user = userReg{
		uType:     args[0],
		uID:       args[1],
		fName:     args[2],
		lName:     args[3],
		dob:       dob,
		contactNo: contactNo,
		email:     args[6],
	}

	userInBytes, err := json.Marshal(user)
	if err != nil {
		shim.Error("ERROR: Marshalling unsuccessful ")
	}

	err = stub.PutState(args[1], userInBytes)
	if err != nil {
		shim.Error(err.Error())
	}

	CKey1, err := stub.CreateCompositeKey("uType~uID~contactNo",
		[]string{
			user.uType,
			user.uID,
			strconv.FormatUint(user.contactNo, 10)})

	value := []byte{0x00}
	err = stub.PutState(Ckey1, value)
	if err != nil {
		shim.Error(err.Error())
	}
}

func (t *registrationManager) getUserByID(stud shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("ERROR: Incorrect number of arguments. Expecting 1 ")
	}

	userInBytes, err := stub.GetState(string.ToLower(args[0]))
	if err != nil {
		shim.Error("ERROR: Failed to get the state of ID " + args[0])
	}

	if userInBytes == nil {
		shim.Error("ERROR: No data available for user Id " + args[0])
	}

	return shim.Success(userInBytes)
}

func (t *registrationManager) updateUser(stub shim.ChaincodeStubInterface, args []string) {
	if len(args) != 7 {
		return shim.Error("ERROR: Incorrect number of arguments. Expecting 7 ")
	}

	userInBytes, err := stub.GetState(string.ToLower(args[1]))
	if err != nil {
		shim.Error("ERROR: Failed to get the state of ID " + args[0])
	}

	if userInBytes == nil {
		shim.Error("ERROR: No data available for user Id " + args[0])
	}

	user := userReg{}

	json.Unmarshal(userInBytes, &user)

	dob, err := strconv.ParseUint(args[4], 10, 32)
	if err != nil {
		return shim.Error("ERROR: DOB parsing failed ")
	}
	contactNo, err := strconv.ParseUint(args[5], 10, 64)
	if err != nil {
		return shim.Error("ERROR: contactNo parsing failed ")
	}

	user.uType = args[0]
	user.fName = args[2]
	user.lName = args[3]
	user.dob = dob
	user.contactNo = contactNo
	user.email = args[6]

	userInBytes, err = json.Marshal(user)
	if err != nil {
		return shim.Error("ERROR: Marshalling unsuccessful ")
	}

	err := stub.PutState(args[1], userInBytes)
	if err != nil {
		return shim.Error("ERROR: No update made ")
	}

	CKeyUpdate, err := stub.CreateCompositeKey("uType~uID~contactNo",
		[]string{
			user.uType,
			user.uID,
			strconv.FormatUint(user.contactNo, 10)})

	value := []byte{0x00}
	err = stub.PutState(CkeyUpdate, value)
	if err != nil {
		shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *registrationManager) deleteUser(stub shim.ChaincodeStubInterface, args []string) {
	if len(args) != 1 {
		shim.Error("ERROR: Incorrect number of arguments. Exoecting 1 ")
	}

	var user = userReg{}
	userID = args[0]

	userInBytes, err := stub.GetState(userID)
	if err != nil {
		shim.Error(err.Error())
	}

	if userInBytes == nil {
		return shim.Error("ERROR: No data is available for user ID " + args[0])
	}

	err := json.Unmarshal(userInBytes, &user)
	if err != nil {
		shim.Error("ERROR: Unmarshall unsuccessfull")
	}

	err = stub.DelState(userID)
	if err != nil {
		shim.Error("ERROR: Failed to delete state for userID " + args[0])
	}

	CKeyDel, err := stub.CreateCompositeKey("uType~uID~contactNo",
		[]string{
			user.uType,
			user.uID,
			strconv.FormatUint(user.contactNo, 10)})

	err = stub.DelState(userID)
	if err != nil {
		shim.Error("ERROR: Failed to delete state for Composite Key of " + args[0])
	}
}

func (t *registrationManager) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "createUser" {
		return t.createUser(stub, args)
	}

	if function == "getUserByID" {
		return t.getUserByID(stub, args)
	}

	if function == "updateUser" {
		return t.updateUser(stub, args)
	}

	if function == "deleteUser" {
		return t.deleteUser(stub, args)
	}
}
