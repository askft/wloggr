import React, { Component } from "react";
import { Link, Route, Redirect, Switch } from "react-router-dom";

import { PublicLayout, PrivateLayout } from "./util/Layouts";
import { PublicRoute, PrivateRoute } from "./util/Routes";
import Login from "./auth/Login";
import Register from "./auth/Register";
import Dashboard from "./components/Dashboard";
import Profile from "./components/Profile";
import Wlog from "./workout/Wlog";

import "./style.css";

const NotFound = () => <div>404 not found</div>;

export default class App extends Component {
  componentDidMount() {
    document.title = "wloggr";
  }
  render() {
    return (
      <div>
        <Switch>
          <PrivateRoute exact path="/" component={PrivateLayout(Dashboard)} />
          <PrivateRoute path="/profile" component={PrivateLayout(Profile)} />
          <PrivateRoute path="/wlog" component={PrivateLayout(Wlog)} />
          <PublicRoute path="/login" component={PublicLayout(Login)} />
          <PublicRoute path="/register" component={PublicLayout(Register)} />
          <Route component={NotFound} />
        </Switch>
      </div>
    );
  }
}
