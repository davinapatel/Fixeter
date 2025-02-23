import React from 'react'
import {Container, Row, Button} from "react-bootstrap"
import { Link } from 'react-router-dom';

const Home = () => {

    return (
        <Container className="py-2">
        <Row>
          <h3>
            <Link to ="/log-issue">
              <Button className="btn">Log an Issue</Button>
            </Link>
            <Link to ="/portal">
              <Button className="btn">Portal</Button>
            </Link>
          </h3>
        </Row>
      </Container>
    );
};

export default Home