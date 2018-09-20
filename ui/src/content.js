import React from 'react';
import { Table } from 'semantic-ui-react';

export default class DummyContent extends React.Component {
  constructor(props) {
    super(props);
    console.log(props);
    this.state = {};
  }

  componentDidMount() {
    fetch("http://localhost:8000/contents" + this.props.match.url)
      .then(response => response.json())
      .then(response => {
        this.setState({ contents: response });
      });
  }

  handleContent(content) {}

  render() {
    if (this.state.contents === undefined) {
      return null;
    }

    var views = [];

    views = this.state.contents.map(content => {
      switch (content.content_type) {
        case "table":
          var columns = content.data.columns.map(column => {
            return <Table.HeaderCell>{column}</Table.HeaderCell>;
          });

          var rows = content.data.rows.map(row => {
            var cells = row.map(cell => {
              return <Table.Cell>{cell}</Table.Cell>;
            });
            return <Table.Row>{cells}</Table.Row>;
          });

          var title = content.data.title;

          return (
            <div>
              <h2>{title}</h2>
              <Table striped>
                <Table.Header>
                  <Table.Row>{columns}</Table.Row>
                </Table.Header>
                <Table.Body>{rows}</Table.Body>
              </Table>
            </div>
          );
        default:
          return "work in progress for " + content.content_type;
      }
    });

    return <div>{views}</div>;
  }
}
