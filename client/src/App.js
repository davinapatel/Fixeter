import "./App.css";
import Home from "./page/Home";
import IssueForm from "./page/IssueForm";
import Header from "./components/layout/Header";
import { Routes, Route } from "react-router-dom";

function App() {
  return (
    <>
    <Header />
    <Routes>
      <Route path = "/" element ={<Home />} />
      <Route path = "/log-issue" element={<IssueForm />} />
    </Routes>
    </>
  );
}

export default App;
