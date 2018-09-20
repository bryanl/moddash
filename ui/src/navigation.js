import * as path from 'path-browserify';
import React from 'react';
import { Menu } from 'semantic-ui-react';


export default class Navigation extends React.Component {
  render() {
    var results = this.props.results;
    if (results === undefined || results.length === 0) {
      return null;
    }


    const entries = results.map((result, index) => {
      var currentPath = path.join(this.props.parentPath || '', result.path)
      var subMe = (
        <Menu.Menu>
          <Navigation results={result.subs} parentPath={currentPath} />
        </Menu.Menu>
      );

      return (
        <Menu.Item href={currentPath}>
          {result.name}
          {subMe}
        </Menu.Item>
      );
    });

    return entries;
  }
}
