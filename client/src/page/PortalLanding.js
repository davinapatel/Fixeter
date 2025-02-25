import React from 'react';
import {Container, Card, CardGroup, Button} from "react-bootstrap";
import {Link} from "react-router-dom";
import trackIssueIcon from "../images/track-issue-icon.png";
import reportIssueIcon from "../images/report-issue-icon.png";

const PortalLanding = () => {

    return (
        <Container>
            <h3 style={{ textAlign: 'center'}} className="mb-5">Welcome back to the Fixeter Portal!</h3>
            <div className="text-center">
                <CardGroup className="w-50 mx-auto small">
                    <Card style={{ width: '18rem '}}>
                        <Card.Img variant="top" src={trackIssueIcon} />
                        <Card.Body>
                            <Card.Title>Track your Issues</Card.Title>
                            <Card.Text>
                                View your past and current Issues logged.
                            </Card.Text>
                            <Link to="/track-issues">
                                <Button className="pimary">Track</Button>
                            </Link>
                        </Card.Body>
                    </Card>
                    <Card style={{ width: '18rem '}}>
                        <Card.Img variant="top" src={reportIssueIcon} />
                        <Card.Body>
                            <Card.Title >Report an Issue</Card.Title>
                            <Card.Text>
                                Report a new Issue in your local community.
                            </Card.Text>
                            <Link to ="/log-issue">
                                <Button className="pimary">Report</Button>
                            </Link>
                        </Card.Body>
                    </Card>
                </CardGroup>
            </div>
        </Container>
    );
};

export default PortalLanding;