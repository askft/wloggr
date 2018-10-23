// Thank you https://github.com/robinskafte

const TOKEN_KEY = "token";

export default {
  TokenHeader() {
    return {
      Authorization: "Bearer " + localStorage.getItem(TOKEN_KEY)
    };
  },

  getToken() {
    return localStorage.getItem(TOKEN_KEY);
  },

  setToken(token) {
    if (typeof token !== "undefined") {
      localStorage.setItem(TOKEN_KEY, token);
    } else {
      console.log("error: token is undefined");
    }
  },

  removeToken() {
    return localStorage.removeItem(TOKEN_KEY);
  },

  isLoggedIn() {
    const token = this.getToken();
    let payload;

    if (token) {
      [, payload] = token.split(".");
      payload = window.atob(payload);
      payload = JSON.parse(payload);

      return payload.exp > Date.now() / 1000;
    }
    return false;
  }

  // getUser() {
  //   if (this.isLoggedIn()) {
  //     return "user";
  //   }
  // }
};
