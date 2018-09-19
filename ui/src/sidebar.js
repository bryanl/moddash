import React from 'react';

import Navigation from './navigation';

export default class Sidebar extends React.Component {
  state = {
    results: []
  };

  componentDidMount() {
    fetch("http://localhost:8000/navigation")
      .then(response => response.json())
      .then(response => {
        this.setState({ results: response });
      });
  }

  render() {
    return <Navigation results={this.state.results} />;
  }
}
