const fs = require("fs");
fs.readFile("day4.txt", "utf8", (err, data) => {
  if (err) {
    console.error(err);
  }
  //console.log(data);
  const dataArray = data.split("\n");
  findOverlap(dataArray);
});

const findOverlap = (data) => {
  let count = 0;
  let linecount = 0;
  data.map((line) => {
    const [left, right] = line.split(",");
    linecount++;
    if (checkForAnyOverlap(left, right)) {
      //console.log("overlap");
      //  console.log(`left:${left} right:${right}`);
      count += 1;
    }
  });
  console.log("count", count);
  console.log(linecount);
};

const checkForOverlap = (left, right) => {
  let [leftLower, leftUpper] = left.split("-");
  let [rightLower, rightUpper] = right.split("-");

  leftLower = parseInt(leftLower);
  rightLower = parseInt(rightLower);
  leftUpper = parseInt(leftUpper);
  rightUpper = parseInt(rightUpper);

  if (rightLower >= leftLower && rightUpper <= leftUpper) {
    return true;
  } else if (leftLower >= rightLower && leftUpper <= rightUpper) {
    return true;
  }
  return false;
};

const checkForAnyOverlap = (left, right) => {
  let [leftLower, leftUpper] = left.split("-");
  let [rightLower, rightUpper] = right.split("-");

  leftLower = parseInt(leftLower);
  rightLower = parseInt(rightLower);
  leftUpper = parseInt(leftUpper);
  rightUpper = parseInt(rightUpper);

  //10-12 6-9
  if (leftUpper >= rightLower && !(leftLower > rightLower)) {
    return true;
  }
  if (rightUpper >= leftLower && !(rightLower > leftLower)) {
    return true;
  }
  return false;
};
