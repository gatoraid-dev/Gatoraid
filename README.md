# Gatoraid
###### (lol, get it?)
[![CircleCI](https://circleci.com/gh/csharpdf/Gatoraid/tree/master.svg?style=svg)](https://circleci.com/gh/csharpdf/Gatoraid/tree/master)

A simple Text Dungeon game for the 2021 Game Off


## Info

This is my first game made with golang, and as I joined the Game Off a little late, I decided to make a text dungeon while using this game jam as an excuse to learn more game and graphics libraries like SDL, OpenGL, GLFW, etc.. I'm also busy with personal things at the moment, so I don't have much time to put a ton of effort into the Game Off(I'm sorry).

## Building From Source

If you're going to build the game from source, make sure you have the Go 1.17.1 Compiler installed on your pc, then run either `run.bat` or `run.sh` depending on your system(sorry macOS I don't like you).

## Updates

I will post information about what I've done for each development day here.

#### Day 1

I started development, attempted to create a room size generator that worked(though it was not truly random) and got an error in `visual.go` as I finished the day, leaving it for the next.
(No image avaliable)
#### Day 2

Today I accomplished lots. I updated the room generator, added user input for the game difficulty, and I now have a clearer concept of how to make infinitely generated rooms.
Accomplishments:
- Updated room size generator(true randomness, difficulties are more unique)
- Added user input for difficulty
- Completed Room Map Generator (shows a box with the character in it to represent the current room and player position, player currently doesn't show)

![Image showing difficulty input and slight view of room map](https://cdn.discordapp.com/attachments/666427327437340687/906010154024185896/unknown.png)
