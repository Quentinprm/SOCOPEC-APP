import React, { Component, PropTypes } from 'react';
import { Button } from 'react-bootstrap/lib';
//import axios from 'axios';
//import Urls from '../util/Urls.js';

class StatusRow extends Component {
  constructor(props) {
    super(props);
    this.state = {
      errors: [],
      isDeleteDisabled: false,
      isDeleteLoading: false,
      isEditDisabled: false,
    };
  }

  resetButtonsState() {
    this.setState({
      isDeleteLoading: false,
      isDeleteDisabled: false,
      isEditDisabled: false,
    });
  }


  render() {
    const { status } = this.props;
    return (
      <tr>
        <td>{status.ID}</td>
        <td>{status.Value}</td>
      </tr>
    );
  }
}

StatusRow.propTypes = {
  status: PropTypes.shape({
    Value: PropTypes.string.isRequired,
    ID: PropTypes.number.isRequired,
  }),
  addError: PropTypes.func.isRequired,
  clearErrors: PropTypes.func.isRequired,
  index: PropTypes.number.isRequired,
};

export default StatusRow;