import React from 'react';
import Home from './components/Home';
import Users from './components/Users';
import Microposts from './components/Microposts';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/users" element={<Users />} />
        <Route path="/microposts" element={<Microposts />} />
      </Routes>
    </Router>
  );
}

export default App;
