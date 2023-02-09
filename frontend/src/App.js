import React from "react";
import {Wrapper} from "./components/wrapper";
import { Routes, Route } from "react-router-dom";
import {Chapter} from "./components/chapter";
import {NotFound} from "./components/notfound";


export function App() {
    return (
        <div>
            <Routes>
                <Route index element={<Wrapper />} />
                <Route path={"/chapter/:id"} element={<Chapter/>} />
                <Route path='*' element={<NotFound />}/>
            </Routes>
        </div>
    );
}
