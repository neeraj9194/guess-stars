# guess-stars

A “Guess the Stars” CLI game based on the GitHub API. Allows users to choose language (optionally) and shows trending repositories and allows you to guess the number of stars that repo has. Total 5 rounds. Each round presents you with a new repo. You win a round by guessing the stars within 10% tolerance. You win the game if you get 4 or more correct answers out of 5.


API -> to fetch trending githib repos. (fetch 5 random repo from last 7 days trending) 
 returns: {repoName: xx, repo description: xxx, stars: 1234, author: ABC}

cli game ->  loop (5) -> show repo, ask for the stars, store result -> 



