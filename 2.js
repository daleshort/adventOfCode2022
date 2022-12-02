const fs = require("fs");
fs.readFile("adventOfCodeDay2.txt", "utf8", (err, data) => {
  if (err) {
    console.error(err);
  }
  //console.log(data);
  const dataArray = data.split("\n");
  scoreGame(dataArray);
});

const getScoreForChoice = (choice) => {
  switch (choice) {
    case "X":
      return 1;
      break;
    case "Y":
      return 2;
      break;
    case "Z":
      return 3;
      break;
    case "A":
      return 1;
      break;
    case "B":
      return 2;
      break;
    case "C":
      return 3;
      break;
  }
};

const getWinner = (round) => {
  const choice1 = getScoreForChoice(round[0]) - 1;
  const choice2 = getScoreForChoice(round[2]) - 1;
  var choices = ["rock", "paper", "scissors"];
  console.log("c1", choice1, "c2", choice2);

  if (choice1 == choice2) {
    return 0;
  }
  if (choice1 == choices.length - 1 && choice2 == 0) {
    return 2;
  }
  if (choice2 == choices.length - 1 && choice1 == 0) {
    return 1;
  }
  if (choice1 > choice2) {
    return 1;
  } else {
    return 2;
  }
};

const scoreRound = (round) => {
  const result = getWinner(round);
  console.log("result", result);
  let player1Points = getScoreForChoice(round[0]);
  let player2Points = getScoreForChoice(round[2]);

  if (result == 0) {
    player1Points += 3;
    player2Points += 3;
  } else if (result == 1) {
    player1Points += 6;
  } else if (result == 2) {
    player2Points += 6;
  }

  return { player1Points, player2Points };
};

const getWinnerFromLetter = (letter) => {
  switch (letter) {
    case "X":
      return 1;
      break;
    case "Y":
      return 0;
      break;
    case "Z":
      return 2;
      break;
  }
};

const getPlayer2EqFromPlayer1 = (player1) => {
  switch (player1) {
    case "A":
      return "X";
      break;
    case "B":
      return "Y";
      break;
    case "C":
      return "Z";
      break;
  }
};

const getPlayForResult = (player1Choice, desiredResult) => {
  if (desiredResult == 0) {
    return getPlayer2EqFromPlayer1(player1Choice);
  }
  if (desiredResult == 1) {
    if (player1Choice == "A") {
      return "Z";
    }
    if (player1Choice == "B") {
      return "X";
    }
    if (player1Choice == "C") {
      return "Y";
    }
  }
  if (desiredResult == 2) {
    if (player1Choice == "A") {
      return "Y";
    }
    if (player1Choice == "B") {
      return "Z";
    }
    if (player1Choice == "C") {
      return "X";
    }
  }
};

const scoreGame = (data) => {
  let player1 = 0;
  let player2 = 0;

  data.map((round) => {
    //draw Y
    // lose X (1 wins)
    // win Z (2 wins)
    const player1Choice = round[0];
    const desiredResult = getWinnerFromLetter(round[2]);
    const player2Choice = getPlayForResult(player1Choice, desiredResult);
    const newRound = player1Choice + " " + player2Choice;

    const { player1Points: points1, player2Points: points2 } =
      scoreRound(newRound);
    player1 += points1;
    player2 += points2;

    console.log(`Round ${round} | P1 ${points1} | P2 ${points2}`);
  });
  console.log(`Player 1 ${player1} | Player 2 ${player2}`);
};
