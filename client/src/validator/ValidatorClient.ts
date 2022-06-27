
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
        return data
};

export const validateDeposit = async(txHash, fromChain, toChain, symbol, name, isBurn) => {
  console.log(txHash)
  console.log(isBurn)
  console.log(fromChain)
  console.log(toChain)
  console.log(symbol)
    await fetch("http://localhost:8080/transactions", {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ lockTransactionHash: txHash, isBurn: isBurn, fromChainId: fromChain, toChainId: toChain, symbol: symbol, name: name }),
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