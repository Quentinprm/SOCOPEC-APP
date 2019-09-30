import React, { Component, PropTypes } from 'react';
import { Table, Alert } from 'react-bootstrap/lib';
import StatusRow from './StatusRow';

class StatusBoard extends Component {
  constructor(props) {
    super(props);
    this.state = { error: '' };
  }

  addError(error) {
    this.setState({ error });
  }

  clearErrors() {
    this.setState({ error: '' });
  }

  makeStatusRows() {
    const { status } = this.props;
    return status.map((stat, i) =>
      <StatusRow
        key={i}
        index={i}
        status={stat}
        addError={this.addError.bind(this)}
        clearErrors={this.clearErrors.bind(this)}
      />,
    );
  }

  makeError() {
    const { error } = this.state;
    if (error) {
      return <Alert bsStyle="danger">{error}</Alert>;
    }

    return <div />;
  }

  render() {
    return (
      <div style={{ marginTop: '10px' }}>
        {this.makeError()}
        <Table striped bordered condensed hover>
          <thead>
            <tr>
              <th>ID</th>
              <th>Value</th>
            </tr>
          </thead>
          <tbody>
            {this.makeStatusRows()}
          </tbody>
        </Table>
      </div>
    );
  }
}

StatusBoard.propTypes = {
  status: PropTypes.arrayOf(PropTypes.object),
};

export default StatusBoard;