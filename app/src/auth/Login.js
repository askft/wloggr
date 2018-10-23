import React, { Component } from "react";
import LoginForm from "./LoginForm";

class Login extends Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    document.title = "Log in to wloggr";
  }

  render() {
    return (
      <div>
        <h2>Log in</h2>
        <LoginForm />
      </div>
    );
  }
}

export default Login;
