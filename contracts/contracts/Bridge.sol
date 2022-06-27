//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;
import "./WrappedToken.sol";

contract Bridge {
    mapping (string => bool) public isProccessed;
    mapping(string => address[]) public signaturesForTransaction;
    mapping(address => address) public wrappedTokenContracts;
    mapping(address => address) public nativeTokenContracts;
    mapping (address => bool) private isValidator;
    event NewTokenDeployed(address indexed tokenContract);
    event TokenUnlock(address indexed from, address indexed sourceTokenAddress, uint256 amount);
    event TokenLock(address indexed from, address indexed sourceTokenAddress, uint256 amount);

    constructor (address[] memory _validators){
        for (uint256 i = 0; i < _validators.length; i++) {
            isValidator[_validators[i]]=true;
        }
    }

    function isContract(address _address) internal view returns(bool){
        uint32 size;
        assembly {
            size := extcodesize(_address)
        }
        return (size > 0);    
    }    

    function validate(string memory _transaction, address _token, uint256 _amount, address _receiver, uint8[]memory v, bytes32[]memory r, bytes32[]memory s) internal{
        bytes32 messageDigest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32",keccak256(abi.encodePacked(block.chainid,_transaction,_token,_amount,_receiver))));
        for (uint i = 0; i < v.length; i++) {
            address currentAddress = ecrecover(messageDigest, v[i], r[i], s[i]);
            require(isValidator[ecrecover(messageDigest, v[i], r[i], s[i])], "Wrong signature");
            signaturesForTransaction[_transaction].push(currentAddress);            
        }
        require(signaturesForTransaction[_transaction].length>1, "Need at least 2 signatures");
    }

    function unlockTokens(address _token, uint256 _amount, string memory _transaction, uint8[]memory _v, bytes32[]memory _r, bytes32[]memory _s) public{
        require(!isProccessed[_transaction], "Tokens already unlocked.");
        validate(_transaction, _token, _amount, msg.sender, _v, _r, _s); 
        WrappedToken token = WrappedToken(_token);
        isProccessed[_transaction]=true;
        token.transfer(msg.sender, _amount);
        emit TokenUnlock(msg.sender, _token, _amount);
    }

    function mintTokens(address _nativeTokenAddress, uint256 _amount, string memory _transaction, uint8[]memory _v, bytes32[]memory _r, bytes32[]memory _s) public  {
       require(wrappedTokenContracts[_nativeTokenAddress]!= address(0), "No wrapped token contract");
       require(!isProccessed[_transaction], "Tokens already minted.");
       validate(_transaction, _nativeTokenAddress, _amount, msg.sender, _v, _r, _s); 
       WrappedToken wtoken = WrappedToken(wrappedTokenContracts[_nativeTokenAddress]);
       isProccessed[_transaction]=true;
       wtoken.mintTo(msg.sender, _amount);
       emit TokenUnlock(msg.sender, address(wtoken), _amount);
    }

    function claimTokens(address _token, uint256 _amount, string memory _transaction, uint8[]memory _v, bytes32[]memory _r, bytes32[]memory _s) public{
        if(!isContract(_token) || WrappedToken(_token).balanceOf(address(this))==0){
            mintTokens(_token, _amount, _transaction, _v, _r, _s);
        }else{
            unlockTokens(_token, _amount, _transaction, _v, _r, _s);
        }
    }

    function deployContract(address _nativeTokenAddress, string memory _name, string memory _symbol) public{
        require(wrappedTokenContracts[_nativeTokenAddress] == address(0),"Contract already exists");
        address newTokenAddress = address(new WrappedToken(_name, _symbol));
        wrappedTokenContracts[_nativeTokenAddress] = newTokenAddress;
        nativeTokenContracts[newTokenAddress] = _nativeTokenAddress;
        emit NewTokenDeployed(wrappedTokenContracts[_nativeTokenAddress]);
    }
    
    function lock(address _owner, address _token, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) public payable{
        WrappedToken wt = WrappedToken(_token);
        wt.permit(_owner, address(this), _amount, _deadline, _v, _r, _s);
        wt.transferFrom(_owner, address(this), _amount);
        emit TokenLock(_owner, _token, _amount);
    } 

    function burn(address _owner, address _token, uint256 _amount, uint8 _v, bytes32 _r, bytes32 _s) public payable {
        bytes32 messageDigest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32",keccak256(abi.encodePacked(_owner, address(this),_token,_amount))));
        require(_owner == ecrecover(messageDigest, _v, _r, _s), "Wrong Signature");            
        WrappedToken wtoken = WrappedToken(_token);
        wtoken.burnFrom(_owner, _amount);
        //Emits native token address from the other chain, so the other chain knows which token to unlock
        emit TokenLock(_owner, nativeTokenContracts[_token], _amount); 
    }
}
