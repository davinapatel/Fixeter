import React from "react";
import { Col, Container, Row } from "react-bootstrap";
import { Link } from "react-router-dom";


const Header = () => {
  return (
      <header className="header">
        <div className="logo">Fixeter</div>
        <nav>
          <ul className="nav-links">
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/log-issue">Log Issue</Link>
            </li>
          </ul>
        </nav>
      </header>
  );
};

export default Header;