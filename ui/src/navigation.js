import React from 'react';
import { NavLink } from 'react-router-dom';
import { Menu } from 'semantic-ui-react';

export default class Navigation extends React.Component {
  render() {
    var results = this.props.results;
    if (results === undefined || results.length === 0) {
      return null;
    }

    const entries = results.map((entry, index) => {
      var subMenu = (
        <Menu.Menu >
          <Navigation results={entry.subs}/>
        </Menu.Menu>
      );

      if (!entry.path.startsWith("/")) {
        entry.path = "/" + entry.path;
      }

      return (
        <Menu.Item key={"item_" + index} >
          <NavLink to={entry.path}>
          {entry.key}
          </NavLink>
          {subMenu}
        </Menu.Item>
      );
    });

    return entries;
  }
}
