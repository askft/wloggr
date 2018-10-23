import React, { Component } from "react";
import { Route, Link, Switch, Redirect, withRouter } from "react-router-dom";
import shortid from "shortid";

import history from "../util/history";

import WorkoutList from "./WorkoutList";
import Workout from "./Workout";

import Helpers from "../util/Helpers";
import Http from "./Http";

class Wlog extends Component {
  constructor(props) {
    super(props);
    this.state = {
      dates: [],
      redirect: false
    };
    this.handleNewWorkout = this.handleNewWorkout.bind(this);
    this.handleRemoveWorkout = this.handleRemoveWorkout.bind(this);
  }

  componentDidMount() {
    Http.getWorkoutDates().then(data => {
      this.setState({ dates: data.workoutDates });
    });
  }

  handleNewWorkout(event) {
    Http.createWorkout(this).then(data => {
      this.setState({
        dates: [...this.state.dates, data]
      });
    });
  }

  handleRemoveWorkout(n, event) {
    n = this.state.dates.length - n - 1; // because it's backwards
    if (!confirm("Delete workout " + this.state.dates[n] + "?")) {
      return;
    }
    Http.deleteWorkout(this.state.dates[n])
      .then(res => {
        let dates = Helpers.spliced(this.state.dates, n);
        console.log(dates);
        this.setState({ dates }, () => {
          history.push("/wlog");
          console.log(this.state);
        });
        // this.setState({ redirect: true });
      })
      .catch(err => {
        console.log("error: could not delete");
      });
  }

  render() {
    return (
      <div>
        {/* TODO remove WorkoutList component */}
        <WorkoutList>
          <input
            type="button"
            value="New workout"
            onClick={this.handleNewWorkout}
          />
          {this.state.dates &&
            this.state.dates
              .slice(0)
              .reverse()
              .map((wid, i) => (
                <li key={shortid.generate()}>
                  <Link to={this.props.match.url + "/" + wid}>{wid}</Link>
                  &nbsp;
                  <input
                    type="button"
                    value="Remove workout"
                    onClick={e => this.handleRemoveWorkout(i, e)}
                  />
                  {/* <form onSubmit={e => this.handleRemoveWorkout(i, e)}>
                    <input type="submit" value="Remove workout" />
                  </form> */}
                </li>
              ))}
        </WorkoutList>

        {/* if wid in dates then render workout else redirect */}
        {/* <Route
          path={this.props.match.url + "/:wid"}
          render={() => {
            if (this.state.redirect) {
              // this.setState({ redirect: false });
              return <Workout />;
            } else {
              return <Redirect to="/wlog" />;
            }
          }}
        /> */}

        <Route path={this.props.match.url + "/:wid"} component={Workout} />
      </div>
    );
  }
}

export default Wlog;
