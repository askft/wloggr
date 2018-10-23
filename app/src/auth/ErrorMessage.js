import React, { Component } from "react";

export default class ErrorMessage extends Component {
  constructor(props) {
    super(props);
  }
  render() {
    return (
      <div className="errorMessage">
        <br />
        {this.props.err && format(this.props.err)}
      </div>
    );
  }
}

function format(string) {
  return string.charAt(0).toUpperCase() + string.slice(1).trim() + ".";
}
