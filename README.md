# Discord Reddit Reader Bot Written in Golang

A personal project to read Reddit posts on Discord using TextToSpeech voice.

# How to Run

You need to have a working bot first. If you don't have one, here's a simple guide:  
https://discord.com/developers/docs/quick-start/getting-started

* First, you need to populate your `.env` file

    <details>
    <summary>.env example</summary>
  
    ```
    GO_REDDIT_CLIENT_ID=""
    GO_REDDIT_CLIENT_SECRET=""
    GO_REDDIT_CLIENT_USERNAME=""
    GO_REDDIT_CLIENT_PASSWORD=""
    DISCORD_BOT_SECRET="Bot your_secret"
    ```
    </details>

* Then you can run it by using a simple Makefile: `make run`

# How to Use

With the bot active on Discord, you just write this command:

`!read`

Available arguments for the `!read` command:
* `--subreddit` determines from which subreddit you want to read posts.
* `--number` determines the number of posts you want to be read.
* `--timeframe` determines the timeframe you want to include:
    * <details>
         <summary>Possible options</summary>

        * hour
        * day
        * week
        * month
        * year
        * all

        </details>
<details>
    <summary>If none of the arguments are included, the bot will use default values</summary>
    
        --subreddit=golang
        --number=1
        --timeframe=all
</details>

Example: `!read --subreddit=TwoSentenceHorror --number=3 --timeframe=all`

# Motivation

I wanted to listen to Reddit stories with my friends on Discord. We were visiting certain Reddit pages and reading the posts ourselves. I wanted us to listen to stories while we could do something else while spending time on Discord. Therefore, I came up with this idea. Discord TTS would also be an option for reading the posts, but I wanted to challenge myself and try to make something on my own.

I tried to make the project quite extendable, so creating new commands shouldn't be much of a problem. Therefore, I encourage everyone who stumbles upon this project to create their own fork and try adding new features.

# Special Thanks

I want to thank bwmarrin for audio streaming on Discord.  
https://github.com/bwmarrin/dca

