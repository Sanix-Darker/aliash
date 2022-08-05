import { Router } from 'preact-router';

import Home from '../routes/home';
import Profile from '../routes/profile';
import './index.css'


const containerStyle = {
    width: "70em",
    margin: "auto",
    padding: "10px"
}

const App = () => (
	<div id="app">
        <div style={containerStyle} className='container'>
            <Router>
                <Home path="/" />
                <Profile path="/profile/" user="me" />
                <Profile path="/profile/:user" />
            </Router>
        </div>
	</div>
)

export default App;
