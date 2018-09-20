import React, { Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';

import DummyContent from './content';
import Sidebar from './sidebar';

class App extends Component {
  render() {
    return (
      <Router>
        <div className="App">
          <div className="ui fixed inverted menu">
            <div className="ui container">
              <a href="/" className="header item">
                moddash
              </a>
            </div>
          </div>
          <div className="ui grid">
            <div className="row">
              <div className="three wide column">
                <Sidebar />
              </div>
              <div className="thirteen wide column">
                <Route path="/:path" component={DummyContent} />
              </div>
            </div>
          </div>
        </div>
      </Router>
    );
  }
}
export default App;
