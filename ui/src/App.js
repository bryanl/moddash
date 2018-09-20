import React, { Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { Menu } from 'semantic-ui-react';

import DummyContent from './content';
import Navigation from './navigation';


class App extends Component {
  state = {
    navigationEntries: []
  };

  componentDidMount() {
    fetch("http://localhost:8000/navigation")
      .then(response => response.json())
      .then(response => {
        this.setState({ navigationEntries: response });
      });
  }

  entryPath(entry) {
    return "/" + entry.path;
  }

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
                <Menu vertical size="large" fluid={true}>
                  <Navigation results={this.state.navigationEntries} />
                </Menu>
              </div>
              <div className="twelve wide column">
                {this.state.navigationEntries.map((entry, index) => {
                  var routes = [];
                  routes.push(
                    <Route
                      key={entry.key + index}
                      path={this.entryPath(entry)}
                      component={DummyContent}
                      exact
                    />
                  );

                  entry.subs.map((sub, index) =>
                    routes.push(
                      <Route
                        key={entry.key + "/" + sub.key + index}
                        path={sub.path}
                        component={DummyContent}
                      />
                    )
                  );
                  return routes.flat();
                })}
              </div>
            </div>
          </div>
        </div>
      </Router>
    );
  }
}
export default App;
