import React, { Component } from "react";
import RegisterForm from "./RegisterForm";

class Register extends Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    document.title = "Register on wloggr";
  }

  render() {
    return (
      <div>
        <h2>Register</h2>
        <RegisterForm />
      </div>
    );
  }
}

export default Register;
