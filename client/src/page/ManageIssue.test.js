import React from "react";
import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import ManageIssue from "./ManageIssue";
import "@testing-library/jest-dom";

describe("Manage Issue Page", () => {

    it("Renders the title", () => {
        render(
            <MemoryRouter>
                <ManageIssue />
            </MemoryRouter>
        );
        expect(screen.getByText(/Manage Issue/i))
    });

    it("Renders the Logged, In Progress and Closed cards", () => {
        render(
            <MemoryRouter>
                <ManageIssue />
            </MemoryRouter>
        );
        expect(screen.getAllByText("Logged Issues"))
        expect(screen.getByText("Resolve Logged Issues."))
        expect(screen.getAllByText("In Progress Issues"))
        expect(screen.getByText("Review In Progress Issues."))
        expect(screen.getAllByText("Closed Issues"))
        expect(screen.getByText("View recently Closed Issues."))
    })

    it("Renders the buttons and links", () => {
        render(
            <MemoryRouter>
                <ManageIssue />
            </MemoryRouter>
        );

        const logButton = screen.getByRole("button", { name: /Logged Issues/i});
        expect(logButton).toBeInTheDocument();

        const link = screen.getByRole("link", { name: /Logged Issues/i});
        expect(link).toHaveAttribute("href", "/issues/logged")

        const progressButton = screen.getByRole("button", { name: /In Progress Issues/i});
        expect(progressButton).toBeInTheDocument();

        const link2 = screen.getByRole("link", { name: /In Progress Issues/i});
        expect(link2).toHaveAttribute("href", "/issues/progress")

        const closedButton = screen.getByRole("button", { name: /Closed Issues/i});
        expect(closedButton).toBeInTheDocument();

        const link3 = screen.getByRole("link", { name: /Closed Issues/i});
        expect(link3).toHaveAttribute("href", "/issues/closed")
    })

    it("Renders all images", () => {
        render(
            <MemoryRouter>
                <ManageIssue />
            </MemoryRouter>
        );

        const loggedImg = screen.getByAltText(/A clipboard with writing and a pencil next to it./i);
        const progressImg = screen.getByAltText(/A clipboard with writing and a clock next to it./i);
        const closedImg = screen.getByAltText(/A clipboard with ticks on it./i);

        expect(loggedImg).toBeInTheDocument();
        expect(progressImg).toBeInTheDocument();
        expect(closedImg).toBeInTheDocument();
    })
})