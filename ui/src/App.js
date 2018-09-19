import React, { Component } from 'react';

import Sidebar from './sidebar';

class App extends Component {
  render() {
    return (
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
            <div className="thirteen wide column">content</div>
          </div>
        </div>
      </div>
    );
  }
}
export default App;
