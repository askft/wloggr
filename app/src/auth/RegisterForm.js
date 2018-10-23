import React, { Component } from "react";
import { withRouter } from "react-router-dom";
import axios from "axios";

import { URL_API } from "../util/Constants";
import InputComponent from "../components/InputComponent";
import ErrorMessage from "./ErrorMessage";

class RegisterForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      credentials: { email: "", password: "" },
      error: ""
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  // TODO
  //  There is no need to pass the value, as we can just
  //  get it through event.target.value. See Wlog -> Set.
  handleChange(value, event) {
    this.setState({ [value]: event.target.value });
  }

  handleSubmit(event) {
    console.log("Data was submitted: ", JSON.stringify(this.state));
    event.preventDefault();

    axios
      .post(URL_API + "/user/signup", this.state)
      .then(res => {
        console.log(res);
        this.props.history.push("/");
        // TOOD say that registration was successful
      })
      .catch(err => {
        console.log(err);
        console.log(err.response);
        this.setState({ error: err.response.data });
      });
  }

  render() {
    return (
      <div className="container">
        <form onSubmit={this.handleSubmit}>
          <InputComponent
            title="Email"
            name="email"
            type="text"
            onChange={this.handleChange}
          />
          <br />
          <InputComponent
            title="Password"
            name="password"
            type="password"
            onChange={this.handleChange}
          />
          <input type="submit" value="Submit" />
        </form>
        <ErrorMessage err={this.state.error} />
      </div>
    );
  }
}

export default withRouter(RegisterForm);
