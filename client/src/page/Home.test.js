import React from 'react';
import { render, screen } from '@testing-library/react';
import Home from './Home';
import '@testing-library/jest-dom';

describe("Home Page", () => {

  it("Renders the Welcome Message", () => {
    render(<Home />);
    expect(screen.getByText(/Welcome to the Fixeter Home Page!Fixeter is an application where Exeter Residents in their local community to Chalkstone Council./i)).toBeInTheDocument();
  });

  it("Renders the Fixeter title", () => {
    render(<Home />);
    expect(screen.getByRole('heading', { name: /Fixeter/i })).toBeInTheDocument();
  });

  it("Renders the Statistic cards", () => {
    render(<Home />);
    expect(screen.getByText("1024")).toBeInTheDocument();
    expect(screen.getByText("Total number of issues reported to Fixeter in the last 3 months.")).toBeInTheDocument();
    expect(screen.getByText("95%")).toBeInTheDocument();
    expect(screen.getByText("Percentage of residents satisfied with Fixeter services and response times to issues they reported.")).toBeInTheDocument();
  });

  it("Renders the Issues You Can Report section", () => {
    render(<Home />);
    expect(screen.getByRole("heading", { name: /Issues You Can Report/i })).toBeInTheDocument();
  });

  it("Renders the Issue Cards", () => {
    render(<Home />);
    expect(screen.getByText("Potholes")).toBeInTheDocument();
    expect(screen.getByText("Potholes aren't just annoying, they can be dangerous too if not fixed.Report Potholes you spot as drive along roads or notice when walking")).toBeInTheDocument();
    expect(screen.getByText("Graffiti")).toBeInTheDocument();
    expect(screen.getByText("Graffiti on Public Buildings/Properties are not allowed unless permission has been granted by the owner. Report any Graffiti which you think shouldn't be there.")).toBeInTheDocument();
    expect(screen.getByText("Bins")).toBeInTheDocument();
    expect(screen.getByText("No one likes overflowing bins. They smell, attract rodents and are unpleasant overall. Use our website to report bins which have not been emptied.")).toBeInTheDocument();
    expect(screen.getByText("Broken Streetlights")).toBeInTheDocument();
    expect(screen.getByText("Report any broken/flashing streetlights to us and we guarantee we will get someone out there the same day to take a look.")).toBeInTheDocument();
  });

  it('Renders the images', () => {
    render(<Home />);
    expect(screen.getByAltText("An image of a man using his phone to pinpoint his location.")).toBeInTheDocument();
    expect(screen.getByAltText("An icon of a figure looking at a pothole.")).toBeInTheDocument();
    expect(screen.getByAltText("An icon of a spray can being sprayed onto a brick wall.")).toBeInTheDocument();
    expect(screen.getByAltText("An icon of a rubbish bin.")).toBeInTheDocument();
    expect(screen.getByAltText("An icon of two streetlights on.")).toBeInTheDocument();
  });

  it("Renders the motto", () => {
    render(<Home />);
    expect(screen.getByText("Help us fix the issues in your local community by reporting issues directly to us!")).toBeInTheDocument();
  });
});