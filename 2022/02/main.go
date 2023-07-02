package main

import (
	"bufio"
	"fmt"
	"os"
)


/*

A -> Rock
B -> Paper
C -> Scissors

X -> Rock
Y -> Paper
Z -> Scissors

Game = Many rounds
Total score = Sum of scores for each round
Score of each round = shape_score + round_score
shape_score = match shape_selected
  | Rock -> 1
  | Paper -> 2
  | Scissors -> 3

round_score = match round_outcome
  | Lose -> 0
  | Draw -> 3
  | Win -> 6

*/

func part_one() {
  score := 0
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()

    their := line[0]
    our := line[2]

    if our == 'X' {
      if their == 'A' {
        score += 4
      } else if their == 'B' {
        score += 1
      } else if their == 'C' {
        score += 7
      }
    } else if our == 'Y' {
      if their == 'A' {
        score += 8
      } else if their == 'B' {
        score += 5
      } else if their == 'C' {
        score += 2
      }
    } else if our == 'Z' {
      if their == 'A' {
        score += 3
      } else if their == 'B' {
        score += 9
      } else if their == 'C' {
        score += 6
      }
    }
  }

  fmt.Println(score)
}

/*

A -> Rock
B -> Paper
C -> Scissors

X -> Lose
Y -> Draw
Z -> Win

shape_score = match shape_selected
  | Rock -> 1
  | Paper -> 2
  | Scissors -> 3

round_score = match round_outcome
  | Lose -> 0
  | Draw -> 3
  | Win -> 6

*/

func main() {

  score := 0

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()

    their := line[0]
    outcome := line[2]

    if outcome == 'X' {
      if their == 'A' {
        score += 3
      } else if their == 'B' {
        score += 1
      } else if their == 'C' {
        score += 2
      }
    } else if outcome == 'Y' {
      if their == 'A' {
        score += 4
      } else if their == 'B' {
        score += 5
      } else if their == 'C' {
        score += 6
      }
    } else if outcome == 'Z' {
      if their == 'A' {
        score += 8
      } else if their == 'B' {
        score += 9
      } else if their == 'C' {
        score += 7
      }
    }
  }

  fmt.Println(score)
}
