import React, { Component } from "react";
import axios from "axios";

import Auth from "../auth/Auth";
import { URL_API } from "../util/Constants";

export default class Profile extends Component {
  constructor(props) {
    super(props);
    this.state = {
      user: {},
      fullName: ""
    };
    this.handleFullNameChange = this.handleFullNameChange.bind(this);
    this.handleFullNameSubmit = this.handleFullNameSubmit.bind(this);
  }

  handleFullNameChange(e) {
    this.setState({ fullName: e.target.value });
  }

  handleFullNameSubmit(e) {
    // e.preventDefault();
    console.log(this.state.fullName);
    axios
      .put(
        URL_API + "/user/profile/fullname",
        { fullName: this.state.fullName },
        {
          headers: {
            "Content-Type": "application/json",
            Authorization: "Bearer " + Auth.getToken()
          }
        }
      )
      .then(res => {
        console.log(res);
        axios
          .get(URL_API + "/user/profile", {
            headers: {
              Authorization: "Bearer " + Auth.getToken()
            }
          })
          .then(res => {
            this.setState({ user: res.data });
          })
          .catch(err => {
            console.log(err);
          });
      })
      .catch(err => {
        console.log(err);
        console.log(err.response);
      });
  }

  componentDidMount() {
    document.title = "Your wloggr profile";
    axios
      .get(URL_API + "/user/profile", {
        headers: {
          Authorization: "Bearer " + Auth.getToken()
        }
      })
      .then(res => {
        console.log(res);
        this.setState({ user: res.data });
      })
      .catch(err => {
        console.log(err);
        console.log(err.response);
      });
  }

  render() {
    return (
      <div>
        <h2>Your profile</h2>
        Hello, user. This is your profile page. You are:
        <br />
        <table border="1px thin sold">
          <tr>
            <td>Email</td>
            <td>Full name</td>
          </tr>
          <tr>
            <td>{this.state.user.email}</td>
            <td>{this.state.user.fullName}</td>
          </tr>
        </table>
        {/* <pre>{JSON.stringify(this.state.user.email, null, 2)}</pre> */}
        <br />
        <input
          type="text"
          placeholder="Full name"
          onChange={this.handleFullNameChange}
        />
        <input
          type="button"
          value="Submit"
          onClick={this.handleFullNameSubmit}
        />
      </div>
    );
  }
}
