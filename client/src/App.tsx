import "bootstrap/dist/css/bootstrap.min.css";
import { Container, Row, Col } from "react-bootstrap";
import { useEffect, useState } from "react";
import { properties } from "./properties.js";
import { ethers } from "ethers";
import bridgeAbi from "./contracts/Bridge.json";
import wtAbi from "./contracts/WrappedToken.json";
import { History } from "./components/History";
import { Footer } from "./components/Footer";
import { Bridge } from "./components/Bridge";
import { Header } from "./components/Header";
import { getHistory, validateDeposit, validateClaim } from "./validator/ValidatorClient";
import { burnTokens, lockTokens, claimTokens } from "./bridge/BridgeClient";

function App() {
  const [waiting, setWaiting] = useState(false);
  const [buttonText, setButtonText] = useState("Connect");
  const [walletAddress, setWalletAddress] = useState("");
  const [provider, setProvider] = useState<any>(new ethers.providers.Web3Provider(window.ethereum));
  const [chainId, setChainId] = useState("");
  const [bridgeContract, setBridgeContract] = useState(new ethers.Contract("", bridgeAbi.abi, provider));
  const [testTokenContract, setTestTokenContract] = useState(new ethers.Contract("", wtAbi.abi, provider));
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  interface Transaction {
    lockTransactionHash: string;
    claimTransactionHash: string;
    signatures: string[];
    tokenAddress: string;
    symbol: string;
    amount: string;
    recipient: string;
    fromChainId: string;
    toChainId: string;
    date: number;
    claimed: boolean;
  }

  async function requestAccount() {
    if (window.ethereum) {
      try {
        const accounts = await window.ethereum.request({
          method: "eth_requestAccounts",
        });
        setButtonText("Connected");
        setWalletAddress(accounts[0]);
      } catch (error) {
        console.log("Error connecting...");
      }
    } else {
      alert("Meta Mask not detected");
    }
  }

  // Create a provider to interact with a smart contract
  const connectWallet = async () => {
    if (typeof window.ethereum !== "undefined") {
      await requestAccount();
      setProvider(new ethers.providers.Web3Provider(window.ethereum));
    }
    if (walletAddress) {
      setWalletAddress("");
      window.location.reload();
    }
  };

  const switchNetwork = async () => {
    if (chainId === "4")
      await window.ethereum.request({
        method: "wallet_switchEthereumChain",
        params: [{ chainId: "0x" + "3" }],
      });
    else {
      await window.ethereum.request({
        method: "wallet_switchEthereumChain",
        params: [{ chainId: "0x" + "4" }],
      });
    }
  };

  useEffect(() => {
    window.ethereum.on("chainChanged", () => {
      setProvider(new ethers.providers.Web3Provider(window.ethereum));
      console.log("network changed");
    });
  }, []);

  useEffect(() => {
    provider.getNetwork().then((chain) => {
      console.log("netowrk id : " + chain.chainId);
      setChainId(JSON.stringify(chain.chainId));
    });
  }, [provider]);

  useEffect(() => {
    if (chainId === "3") {
      setTestTokenContract(new ethers.Contract(properties.TEST_TOKEN_ADDRESS_ROPSTEN, wtAbi.abi, provider));
      setBridgeContract(new ethers.Contract(properties.BRIDGE_ADDRESS_ROPSTEN, bridgeAbi.abi, provider));
    } else if (chainId === "4") {
      setTestTokenContract(new ethers.Contract(properties.TEST_TOKEN_ADDRESS_RINKEBY, wtAbi.abi, provider));
      setBridgeContract(new ethers.Contract(properties.BRIDGE_ADDRESS_RINKEBY, bridgeAbi.abi, provider));
    } else {
      setTestTokenContract(new ethers.Contract("", wtAbi.abi, provider));
      setBridgeContract(new ethers.Contract("", bridgeAbi.abi, provider));
    }
  }, [chainId]);

  useEffect(() => {
    window.ethereum.on("accountsChanged", function (accounts) {
      setWalletAddress(accounts[0]);
      console.log(walletAddress);
    });
  }, []);

  useEffect(() => {
    updateHistory();
  }, [walletAddress]);

  const updateHistory = () => {
    if (walletAddress) {
      getHistory(walletAddress).then((data) => {
        setTransactions(data);
      });
    }
  };

  const handleBridge = async (e, tokenToBridge, amount) => {
    console.log("handle bridge")
    console.log(tokenToBridge)
    e.preventDefault();
    let token = new ethers.Contract(tokenToBridge, wtAbi.abi, provider);
    if ((await token.owner()) == bridgeContract.address) {
      validateLockOrBurn(tokenToBridge, await burnTokens(provider, bridgeContract, token, amount));
    } else {
      validateLockOrBurn(tokenToBridge, await lockTokens(provider, bridgeContract, token, amount));
    }
  };

  const handleTokenClaim = async (txHash: string) => {
    let lockTransaction = transactions.find((i) => i.lockTransactionHash === txHash);
    if (lockTransaction !== undefined) {
      await validateClaim(await claimTokens(provider, bridgeContract, lockTransaction));
      updateHistory();
    }
  };

  const validateLockOrBurn = async (tokenToBridge, receipt) => {
    let toChainId;
    if (chainId == "3") {
      toChainId = "4";
    } else {
      toChainId = "3";
    }
    let wtc = new ethers.Contract(tokenToBridge, wtAbi.abi, provider);
    let symbol = await wtc.symbol();
    validateDeposit(receipt.hash, chainId, toChainId, symbol).then(() => {
      updateHistory();
    });
  };

  return (
    <div className="App">
      <Header walletAddress={walletAddress} chainId={chainId} connectWallet={connectWallet} buttonText={buttonText}></Header>
      <Container>
        <Row>
          <Col>
            <Bridge handleBridge={handleBridge}></Bridge>
          </Col>
        </Row>
        <Col>
          <History transactions={transactions} chainId={chainId} handleTokenClaim={handleTokenClaim} switchNetwork={switchNetwork}></History>
        </Col>
      </Container>
      <Footer provider={provider} bridgeContract={bridgeContract} testContract={testTokenContract}></Footer>
    </div>
  );
}

export default App;
