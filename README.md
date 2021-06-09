# guess-stars

A “Guess the Stars” CLI game based on the GitHub API. Allows users to choose language (optionally) and shows trending repositories and allows you to guess the number of stars that repo has. 


## Install

To install you can use makefile or build using commands

```
make build
OR
go build src/cli/game.go
```

To run

```
make run
OR
./game   // after building
```


## Usage

The game is simple you, first you have to select a language in which you want to search repositories(*Tip: enter 0 for any language*).

Then you will be shown repo one by one, enter the number of stars you think the repo has. Total 5 rounds. Each round presents you with a new repo. You win a round by guessing the stars within 10% tolerance. 

You win the game if you get 4 or more correct answers out of 5.

Lets play!


