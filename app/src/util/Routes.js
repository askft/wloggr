import React, { Component } from "react";
import { Route, Redirect } from "react-router-dom";
import Auth from "../auth/Auth";

export class PublicRoute extends Component {
  constructor(props) {
    super(props);
  }
  render() {
    const { component, ...rest } = this.props;
    return (
      <Route
        {...rest}
        render={props =>
          !Auth.isLoggedIn() ? (
            React.createElement(component, props)
          ) : (
            <Redirect
              to={{
                pathname: "/",
                state: { from: props.location }
              }}
            />
          )
        }
      />
    );
  }
}

export class PrivateRoute extends Component {
  constructor(props) {
    super(props);
  }
  render() {
    const { component, ...rest } = this.props;
    return (
      <Route
        {...rest}
        render={props =>
          Auth.isLoggedIn() ? (
            React.createElement(component, props)
          ) : (
            <Redirect
              to={{
                pathname: "/login",
                state: { from: props.location }
              }}
            />
          )
        }
      />
    );
  }
}
