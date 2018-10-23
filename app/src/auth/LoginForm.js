import React, { Component } from "react";
import axios from "axios";
import { withRouter } from "react-router-dom";

import Auth from "./Auth";
import { URL_API } from "../util/Constants";
import InputComponent from "../components/InputComponent";
import ErrorMessage from "./ErrorMessage";

class LoginForm extends Component {
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
    // Not pretty, but it does the job.
    this.setState({
      ...this.state,
      credentials: {
        ...this.state.credentials,
        [value]: event.target.value
      }
    });
  }

  handleSubmit(event) {
    console.log("Data was submitted: ", JSON.stringify(this.state.credentials));
    event.preventDefault();
    axios
      .post(URL_API + "/user/signin", this.state.credentials)
      .then(res => {
        console.log(res.data);
        Auth.setToken(res.data);
        this.props.history.push("/");
      })
      .catch(err => {
        console.log(err);
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

export default withRouter(LoginForm);
