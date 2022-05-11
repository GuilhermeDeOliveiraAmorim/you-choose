import Menu from "components/menu";
import Home from "pages/home";
import Score from "pages/score";
import YouChoose from "pages/youchoose";
import List from "pages/youchoose/score/list";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

export default function AppRouter() {
    return (
        <main>
            <Router>
                <Menu />
                <Routes>
                    <Route path="/home" element={<Home />} />
                    <Route path="/lists" element={<List />} />
                    <Route path="/score" element={<Score />} />
                    <Route path="/you-choose" element={<YouChoose />} />
                </Routes>
            </Router>
        </main>
    );
}
