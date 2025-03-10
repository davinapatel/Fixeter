import React from "react";
import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import PortalLanding from "./PortalLanding";
import "@testing-library/jest-dom";

describe("Portal Landing Page", () => {

    it("Renders the Welcome Message", () => {
        render(
            <MemoryRouter>
                <PortalLanding />
            </MemoryRouter>
        );
        expect(screen.getByText(/Welcome back to the Fixeter Portal!/i))
    });

    it("Renders the Track Issue and Report Issue cards", () => {
        render(
            <MemoryRouter>
                <PortalLanding />
            </MemoryRouter>
        );
        expect(screen.getByText("Track your Issues"))
        expect(screen.getByText("View your past and current Issues logged."))
        expect(screen.getByText("Report an Issue"))
        expect(screen.getByText("Report a new Issue in your local community."))
    })

    it("Renders the Report Button and link", () => {
        render(
            <MemoryRouter>
                <PortalLanding />
            </MemoryRouter>
        );

        const reportButton = screen.getByRole("button", { name: /Report/i});
        expect(reportButton).toBeInTheDocument();

        const link = screen.getByRole("link", { name: /Report/i});
        expect(link).toHaveAttribute("href", "/log-issue")
    })

    it("Renders the Track Button and link", () => {
        render(
            <MemoryRouter>
                <PortalLanding />
            </MemoryRouter>
        );

        const trackButton = screen.getByRole("button", { name: /Track/i});
        expect(trackButton).toBeInTheDocument();

        const link = screen.getByRole("link", { name: /Track/i});
        expect(link).toHaveAttribute("href", "/track-issues")
    })

    it("Renders both images", () => {
        render(
            <MemoryRouter>
                <PortalLanding />
            </MemoryRouter>
        );

        const trackImg = screen.getByAltText(/A location marker attached to a paper./i); // Alt text based on the image's source or description
        const reportImg = screen.getByAltText(/A pen on a piece of paper as if it's writing./i);

        expect(trackImg).toBeInTheDocument();
        expect(reportImg).toBeInTheDocument();
    })
})