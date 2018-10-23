import React, { Component } from "react";
import { NavLink } from "react-router-dom";

class PublicNav extends Component {
  constructor(props) {
    super(props);
  }
  render() {
    return (
      <nav>
        <ul>
          <li>
            <NavLink to="/login">Log in</NavLink>
          </li>
          .
          <li>
            <NavLink to="/register">Register</NavLink>
          </li>
        </ul>
      </nav>
    );
  }
}

export default PublicNav;
