import React, { Component } from "react";

export const Set = props => {
  const { n, editing, change, edit, done, copy, remove } = this.props;
  const { reps, weight } = this.props.set;
  return (
    <div>
      &nbsp;Set {n + 1}
      &nbsp;
      <input
        title="Reps"
        name="reps"
        placeholder="Reps"
        type="number"
        min="0"
        defaultValue={reps}
        disabled={!editing}
        onChange={e => change(n, e)}
      />
      &nbsp;x&nbsp;
      <input
        title="Weight"
        name="weight"
        placeholder="Weight"
        type="number"
        // and no minimum amount of weight: could be assisted!
        defaultValue={weight}
        disabled={!editing}
        onChange={e => change(n, e)}
      />
      &nbsp;
      <input type="button" value="✎" onClick={e => edit(n, e)} />
      <input type="button" value="✓" onClick={e => done(n, e)} />
      <input type="button" value="↓" onClick={e => copy(n, e)} />
      <input type="button" value="✗" onClick={e => remove(n, e)} />
      {/* <input type="button" value="+" onClick={e => this.handleNew(e)} /> */}
      &nbsp;
    </div>
  );
};

// export default class Set extends Component {
//   constructor(props) {
//     super(props);
//   }

//   render() {
//     const { n, editing, change, edit, done, copy, remove } = this.props;
//     const { reps, weight } = this.props.set;
//     return (
//       <div>
//         &nbsp;Set {n + 1}
//         &nbsp;
//         <input
//           title="Reps"
//           name="reps"
//           placeholder="Reps"
//           type="number"
//           min="0"
//           defaultValue={reps}
//           disabled={!editing}
//           onChange={e => change(n, e)}
//         />
//         &nbsp;x&nbsp;
//         <input
//           title="Weight"
//           name="weight"
//           placeholder="Weight"
//           type="number"
//           // and no minimum amount of weight: could be assisted!
//           defaultValue={weight}
//           disabled={!editing}
//           onChange={e => change(n, e)}
//         />
//         &nbsp;
//         <input type="button" value="✎" onClick={e => edit(n, e)} />
//         <input type="button" value="✓" onClick={e => done(n, e)} />
//         <input type="button" value="↓" onClick={e => copy(n, e)} />
//         <input type="button" value="✗" onClick={e => remove(n, e)} />
//         {/* <input type="button" value="+" onClick={e => this.handleNew(e)} /> */}
//         &nbsp;
//       </div>
//     );
//   }
// }

// <Set
//   key={set.id}
//   set={set}
//   n={i}
//   // editing={this.state.editing[i]}
//   change={this.handleSetChange}
//   edit={this.handleSetEdit}
//   done={this.handleSetDone}
//   copy={this.handleSetCopy}
//   remove={this.handleSetRemove}
// />
