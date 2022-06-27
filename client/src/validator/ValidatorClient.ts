

export const getHistory = async(walletAddress) => {
    let response = await fetch("http://localhost:8080/transactions/address/" + walletAddress, {
        method: "GET",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
      })
       let json = await response.json()
      let data= json.sort((a, b) => {
            return a.date - b.date;
          });
          //setTransactions(getHistorya(walletAddress));
       // console.log(data)
        return data
};

export const validateDeposit = async(txHash, fromChain, toChain, symbol) => {
    await fetch("http://localhost:8080/transactions", {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ lockTransactionHash: txHash, fromChainId: fromChain, toChainId: toChain, symbol: symbol }),
      })
};

export const validateClaim = async(transaction)=>{
    await fetch("http://localhost:8080/transactions/" + transaction.lockTransactionHash, {
      method: "PATCH",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify(transaction),
    })
}