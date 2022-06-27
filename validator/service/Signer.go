package service

import (
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getContractFromTransaction(tx Transaction) *Service {
	var infuraUrl string
	var contractAddress common.Address
	if tx.FromChainId == "3" {
		infuraUrl = "wss://ropsten.infura.io/ws/v3/" + configuration.InfuraKey
		contractAddress = common.HexToAddress(configuration.BridgeContractAddressRopsten)
	} else {
		infuraUrl = "wss://rinkeby.infura.io/ws/v3/" + configuration.InfuraKey
		contractAddress = common.HexToAddress(configuration.BridgeContractAddressRinkeby)
	}
	client, err := ethclient.Dial(infuraUrl)
	if err != nil {
		panic(err)
	}
	contract, _ := NewService(contractAddress, client)
	return contract
}

func findLockTransaction(tx Transaction) (Transaction, error) {
	if(tx.IsBurn){
		locks, err := getContractFromTransaction(tx).FilterTokenBurn(&bind.FilterOpts{}, nil, nil)
		if err != nil {
			panic(err)
		}
		found := false
		for locks.Next() {
			event := locks.Event
			if event.Raw.TxHash.String() == tx.LockTransactionHash {
				found = true
				tx.LockTransactionHash = event.Raw.TxHash.String()
				tx.Recipient = event.From.String()
				tx.TokenAddress = event.SourceTokenAddress.String()
				tx.Amount = event.Amount.String()
				break
			}
		}
	
		if !found {
			return Transaction{}, err
		}
	}else{
		locks, err := getContractFromTransaction(tx).FilterTokenLock(&bind.FilterOpts{}, nil, nil)
		if err != nil {
			panic(err)
		}
		found := false
		for locks.Next() {
			event := locks.Event
			if event.Raw.TxHash.String() == tx.LockTransactionHash {
				found = true
				tx.LockTransactionHash = event.Raw.TxHash.String()
				tx.Recipient = event.From.String()
				tx.TokenAddress = event.SourceTokenAddress.String()
				tx.Amount = event.Amount.String()
				break
			}
		}
	
		if !found {
			return Transaction{}, err
		}
	}
	return tx, nil
}

func signTransaction(tx Transaction) (Transaction, error) {
	validator1PrivateKey, err := crypto.HexToECDSA(configuration.Validator1PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	validator2PrivateKey, err := crypto.HexToECDSA(configuration.Validator2PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	uint256Type, _ := abi.NewType("uint256", "", nil)
	args := abi.Arguments{
		{Type: uint256Type},
	}
	transaction := []byte(tx.LockTransactionHash)
	addressToken := common.HexToAddress(tx.TokenAddress).Bytes()
	var symbol []byte
	var name []byte
	if(tx.IsBurn){
		name = []byte(tx.Name)
		symbol = []byte(tx.Symbol)
	}else{
		name = []byte("Wrapped " + tx.Name)
		symbol = []byte("W" + tx.Symbol)
	}
	
	n, _ := new(big.Int).SetString(tx.Amount, 0)
	n1, _ := new(big.Int).SetString(tx.ToChainId, 0)
	amount, _ := args.Pack(n)
	toChainId, _ := args.Pack(n1)
	addressReceiver := common.HexToAddress(tx.Recipient).Bytes()
	hash := crypto.Keccak256Hash(toChainId, transaction, addressToken, name, symbol, amount, addressReceiver)
	msg := crypto.Keccak256([]byte("\x19Ethereum Signed Message:\n32"), hash.Bytes())
	signature1, err := crypto.Sign(msg, validator1PrivateKey)
	if err != nil {
		log.Fatal(err)
		return tx, err
	}
	signature2, err := crypto.Sign(msg, validator2PrivateKey)
	if err != nil {
		log.Fatal(err)
		return tx, err
	}
	tx.Signatures[0] = hexutil.Encode(signature1)
	tx.Signatures[1] = hexutil.Encode(signature2)
	tx.Date = time.Now().Unix()
	return tx, nil

}
func Sign(tx Transaction) (Transaction, error) {
	lockTransaction, err := findLockTransaction(tx)
	if err != nil {
		return Transaction{}, err
	}

	signedTransaction, err := signTransaction(lockTransaction)
	if err != nil {
		return Transaction{}, err
	}

	return signedTransaction, nil
}
func UpdateClaimed(tx Transaction) (Transaction, error) {
	var infuraUrl string
	var contractAddress common.Address
	if tx.ToChainId == "3" {
		infuraUrl = "wss://ropsten.infura.io/ws/v3/" + configuration.InfuraKey
		contractAddress = common.HexToAddress(configuration.BridgeContractAddressRopsten)
	} else if tx.ToChainId == "4" {
		infuraUrl = "wss://rinkeby.infura.io/ws/v3/" + configuration.InfuraKey
		contractAddress = common.HexToAddress(configuration.BridgeContractAddressRinkeby)
	}

	client, err := ethclient.Dial(infuraUrl)
	if err != nil {
		panic(err)
	}
	contract, _ := NewService(contractAddress, client)
	claimed, _ := contract.IsProccessed(&bind.CallOpts{}, tx.LockTransactionHash)
	tx.Claimed = claimed
	return tx, nil
}
