# javascriptquizgame.wtf

Live at [javascriptquizgame.wtf](https://javascriptquizgame.wtf/) 🥳

## What is this ?
It's just a simple gamified version of the Javascript Questions by [Lydia Hallie](https://github.com/lydiahallie).
You can find her questions [here](https://github.com/lydiahallie/javascript-questions). 

These questions are incredibly fun and brain teasing. Just thought that there should be a gamified version of this 😁

## How does this work ?

There is a backend service written in Go that fetches the content of [this](https://github.com/lydiahallie/javascript-questions/blob/master/en-EN/README.md) readme file, parses it, and stores it in a db as json.

The API serves up json to the frontend when requested. The database is updated once every 24 hours.


## More ideas ?
If you have any ideas to improve this game, feel free to add issues and open PR's ✌🏻

<br></br>

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-javascript.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/you-didnt-ask-for-this.svg)](https://forthebadge.com)

