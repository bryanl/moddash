import React from 'react';
import { Table } from 'semantic-ui-react';

export default class DummyContent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  componentDidMount() {
    var contentURL = this.props.match.url;
    if (contentURL === "/") {
      contentURL = "/md-overview";
    }
    fetch("http://localhost:8000/contents" + contentURL)
      .then(response => response.json())
      .then(response => {
        this.setState({ contents: response });
      });
  }

  render() {
    if (this.state.contents === undefined) {
      return null;
    }

    var views = [];

    views = this.state.contents.map((content, index) => {
      switch (content.content_type) {
        case "table":
          return showTable(content, index);
        default:
          return "work in progress for " + content.content_type;
      }
    });

    return <div>{views}</div>;
  }
}

function showTable(content, index) {
  var columns = content.data.columns.map((column, index) => {
    return <Table.HeaderCell key={"cell_" + index}>{column}</Table.HeaderCell>;
  });

  var rows = content.data.rows.map((row, rowIndex) => {
    var cells = row.map((cell, cellIndex) => {
      return <Table.Cell key={"cell_" + cellIndex}>{cell}</Table.Cell>;
    });
    return <Table.Row key={"row_" + rowIndex}>{cells}</Table.Row>;
  });

  var title = content.data.title;

  return (
    <div key={"table_" + index}>
      <h2>{title}</h2>
      <Table striped>
        <Table.Header>
          <Table.Row>{columns}</Table.Row>
        </Table.Header>
        <Table.Body>{rows}</Table.Body>
      </Table>
    </div>
  );
}
