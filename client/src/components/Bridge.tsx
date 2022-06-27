import { Button, Container, Card, Form, InputGroup } from "react-bootstrap";
import { useState } from "react";
import { ethers } from "ethers";

export function Bridge(props) {
  const [token, setToken] = useState("");
  const [amount, setAmount] = useState("");
  return (
    <Container className="mb-1" style={{ display: "flex", width: "78.5rem" }}>
      <Card className="text-center" bg="dark" text="light" style={{ width: "78rem" }}>
        <Card.Body style={{ display: "center" }}>
          <Card.Title>Bridge Tokens</Card.Title>
          <Card.Text>Rinkeby ⇆ Ropsten</Card.Text>
        </Card.Body>
        <Form onSubmit={(e)=> props.handleBridge(e, token, amount)}>
          <Form.Label>
            <InputGroup className="mb-3">
              <Form.Control placeholder="Token Address" aria-label="TokenAddress" onChange={(e) => setToken(e.target.value)} />
              <Form.Control placeholder="Token Amount" aria-label="TokenAamount" onChange={(e) => setAmount(ethers.utils.parseEther(e.target.value).toString())} />
              <Button type="submit" variant="primary">
                Bridge
              </Button>
            </InputGroup>
          </Form.Label>
        </Form>
      </Card>
    </Container>
  );
}
