import { createRoot } from "react-dom/client";
import { App } from "./App";
import './style.css';
import { BrowserRouter } from "react-router-dom";

const container = document.getElementById("app");
const root = createRoot(container)
root.render(<BrowserRouter><App/></BrowserRouter>);