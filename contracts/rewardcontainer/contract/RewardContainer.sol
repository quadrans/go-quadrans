pragma solidity ^0.5.4;

contract RewardContainer {
    
    address payable private _owner;
    address private _rewardImplementation;
    uint private _totalRewarded;
    
    modifier isOwner() {
        require(msg.sender == _owner);
        _;
    }
    
    modifier isRewardImplementation() {
        require(_rewardImplementation != address(0));
        require(msg.sender == _rewardImplementation);
        _;
    }
    
    event ChangedOwner(address indexed owner);
    event ChangedRewardImplementation(address indexed rewardImplementation);
    event Sent(address indexed to, uint amount);
    event Mined(uint amount);
    
    constructor() public {
        _owner = msg.sender;
    }
    
    function changeOwner(address payable newOwner) public isOwner {
        _owner = newOwner;
        emit ChangedOwner(newOwner);
    }
    
    function getActualReward() public view returns (uint) {
        return address(this).balance;
    }
    
    function getTotalRewarded() public view returns (uint) {
        return _totalRewarded;
    }
    
    function changeRewardImplementation(address newImpl) public isOwner {
        _rewardImplementation = newImpl;
        emit ChangedRewardImplementation(newImpl);
    }
    
    function send(address payable to, uint amount) public isRewardImplementation {
        require(address(this).balance > 0);
        to.transfer(amount);
        _totalRewarded += amount;
        emit Sent(to, amount);
    }
    
    function () external payable {
        emit Mined(msg.value);
    }
    
    function destroy() public isOwner {
        selfdestruct(_owner);
    }
    
}
