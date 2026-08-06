# Number Guessing Game

A CLI number guessing game in Go. The computer picks a random number between 1 and 100; you guess it within a limited number of chances that depend on the chosen difficulty.

Project page [Number Guessing Game](https://roadmap.sh/projects/number-guessing-game)

## How to run

Requires Go 1.26+ (see `go.mod`).

```bash
# Play the game
go run .

# Or build and run the binary
go build -o number-guessing-game-cli .
./number-guessing-game-cli
```

## How to test

```bash
go test ./...
```

## How to play

1. On start, the game prints a welcome message, the current high scores, and a difficulty menu.
2. Choose a difficulty — it determines how many chances you get:

   | Level | Chances |
   |-------|---------|
   | 1. Easy   | 10 |
   | 2. Medium | 5  |
   | 3. Hard   | 3  |

3. Guess the secret number (1–100). Each wrong guess tells you whether the secret is greater or less.
4. After **3 consecutive wrong guesses**, you get a hint: the narrowed range the secret must be in (e.g. `Hint: the number is between 25 and 50.`).
5. Win by guessing correctly; lose when you run out of chances. The round summary shows your attempts and elapsed time.
6. If you beat the stored best for that difficulty, you get a `New high score!` — scores persist in `highscores.json` (fewest attempts per difficulty).
7. After each round, answer `Play again? (y/n)` to start a new round with a fresh difficulty choice, or exit.

Example round:

```bash
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
High scores (fewest attempts):
  Easy: -   Medium: -   Hard: -
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
Enter your choice: 2
Great! You have selected the Medium difficulty level.
Let's start the game!
Enter your guess: 50
Incorrect! The number is less than 50.
Enter your guess: 25
Incorrect! The number is greater than 25.
Enter your guess: 35
Incorrect! The number is less than 35.
Enter your guess: 30
Congratulations! You guessed the correct number in 4 attempts.
You took 37.2s.
Play again? (y/n): n
```

## Project structure

```
number-guessing-game/
  go.mod
  main.go                  # CLI: menus, input/output, round + play-again loop
  highscores.json          # generated at runtime (fewest attempts per difficulty)
  internal/game/
    game.go                # Difficulty, Round, Guess, hint logic
    game_test.go           # unit tests for core logic
    highscores.go          # high-score load/save/compare
    highscores_test.go     # unit tests for high scores
```
