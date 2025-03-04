import React from "react";
import {Button} from "react-bootstrap";
import { Link } from "react-router-dom";


const Header = () => {
  return (
      <header className="header">
        <div className="logo">Fixeter</div>
        <nav>
          <ul className="nav-links">
            <li>
              <Link to="/">
                <Button>Home</Button>
              </Link>
            </li>
            <li>
              <Link to="/staff-portal">
                <Button>Staff Portal</Button>
              </Link>
            </li>
            <li>
              <Link to="/portal">
                  <Button>Resident Portal</Button>
              </Link>
            </li>
          </ul>
        </nav>
      </header>
  );
};

export default Header;