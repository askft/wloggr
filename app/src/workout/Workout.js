import React, { Component } from "react";
import shortid from "shortid";

import Http from "./Http";
import Helpers from "../util/Helpers";

class Workout extends Component {
  constructor(props) {
    super(props);
    this.state = {
      workout: {
        exercises: []
      },
      newExerciseName: ""
    };
    this.handleExerciseNew = this.handleExerciseNew.bind(this);
    this.handleExerciseRemove = this.handleExerciseRemove.bind(this);
    this.handleSetChange = this.handleSetChange.bind(this);
    this.handleSetEdit = this.handleSetEdit.bind(this);
    this.handleSetDone = this.handleSetDone.bind(this);
    this.handleSetCopy = this.handleSetCopy.bind(this);
    this.handleSetRemove = this.handleSetRemove.bind(this);
    this.handleSetNew = this.handleSetNew.bind(this);
    this.handleBlur = this.handleBlur.bind(this);
  }

  componentDidMount() {
    Http.getWorkout(this.props.match.params.wid)
      .then(data => {
        this.setState({ workout: data });
      })
      .catch(err => {
        console.log("Response status: " + err.response.status);
      });
  }

  // componentDidUpdate(prevProps, prevState) {
  //   console.log("componentDidUpdate was called.");
  // }

  // Done.
  handleExerciseNew(event) {
    event.preventDefault();
    let workout = {
      ...this.state.workout,
      exercises: [
        ...this.state.workout.exercises,
        {
          name: this.state.newExerciseName,
          sets: [],
          editing: [],
          id: shortid.generate()
        }
      ]
    };
    this.setState({ workout });
    Http.updateWorkout(workout, this.props.match.params.wid);
    console.log(workout);
  }

  // Must change state inside Wlog -
  handleExerciseRemove(n, event) {
    let workout = {
      ...this.state.workout,
      exercises: Helpers.spliced(this.state.workout.exercises, n)
    };
    this.setState({ workout });
    Http.updateWorkout(workout, this.props.match.params.wid);
    // history.push("/wlog");
  }

  handleSetChange(eIndex, sIndex, event) {
    let set = {
      ...this.state.workout.exercises[eIndex].sets[sIndex],
      [event.target.name]: event.target.value
    };
    let workout = {
      ...this.state.workout,
      exercises: Helpers.set(this.state.workout.exercises, eIndex, {
        ...this.state.workout.exercises[eIndex],
        sets: Helpers.set(
          this.state.workout.exercises[eIndex].sets,
          sIndex,
          set
        )
      })
    };
    this.setState({ workout });
  }

  handleSetEdit(eIndex, sIndex, event) {
    let editing = Helpers.set(
      this.state.workout.exercises[eIndex].editing,
      sIndex,
      true
    );
    let workout = {
      ...this.state.workout,
      exercises: Helpers.set(this.state.workout.exercises, eIndex, {
        ...this.state.workout.exercises[eIndex],
        editing: editing
      })
    };
    this.setState({ workout });
  }

  handleSetDone(eIndex, sIndex, event) {
    let editing = Helpers.set(
      this.state.workout.exercises[eIndex].editing,
      sIndex,
      false
    );
    let workout = {
      ...this.state.workout,
      exercises: Helpers.set(this.state.workout.exercises, eIndex, {
        ...this.state.workout.exercises[eIndex],
        editing: editing
      })
    };
    this.setState({ workout });

    // TODO: Instead of this, update on onBlur.
    Http.updateWorkout(workout, this.props.match.params.wid);
  }

  handleSetCopy(eIndex, sIndex, event) {
    let set = {
      ...this.state.workout.exercises[eIndex].sets[sIndex],
      id: shortid.generate()
    };
    let exercise = {
      ...this.state.workout.exercises[eIndex],
      sets: Helpers.inserted(
        this.state.workout.exercises[eIndex].sets,
        sIndex + 1,
        set
      )
    };
    let workout = {
      ...this.state.workout,
      exercises: Helpers.set(this.state.workout.exercises, eIndex, exercise)
    };
    console.log(workout);
    this.setState({ workout });

    Http.updateWorkout(workout, this.props.match.params.wid);
  }

