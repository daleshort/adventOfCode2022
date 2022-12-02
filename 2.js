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

const scoreGame = (data) => {
  let player1 = 0;
  let player2 = 0;

  data.map((round) => {
    const { player1Points: points1, player2Points: points2 } =
      scoreRound(round);
    player1 += points1;
    player2 += points2;

    console.log(`Round ${round} | P1 ${points1} | P2 ${points2}`);
  });
  console.log(`Player 1 ${player1} | Player 2 ${player2}`);
};
