export default {
  shallowCopy(arr) {
    return arr.slice(0);
  },

  set(array, i, value) {
    return array.map((element, j) => {
      if (i === j) {
        return value;
      } else {
        return element;
      }
    });
  },

  spliced(arr, index) {
    let a = arr.slice(0);
    a.splice(index, 1);
    return a;
  },

  appended(arr, elem) {
    return [...arr, elem];
  },

  inserted(arr, index, elem) {
    let a = arr.slice(0);
    a.splice(index, 0, elem);
    return a;
  }
};
