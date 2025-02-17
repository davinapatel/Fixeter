import React from 'react'
import {Container, Row,} from "react-bootstrap"
import { Link } from 'react-router-dom';

const Home = () => {

    return (
        <Container className="py-2">
        <Row>
          <h3>
            <Link to ="/log-issue">
              <button className="btn">Log an Issue</button>
            </Link>
          </h3>
        </Row>
      </Container>
    );
};

export default Home