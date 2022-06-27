//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract WrappedToken is ERC20Permit ,Ownable {
    constructor(string memory _name, string memory _symbol) ERC20Permit(_name) ERC20(_name, _symbol) {
    }
    event newMint(address indexed addr, uint256 indexed amount);

    function mintTo(address _to, uint256 _amount) public onlyOwner{
        _mint(_to, _amount);
        emit newMint(_to, _amount);
    }

    function burnFrom(address _from, uint256 _amount) public onlyOwner{
        _burn(_from, _amount);
    }
}
