import React, { Component } from "react";
import { NavLink, withRouter } from "react-router-dom";

import Auth from "./../auth/Auth";

class PrivateNav extends Component {
  constructor(props) {
    super(props);
  }
  render() {
    return (
      <nav>
        <ul>
          <li>
            <NavLink exact to="/">
              Dashboard
            </NavLink>
          </li>
          .
          <li>
            <NavLink to="/profile">Profile</NavLink>
          </li>
          .
          <li>
            <NavLink to="/wlog">Workouts</NavLink>
          </li>
        </ul>
        <LogoutButton history={this.props.history} />
        <br />
        <br />
      </nav>
    );
  }
}

class LogoutButton extends Component {
  constructor() {
    super();
    this.handleLogout = this.handleLogout.bind(this);
  }
  handleLogout() {
    Auth.removeToken();
    this.props.history.push("/");
  }
  render() {
    return <button onClick={this.handleLogout}>Log out</button>;
  }
}

export default withRouter(PrivateNav);
