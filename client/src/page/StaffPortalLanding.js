import React from 'react';
import {Container, Card, CardGroup, Button} from "react-bootstrap";
import {Link} from "react-router-dom";
import analyticsIcon from "../images/analytics-icon.png";
import manageIssueIcon from "../images/manage-issue-icon.png";

const StaffPortalLanding = () => {

    return (
        <Container>
            <h3 style={{ textAlign: 'center'}} className="mb-5">Welcome back to the Staff Portal!</h3>
            <div className="text-center">
                <CardGroup className="w-50 mx-auto small">
                    <Card style={{ width: '18rem '}}>
                        <Card.Img variant="top" src={manageIssueIcon} alt="Clipboard with writing on it and pencil next to it." />
                        <Card.Body>
                            <Card.Title>Manage Issues</Card.Title>
                            <Card.Text>
                                View and Manage Issues logged by Local Residents.
                            </Card.Text>
                            <Link to="/manage-issue">
                                <Button className="pimary">Manage</Button>
                            </Link>
                        </Card.Body>
                    </Card>
                    <Card style={{ width: '18rem '}}>
                        <Card.Img variant="top" src={analyticsIcon} alt="An icon of a graph." />
                        <Card.Body>
                            <Card.Title >Analytics</Card.Title>
                            <Card.Text>
                                View analyzed reports on Issues in community.
                            </Card.Text>
                            <Link to ="/analytics">
                                <Button className="pimary">Analytics</Button>
                            </Link>
                        </Card.Body>
                    </Card>
                </CardGroup>
            </div>
        </Container>
    );
};

export default StaffPortalLanding;