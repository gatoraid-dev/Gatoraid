# Gatoraid
###### (lol, get it?)
[![CircleCI](https://circleci.com/gh/csharpdf/Gatoraid/tree/master.svg?style=svg)](https://circleci.com/gh/csharpdf/Gatoraid/tree/master) [![build](https://github.com/csharpdf/Gatoraid/actions/workflows/build.yml/badge.svg)](https://github.com/csharpdf/Gatoraid/actions/workflows/build.yml)

A simple Text Dungeon game for the 2021 Game Off


## Info

This is my first game made with golang, and as I joined the Game Off a little late, I decided to make a text dungeon while using this game jam as an excuse to learn more game and graphics libraries like SDL, OpenGL, GLFW, etc.. I'm also busy with personal things at the moment, so I don't have much time to put a ton of effort into the Game Off(I'm sorry).

## Building From Source

If you're going to build the game from source, make sure you have the Go 1.17.1 Compiler installed on your pc, then run either `run.bat` or `run.sh` depending on your system.

## Updates

I will post information about what I've done for each development day here.

#### Day 1

I started development, attempted to create a room size generator that worked(though it was not truly random) and got an error in `visual.go` as I finished the day, leaving it for the next.

##### (No image avaliable)
#### Day 2

Today I accomplished lots. I updated the room generator, added user input for the game difficulty, and I now have a clearer concept of how to make infinitely generated rooms.
Accomplishments:
- Updated room size generator(true randomness, difficulties are more unique in terms of room sizes)
- Added user input for difficulty
- Completed Room Map Generator (shows a box with the character in it to represent the current room and player position, player currently doesn't show)

![Image showing difficulty input and slight view of room map](https://cdn.discordapp.com/attachments/666427327437340687/906010154024185896/unknown.png)

#### Day 3

To be honest, I didn't do much today, I really was just polishing the map and trying to think of an efficient way to make the system of infinite rooms. You could call it a brainstorm. I finished Day 3 at about 12:10 am, so you might be able to tell that as I was brainstorming I was also just procrastinating, doing other things.

##### (No image avaliable, same as Day 2)

#### Day 4

I had some personal stuff to attend to, so I couldn't keep the development daily. I also realized that giving the game a visual and calling it "simple" at the same time would be kind of misleading(not that it would hurt though). Because of this, I decided to scrap the visual and room idea and just stay completely text based, confronting one enemy after the other with no end. Now it's starting to sound like a generic dungeon-crawler. I worked a little bit on the enemy randomizer and the different commands you could use(`use [item]` and `proceed`), but overall I didn't complete much today either. Just so I can prove that I wasn't just procrastinating(even though I really was), I started learning the Nim programming language, along with doing some personal stuff. This was written 11/17/21.
