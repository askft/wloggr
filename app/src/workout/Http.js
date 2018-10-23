import axios from "axios";

import Auth from "../auth/Auth";
import { URL_API } from "../util/Constants";

export default {
  getWorkoutDates() {
    return axios
      .get(URL_API + "/workout/dates", {
        headers: Auth.TokenHeader()
      })
      .then(res => {
        console.log("WorkoutDates:", res.data.workoutDates);
        return res.data;
      })
      .catch(err => {
        console.log(err);
        console.log(err.response);
      });
  },

  createWorkout(that) {
    return axios
      .post(
        URL_API + "/workout",
        { exercises: [] },
        {
          headers: {
            Authorization: "Bearer " + Auth.getToken(),
            "Content-Type": "application/json"
          }
        }
      )
      .then(res => {
        console.log(res.data);
        return res.data;
      })
      .catch(err => {
        console.log(err);
        console.log(err.response);
      });
  },

  getWorkout(wid) {
    return axios
      .get(URL_API + "/workout/" + wid, {
        headers: Auth.TokenHeader()
      })
      .then(res => {
        console.log("Workout", wid, res.data);
        return res.data;
      })
      .catch(err => {
        console.log(err);
        console.log(err.response);
        return err;
      });
  },

  deleteWorkout(wid) {
    return axios
      .delete(URL_API + "/workout/" + wid, {
        headers: {
          Authorization: "Bearer " + Auth.getToken(),
          "Content-Type": "application/json"
        }
      })
      .then(res => {
        console.log(res);
        return res;
      })
      .catch(err => {
        console.log(err);
        console.log(err.response);
        return err;
      });
  },

  updateWorkout(workout, workoutID) {
    return axios
      .put(URL_API + "/workout/" + workoutID, workout, {
        headers: {
          Authorization: "Bearer " + Auth.getToken(),
          "Content-Type": "application/json"
        }
      })
      .then(res => {
        console.log(res);
        return res;
      })
      .catch(err => {
        console.log(err);
        console.log(err.response);
        return err;
      });
  }
};
