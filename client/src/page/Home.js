import React from 'react'
import {Container, Row, Col, Image, Card} from "react-bootstrap"
import homeImage from "../images/fixeter-home.jpg";
import potholeImage from "../images/pothole.png"; // Icon by <a class="link_pro" href="https://freeicons.io/man-and-a-hole-icon-set/danger-dangerous-hole-man-manhole-risk-walking-icon-78173">Gan Khoon Lay</a> on <a href="https://freeicons.io">freeicons.io</a>
import graffitiImage from "../images/graffiti.png";// Icon by <a class="link_pro" href="https://freeicons.io/hip-hop-icon-set-4/graffiti-canspray-wall-paint-artist-icon-798399">ColourByteDesigns</a> on <a href="https://freeicons.io">freeicons.io</a>
import binsImage from "../images/bins.png";//Icon by <a class="link_pro" href="https://freeicons.io/health-care-icon-set-9/dustbin-health-care-bin-delete-garbage-recycle-remove-trash-icon-771562">ColourByteDesigns</a> on <a href="https://freeicons.io">freeicons.io</a>
import streetlightImage from "../images/streetlight.png";// Icon by <a class="link_pro" href="https://freeicons.io/power-and-energy-icon-set-38317/illuminating-road-streetlight-public-street-light-electricity-icon-1557560">Pexelpy</a> on <a href="https://freeicons.io">freeicons.io</a>
const Home = () => {

  const cardData = [
    {
      title: "Potholes",
      image: potholeImage,
      text: "Potholes aren't just annoying, they can be dangerous too if not fixed.Report Potholes you spot as drive along roads or notice when walking",
      altText: "An icon of a figure looking at a pothole."
    },
    {
      title: "Graffiti",
      image: graffitiImage,
      text: "Graffiti on Public Buildings/Properties are not allowed unless permission has been granted by the owner. Report any Graffiti which you think shouldn't be there.",
      altText: "An icon of a spray can being sprayed onto a brick wall."
    },
    {
      title: "Bins",
      image: binsImage,
      text: "No one likes overflowing bins. They smell, attract rodents and are unpleasant overall. Use our website to report bins which have not been emptied.",
      altText: "An icon of a rubbish bin."
    },
    {
      title: "Broken Streetlights",
      image: streetlightImage,
      text: "Report any broken/flashing streetlights to us and we guarantee we will get someone out there the same day to take a look.",
      altText: "An icon of two streetlights on."
    }
  ]

    return (
       <main>
        <Container>
          <Row className="">
            <Col sm={7}>
            <h1 class="font-weigh-light">Fixeter</h1>
              <p class="mt-4">
                Welcome to the Fixeter Home Page!<br></br>
                Fixeter is an application where Exeter Residents in their local community to Chalkstone Council.
              </p>
              <Row xs={1} md={2} className="g-4 my-1">
                <Card>
                  <Card.Body>
                    <Card.Title className="text-center">1024</Card.Title>
                    <Card.Text className="text-center">
                        Total number of issues reported to Fixeter in the last 3 months.
                    </Card.Text>
                    </Card.Body>
                </Card>
                <Card>
                  <Card.Body>
                    <Card.Title className="text-center">95%</Card.Title>
                    <Card.Text className="text-center">
                        Percentage of residents satisfied with Fixeter services and response times to issues they reported.
                    </Card.Text>
                    </Card.Body>
                </Card>
             </Row>
              
            </Col>
            <Col sm={5}> 
              <Image src={homeImage}  fluid rounded alt="An image of a man using his phone to pinpoint his location." />
            </Col>
          </Row>
          <Row>
            <Card className="text-center bg-primary-subtle text-dark fst-italic p-3 my-5 py-4">
              <Card.Body>
                Help us fix the issues in your local community by reporting issues directly to us!
              </Card.Body>
            </Card>
          </Row>
          <Row>
            <h4>Issues You Can Report</h4>
            <Row xs={1} md={2} className="g-4">
              {cardData.map((card) => (
                <Col>
                  <Card className="d-flex flex-row">
                    <Card.Img 
                      variant="left"
                      src={card.image}
                      alt={card.altText}
                      style={{ width: '300px', height: 'auto', objectFit: 'cover', marginRight: '10px' }} 
                      />
                    <Card.Body>
                      <Card.Title>{card.title}</Card.Title>
                      <Card.Text>
                        {card.text}
                      </Card.Text>
                    </Card.Body>
                  </Card>
                </Col>
              ))}
            </Row>
          </Row>
        </Container>
       </main>
    );
};

export default Home