import React from "react";
import { Col, Container, Row, Stack, Image, Nav, NavLink } from "react-bootstrap";
import { Link } from "react-router-dom";
import {
    FaFacebook,
    FaInstagram,
    FaTwitter
} from "react-icons/fa"

const sections = [
    {
        title: "About Us"
    },
    {
        title: "Contact Us"
    },
    {
        title: "Company"
    },
];

const icons = [
    { name:"Facebook", icon: FaFacebook, link: "https://facebook.com/"},
    { name:"Instagram", icon: FaInstagram, link: "https://instagram.com/"},
    { name:"Twitter", icon: FaTwitter, link: "https://twitter.com/"},
];

const Footer = () => {
  return (
    <footer>
        <Container fluid >
            <Row className="bg-primary text-white p-4">
                <Col className="mx-5">
                    <Stack>
                        <h2> Fixeter </h2>
                        <p> Fixing local issues within Exeter </p>
                    </Stack>
                </Col>
                <Col>
                    <Nav clasName="flex-column fs-5">
                        Useful links
                        <NavLink href="#" className="text-white">Home</NavLink>
                        <NavLink href="#" className="text-white">About Us</NavLink>
                        <NavLink href="#" className="text-white">Support Us</NavLink>
                    </Nav>
                </Col>
            </Row>
        </Container>
    </footer>
  );
};

export default Footer;