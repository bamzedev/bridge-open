import { Button, Container, Card, Row, Col, ListGroup, Tabs, Tab } from "react-bootstrap";
import { ethers } from "ethers";

const pendingTransactions = (transactions) => {
  return transactions.filter((tx) => tx.claimed === false);
};

const claimedTransactions = (transactions) => {
  return transactions.filter((tx) => tx.claimed === true);
};

export function History(props) {
  return (
    <Container style={{ display: "flex", width: "77rem" }}>
      <Card className="text-center" bg="dark" text="light" style={{ height: "44rem", width: "77rem" }}>
        <Card.Body>
          <Row>
            <Card.Title>Bridge history</Card.Title>
          </Row>
          <Container>
            <Tabs className="m-3" variant="pills" defaultActiveKey="pending">
              <Tab eventKey="pending" title="Pending">
                <Container style={{ display: "flex", width: "49rem" }}>
                  <ListGroup>
                    {pendingTransactions(props.transactions).map((tx) => (
                      <ListGroup.Item key={tx.lockTransactionHash} id={tx.lockTransactionHash} variant="secondary" style={{ height: "3rem", width: "47rem" }} className="mb-1">
                        <Row>
                          <Col>
                            <h4>{ethers.utils.formatEther(tx.amount).toString() + " " + tx.symbol}</h4>
                          </Col>
                          <Col>
                            <h4>
                              {tx.fromChainId === "4" ? (
                                <a href={"http://rinkeby.etherscan.io/tx/" + tx.lockTransactionHash} target="_blank" rel="noreferrer">
                                  Rinkeby
                                </a>
                              ) : (
                                <a href={"http://ropsten.etherscan.io/tx/" + tx.lockTransactionHash} target="_blank" rel="noreferrer">
                                  Ropsten
                                </a>
                              )}
                            </h4>
                          </Col>
                          <Col>
                            <h3>➔</h3>
                          </Col>
                          <Col>
                            <h4> {tx.toChainId === "4" ? "Rinkeby" : "Ropsten"}</h4>
                          </Col>
                          <Col>
                            {tx.claimed ? (
                              <Button variant="outline-dark" size="sm" disabled>
                                Claimed
                              </Button>
                            ) : (
                              <Button
                                size="sm"
                                onClick={() => {
                                  tx.toChainId === props.chainId ? props.handleTokenClaim(tx.lockTransactionHash) : props.switchNetwork();
                                }}
                              >
                                {tx.toChainId === props.chainId ? "Claim" : "Switch"}
                              </Button>
                            )}
                          </Col>
                        </Row>
                      </ListGroup.Item>
                    ))}
                  </ListGroup>
                </Container>
              </Tab>
              <Tab eventKey="claimed" title="Claimed">
                <Container style={{ display: "flex", width: "49rem" }}>
                  <ListGroup>
                    {claimedTransactions(props.transactions).map((tx) => (
                      <ListGroup.Item key={tx.claimTransactionHash} id={tx.claimTransactionHash} variant="secondary" style={{ height: "3rem", width: "47rem" }} className="mb-1">
                        <Row>
                          <Col>
                            <h4>{ethers.utils.formatEther(tx.amount) + " " + tx.symbol}</h4>
                          </Col>
                          <Col>
                            <h4>
                              {tx.fromChainId === "4" ? (
                                <a href={"http://rinkeby.etherscan.io/tx/" + tx.lockTransactionHash} target="_blank" rel="noreferrer">
                                  Rinkeby
                                </a>
                              ) : (
                                <a href={"http://ropsten.etherscan.io/tx/" + tx.lockTransactionHash} target="_blank" rel="noreferrer">
                                  Ropsten
                                </a>
                              )}
                            </h4>
                          </Col>
                          <Col>
                            <h3>➔</h3>
                          </Col>
                          <Col>
                            <h4>
                              {tx.toChainId === "4" ? (
                                <a href={"http://rinkeby.etherscan.io/tx/" + tx.claimTransactionHash} target="_blank" rel="noreferrer">
                                  Rinkeby
                                </a>
                              ) : (
                                <a href={"http://ropsten.etherscan.io/tx/" + tx.claimTransactionHash} target="_blank" rel="noreferrer">
                                  Ropsten
                                </a>
                              )}
                            </h4>
                          </Col>
                          <Col>
                            <Button variant="outline-dark" size="sm" disabled>
                              Claimed
                            </Button>
                          </Col>
                        </Row>
                      </ListGroup.Item>
                    ))}
                  </ListGroup>
                </Container>
              </Tab>
            </Tabs>
          </Container>
        </Card.Body>
      </Card>
    </Container>
  );
}
