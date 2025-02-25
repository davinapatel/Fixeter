import "./App.css";
import Home from "./page/Home";
import IssueForm from "./page/IssueForm";
import PortalLanding from "./page/PortalLanding";
import TrackIssues from "./page/TrackIssues";
import Header from "./components/layout/Header";
import Footer from "./components/layout/Footer";
import { Routes, Route } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";

function App() {
  return (
    <>
    <div className="d-flex flex-column min-vh-100">
      <Header />
      <main className="flex-grow-1 container py-4">
        <Routes>
          <Route path = "/" element ={<Home />} />
          <Route path = "/log-issue" element={<IssueForm />} />
          <Route path="/portal" element={<PortalLanding />} />
          <Route path="/track-issues" element={<TrackIssues />} />
        </Routes>
      </main>
    </div>
    </>
  );
}

export default App;
