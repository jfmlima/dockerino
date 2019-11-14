import React from 'react';
import logo from '../logo.svg';
import './App.css';
import Docker from "./features/docker/Docker";

const App: React.FC = () => {
    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo"/>
                <Docker></Docker>
            </header>
        </div>
    );
};

export default App;