  handleSetRemove(eIndex, sIndex, event) {
    let workout = {
      ...this.state.workout,
      exercises: Helpers.set(this.state.workout.exercises, eIndex, {
        ...this.state.workout.exercises[eIndex],
        sets: Helpers.spliced(
          this.state.workout.exercises[eIndex].sets,
          sIndex
        ),
        editing: Helpers.spliced(
          this.state.workout.exercises[eIndex].editing,
          sIndex
        )
      })
    };
    console.log(workout);
    this.setState({ workout });
    Http.updateWorkout(workout, this.props.match.params.wid);
  }

  handleSetNew(eIndex, event) {
    let workout = {
      ...this.state.workout,
      exercises: Helpers.set(this.state.workout.exercises, eIndex, {
        ...this.state.workout.exercises[eIndex],
        sets: [
          ...(this.state.workout.exercises[eIndex].sets || []),
          { reps: 0, weight: 0, id: shortid.generate() }
        ],
        editing: [...this.state.workout.exercises[eIndex].editing, true]
      })
    };
    this.setState({ workout });
    Http.updateWorkout(workout, this.props.match.params.wid);
  }

  handleBlur(eIndex, sIndex, event) {
    let set = {
      ...this.state.workout.exercises[eIndex].sets[sIndex],
      [event.target.name]: event.target.value
    };
    let workout = {
      ...this.state.workout,
      exercises: Helpers.set(this.state.workout.exercises, eIndex, {
        ...this.state.workout.exercises[eIndex],
        sets: Helpers.set(
          this.state.workout.exercises[eIndex].sets,
          sIndex,
          set
        )
      })
    };
    Http.updateWorkout(this.state.workout, this.props.match.params.wid)
      .then(res => {
        this.setState({ workout });
      })
      .catch(err => {
        //
      });
  }

  render() {
    let workout = this.state.workout;
    return (
      <div>
        <h3>Workout ({this.props.match.params.wid})</h3>

        <FormExerciseNew
          onSubmit={this.handleExerciseNew}
          onChange={e => {
            this.setState({ newExerciseName: e.target.value });
          }}
        />

        {workout.exercises.map((exercise, eIndex) => (
          <div key={exercise.id}>
            {exercise.name}
            &nbsp;
            {/* {"<" + exercise.id + ">"} */}
            {exercise.sets &&
              exercise.sets.map((set, sIndex) => (
                <div key={set.id}>
                  &nbsp;Set {sIndex + 1}
                  &nbsp;
                  <input
                    title="Reps"
                    name="reps"
                    placeholder="Reps"
                    type="number"
                    min="0"
                    defaultValue={set.reps}
                    // disabled={!editing}
                    onChange={e => this.handleSetChange(eIndex, sIndex, e)}
                    onBlur={e => this.handleBlur(eIndex, sIndex, e)}
                  />
                  &nbsp;x&nbsp;
                  <input
                    title="Weight"
                    name="weight"
                    placeholder="Weight"
                    type="number"
                    // and no minimum amount of weight: could be assisted!
                    defaultValue={set.weight}
                    // disabled={!editing}
                    onChange={e => this.handleSetChange(eIndex, sIndex, e)}
                    onBlur={e => this.handleBlur(eIndex, sIndex, e)}
                  />
                  &nbsp;
                  {/* <input
                    type="button"
                    value="✎"
                    onClick={e => this.handleSetEdit(eIndex, sIndex, e)}
                  />
                  <input
                    type="button"
                    value="✓"
                    onClick={e => this.handleSetDone(eIndex, sIndex, e)}
                  /> */}
                  <input
                    type="button"
                    value="↓"
                    onClick={e => this.handleSetCopy(eIndex, sIndex, e)}
                  />
                  <input
                    type="button"
                    value="✗"
                    onClick={e => this.handleSetRemove(eIndex, sIndex, e)}
                  />
                </div>
              ))}
            &nbsp;
            <input
              type="button"
              value="New set"
              onClick={e => this.handleSetNew(eIndex, e)}
            />
            <input
              type="button"
              value="Remove exercise"
              onClick={e => this.handleExerciseRemove(eIndex, e)}
            />
          </div>
        ))}
        {/* <pre>{JSON.stringify(workout, null, 2)}</pre> */}
      </div>
    );
  }
}

const FormExerciseNew = ({ onSubmit, onChange }) => (
  <form onSubmit={onSubmit}>
    <input
      required
      title="Exercise"
      name="exercise"
      placeholder="Exercise name"
      type="text"
      onChange={onChange}
    />
    <input type="submit" value="New exercise" />
  </form>
);

export default Workout;
