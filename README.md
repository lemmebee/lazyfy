# lazyfy

lazyfy is a spotify TUI, it will get featured songs daily and create custom home made playlist for you

![Alt Text](https://github.com/ehabshaaban/lazyfy/blob/main/lazyfy.gif)

## Purpose

- Learn go
- Understand bubbletea
- Find a lazy/fast way to discover new music

## Setup

You need to register a new spotify app https://developer.spotify.com/dashboard/

lazyfy needs the following environment variables

- SPOTIFY_ID=""
- SPOTIFY_SECRET=""
- REDIRECT_URI=http://localhost:8080/callback

SPOTIFY_ID and SPOTIFY_SECRET can be found in spotify app dashboard that you just registered
From 'Edit Settings', Update 'Redirect URIs' with REDIRECT_URI above

## Test

```
cd test
go test -run "^TestFeaturedPlaylists$"
```

## Run

```
./lazyfy
```

## Github Workflow

`pull_request.yml` will run tests on every pull request
`releaser.yml` will make a release for every merge to main

## Note

Contributions are more than welcome
This project is kinda on it's early phases, there's so many features that could be hocked up to lazyfy by using [spotify endpoints](https://github.com/zmb3/spotify)
