import React from 'react'
import {Container, Card, CardGroup, Button} from "react-bootstrap";
import {Link} from "react-router-dom";
import manageIssueIcon from "../images/manage-issue-icon.png";
import inProgressIssueIcon from "../images/in-progress-icon.png";
import closedIssueIcon from "../images/closed-issue-icon.png";


const ManageIssue = () => {

    return (
        <Container>
            <h3 style={{ textAlign: 'center'}} className="mb-5">Manage Issues</h3>
            <div className="text-center">
                <CardGroup className="w-50 mx-auto small">
                    <Card style={{ width: '18rem '}}>
                        <Card.Img variant="top" src={manageIssueIcon} />
                        <Card.Body>
                            <Card.Title>Logged Issues</Card.Title>
                            <Card.Text>
                                Resolve Logged Issues.
                            </Card.Text>
                            <Link to="/issues/logged">
                                <Button className="pimary">Logged Issues</Button>
                            </Link>
                        </Card.Body>
                    </Card>
                    <Card style={{ width: '18rem '}}>
                        <Card.Img variant="top" src={inProgressIssueIcon} />
                        <Card.Body>
                            <Card.Title >In Progress Issues</Card.Title>
                            <Card.Text>
                                Review In Progress Issues.
                            </Card.Text>
                            <Link to ="/issues/progress">
                                <Button className="pimary">In Progress Issues</Button>
                            </Link>
                        </Card.Body>
                    </Card>
                    <Card style={{ width: '18rem '}}>
                        <Card.Img variant="top" src={closedIssueIcon} />
                        <Card.Body>
                            <Card.Title >Closed Issues</Card.Title>
                            <Card.Text>
                                View recently Closed Issues.
                            </Card.Text>
                            <Link to ="/issues/closed">
                                <Button className="pimary">Closed Issues</Button>
                            </Link>
                        </Card.Body>
                    </Card>
                </CardGroup>
            </div>
        </Container>
    );
};

export default ManageIssue;