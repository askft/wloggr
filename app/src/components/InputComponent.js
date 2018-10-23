import React, { Component } from "react";

export default class InputComponent extends Component {
  constructor(props) {
    super(props);
  }
  render() {
    let { title, name, type, onChange } = this.props;
    return (
      <label>
        <input
          required
          type={type}
          placeholder={title}
          onChange={e => onChange(name, e)}
        />
      </label>
    );
  }
}
