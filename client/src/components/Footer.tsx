import { Button, Container, Navbar,  Card, Form, Row, Col } from "react-bootstrap";
import { useState } from "react";
import { ethers } from "ethers";

export function Footer(props) {
  const [tokenToDeploy, setTokenToDeploy] = useState("");
  const [tokenToDeployName, setTokenToDeployName] = useState("");
  const [tokenToDeploySymbol, setTokenToDeploySymbol] = useState("");
  const handleTokenDeploySubmit = async (e, provider, bridgeContract) => {
    e.preventDefault();
    console.log(provider);
    console.log(provider.getSigner());
    await bridgeContract.connect(provider.getSigner()).deployContract(tokenToDeploy, tokenToDeployName, tokenToDeploySymbol);
  };
  const handleMintTestTokenSubmit = async (e, provider, testTokenContract) => {
    e.preventDefault();
    await testTokenContract.connect(provider.getSigner()).mintTo(await provider.getSigner().getAddress(), ethers.utils.parseEther("1000"));
  };

  return (
    <Navbar bg="dark" variant="dark" fixed="bottom">
      <h2 style={{ color: "white", marginLeft: 30 }}>Admin / Test</h2>
      <Container fluid="md" style={{ color: "white" }}>
        <div style={{ lineHeight: 3, padding: 2 }}>
          <Row>
            <Col>
              <Card className="text-center" bg="dark" text="light" style={{ width: "14rem", height: "20rem" }}>
                <Card.Body style={{ display: "center" }}>
                  <Card.Title>Deploy Token Contract</Card.Title>
                  <Card.Text style={{ lineHeight: 1 }}>Token contract used for minting wrapped tokens on target chain.</Card.Text>
                </Card.Body>
                <Card.Footer>
                  <Form onSubmit={(e) => handleTokenDeploySubmit(e, props.provider, props.bridgeContract)}>
                    <Form.Label>
                      <Form.Control placeholder="Native Token Address" onChange={(e) => setTokenToDeploy(e.target.value)} />
                      <Form.Control placeholder="Token Name" onChange={(e) => setTokenToDeployName(e.target.value)} />
                      <Form.Control placeholder="Token Symbol" onChange={(e) => setTokenToDeploySymbol(e.target.value)} />
                      <Button type="submit" variant="primary">
                        Deploy
                      </Button>
                    </Form.Label>
                  </Form>
                </Card.Footer>
              </Card>
            </Col>
            <Col>
              <Card className="text-center" bg="dark" text="light" style={{ width: "14rem", height: "20rem" }}>
                <Card.Body style={{ display: "center" }}>
                  <Card.Title>Mint Native Test Tokens</Card.Title>
                  <Card.Text style={{ lineHeight: 1 }}>Receive native tokens for testing</Card.Text>
                </Card.Body>
                <Card.Footer>
                  {props.testContract ? <p>{props.testContract.address}</p> : <p>Chain not supported</p>}
                  <Form onSubmit={(e) => handleMintTestTokenSubmit(e, props.provider, props.testContract)}>
                    <Form.Label>
                      <Button type="submit" variant="primary">
                        Mint
                      </Button>
                    </Form.Label>
                  </Form>
                </Card.Footer>
              </Card>
            </Col>
          </Row>
        </div>
      </Container>
    </Navbar>
  );
}
