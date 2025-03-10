import React from "react";
import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import StaffPortalLanding from "./StaffPortalLanding";
import "@testing-library/jest-dom";

describe("Staff Portal Landing Page", () => {

    it("Renders the Welcome Message", () => {
        render(
            <MemoryRouter>
                <StaffPortalLanding />
            </MemoryRouter>
        );
        expect(screen.getByText(/Welcome back to the Staff Portal!/i))
    });

    it("Renders the Manage Issue and Analytics cards", () => {
        render(
            <MemoryRouter>
                <StaffPortalLanding />
            </MemoryRouter>
        );
        expect(screen.getByText("Manage Issues"))
        expect(screen.getByText("View and Manage Issues logged by Local Residents."))
        expect(screen.getAllByText("Analytics"))
        expect(screen.getByText("View analyzed reports on Issues in community."))
    })

    it("Renders the Manage Button and link", () => {
        render(
            <MemoryRouter>
                <StaffPortalLanding />
            </MemoryRouter>
        );

        const manageButton = screen.getByRole("button", { name: /Manage/i});
        expect(manageButton).toBeInTheDocument();

        const link = screen.getByRole("link", { name: /Manage/i});
        expect(link).toHaveAttribute("href", "/manage-issue")
    })

    it("Renders the Analytics Button and link", () => {
        render(
            <MemoryRouter>
                <StaffPortalLanding />
            </MemoryRouter>
        );

        const analyticsButton = screen.getByRole("button", { name: /Analytics/i});
        expect(analyticsButton).toBeInTheDocument();

        const link = screen.getByRole("link", { name: /Analytics/i});
        expect(link).toHaveAttribute("href", "/analytics")
    })

    it("Renders both images", () => {
        render(
            <MemoryRouter>
                <StaffPortalLanding />
            </MemoryRouter>
        );

        const manageIssueImg = screen.getByAltText(/Clipboard with writing on it and pencil next to it./i); // Alt text based on the image's source or description
        const analyticsImg = screen.getByAltText(/An icon of a graph./i);

        expect(manageIssueImg).toBeInTheDocument();
        expect(analyticsImg).toBeInTheDocument();
    })
})