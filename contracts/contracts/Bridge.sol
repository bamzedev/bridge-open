//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;
import "./WrappedToken.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Bridge is Ownable{
    address payable private feeCollector;
    uint256 public fee;
    mapping (string => bool) public isProccessed;
    mapping(string => address[]) public signaturesForTransaction;
    mapping(address => address) public wrappedTokenContracts;
    mapping(address => address) public nativeTokenContracts;
    mapping (address => bool) private isValidator;
    event NewTokenDeployed(address indexed tokenContract);
    event TokenMint(address indexed from, address indexed sourceTokenAddress, uint256 amount);
    event TokenUnlock(address indexed from, address indexed sourceTokenAddress, uint256 amount);
    event TokenBurn(address indexed from, address indexed sourceTokenAddress, uint256 amount);
    event TokenLock(address indexed from, address indexed sourceTokenAddress, uint256 amount);

    constructor (address[] memory _validators, address _feeColector, uint256 _fee){
        for (uint256 i = 0; i < _validators.length; i++) {
            isValidator[_validators[i]]=true;
        }
        feeCollector = payable(_feeColector);
        fee = _fee;
    }

    modifier enoughFee() {
        require(msg.value == fee, "not enough ether");
        _;
    }

    modifier notProccessed(string memory _transaction) {
        require(!isProccessed[_transaction], "Transaction already processed");
        _;
    }

    function validateMint(string memory _transaction, address _token, string memory _name, string memory _symbol, uint256 _amount, address _receiver, uint8[]memory v, bytes32[]memory r, bytes32[]memory s) internal{
        bytes32 messageDigest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32",keccak256(abi.encodePacked(block.chainid,_transaction,_token, _name, _symbol, _amount,_receiver))));
        for (uint i = 0; i < v.length; i++) {
            address currentAddress = ecrecover(messageDigest, v[i], r[i], s[i]);
            require(isValidator[ecrecover(messageDigest, v[i], r[i], s[i])], "Wrong signature");
            signaturesForTransaction[_transaction].push(currentAddress);            
        }
        require(signaturesForTransaction[_transaction].length>1, "Need at least 2 signatures");
    }

    function validateUnlock(string memory _transaction, address _token, uint256 _amount, address _receiver, uint8[]memory v, bytes32[]memory r, bytes32[]memory s) internal{
        bytes32 messageDigest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32",keccak256(abi.encodePacked(block.chainid,_transaction,_token, _amount,_receiver))));
        for (uint i = 0; i < v.length; i++) {
            address currentAddress = ecrecover(messageDigest, v[i], r[i], s[i]);
         require(isValidator[ecrecover(messageDigest, v[i], r[i], s[i])], "Wrong signature");
            signaturesForTransaction[_transaction].push(currentAddress);            
        }
        require(signaturesForTransaction[_transaction].length>1, "Need at least 2 signatures");
    }

    function deployContract(address _nativeTokenAddress, string memory _name, string memory _symbol) private{
      address newTokenAddress = address(new WrappedToken(_name, _symbol));
      wrappedTokenContracts[_nativeTokenAddress] = newTokenAddress;
      nativeTokenContracts[newTokenAddress] = _nativeTokenAddress;
      emit NewTokenDeployed(wrappedTokenContracts[_nativeTokenAddress]);
    }

    function unlockTokens(address _token, uint256 _amount, string memory _transaction, 
    uint8[]memory _v, bytes32[]memory _r, bytes32[]memory _s) public notProccessed(_transaction){
        validateUnlock(_transaction, _token, _amount, msg.sender, _v, _r, _s); 
        WrappedToken token = WrappedToken(_token);
        isProccessed[_transaction]=true;
        token.transfer(msg.sender, _amount);
        emit TokenUnlock(msg.sender, _token, _amount);
    }

    function mintTokens(address _nativeTokenAddress, string memory _name, string memory _symbol, uint256 _amount, string memory _transaction, 
    uint8[]memory _v, bytes32[]memory _r, bytes32[]memory _s) public notProccessed(_transaction){
       validateMint(_transaction, _nativeTokenAddress, _name, _symbol, _amount, msg.sender, _v, _r, _s); 
       if(wrappedTokenContracts[_nativeTokenAddress]== address(0)){
        deployContract(_nativeTokenAddress, _name, _symbol);
       }
       WrappedToken wtoken = WrappedToken(wrappedTokenContracts[_nativeTokenAddress]);
       isProccessed[_transaction]=true;
       wtoken.mintTo(msg.sender, _amount);
       emit TokenMint(msg.sender, address(wtoken), _amount);
    }
    
    function lock(address _owner, address _token, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) public payable enoughFee{
        feeCollector.transfer(1);
        WrappedToken wt = WrappedToken(_token);
        wt.permit(_owner, address(this), _amount, _deadline, _v, _r, _s);
        wt.transferFrom(_owner, address(this), _amount);
        emit TokenLock(_owner, _token, _amount);
    } 

    function burn(address _owner, address _token, uint256 _amount, uint8 _v, bytes32 _r, bytes32 _s) public payable enoughFee{
        feeCollector.transfer(1);
        bytes32 messageDigest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32",keccak256(abi.encodePacked(_owner, address(this),_token,_amount))));
        require(_owner == ecrecover(messageDigest, _v, _r, _s), "Wrong Signature");            
        WrappedToken wtoken = WrappedToken(_token);
        wtoken.burnFrom(_owner, _amount);
        //Emits native token address from the other chain, so the other chain knows which token to unlock
        emit TokenBurn(_owner, nativeTokenContracts[_token], _amount); 
    }
}
