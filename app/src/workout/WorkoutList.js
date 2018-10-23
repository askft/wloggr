import React, { Component } from "react";

export default class WorkoutList extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <h2>Workouts</h2>
        {this.props.children}
      </div>
    );
  }
}
