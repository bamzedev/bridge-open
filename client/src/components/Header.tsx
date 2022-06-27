import { Button, Container, Navbar, Nav } from "react-bootstrap";

export function Header(props){
    return(
        <Navbar bg="dark" variant="dark">
        <Container>
          <Navbar.Brand href="#home">EVM Bridge</Navbar.Brand>
          <Nav>
            <Navbar.Text className="m-1">Network: {props.chainId === "3" ? "Ropsten" : "Rinkeby"}</Navbar.Text>
            <Navbar.Text className="m-1">Wallet Address: {props.walletAddress}</Navbar.Text>
            <Button variant="secondary" className="m-1" onClick={props.connectWallet}>
              {props.buttonText}
            </Button>
          </Nav>
        </Container>
      </Navbar>
    )
}