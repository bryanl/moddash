import * as path from 'path-browserify';
import React from 'react';
import { Link } from 'react-router-dom';

export default class Navigation extends React.Component {
  render() {
    var results = this.props.results;
    if (results === undefined || results.length === 0) {
      return null;
    }


    const entries = results.map((result, index) => {
      var currentPath = path.join(this.props.parentPath || '', result.path)
      var subMe = (
        <ul>
          <Navigation results={result.subs} parentPath={currentPath} />
        </ul>
      );

      return (
        <li>
          <Link to={currentPath}>{result.name}</Link>
          {subMe}
        </li>
      );
    });

    return entries;
  }
}
