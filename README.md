# let's go 🎮

> a little terminal-based story inspired by (un)real events <br>
> enter for your own entertainment <br>
> ⚠️ BEWARE OF DAVE ⚠️

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![Bubble Tea](https://img.shields.io/badge/Bubble%20Tea-Charm-FF69B4)
![Lip Gloss](https://img.shields.io/badge/Lip%20Gloss-Charm-FF69B4)

## ✨ hello ✨

I promise there is no AI here <br>
_(which probably is reflected in the code quality)_

especially for the artistic expression of the storyline <br>
_(nothing can beat authentic human chaos)_

### this project aims to solve several issues I currently face:
1. no reliable way to share my engineering talent with the world
2. recruiters not recognizing once-in-a-lifetime opportunity to hire me
3. not a lot of audience to share silly stories I write
4. too much free time between face masks and my 9-5

## what's in there for you
- configurable UI colors, margins and paddings
- different story branches in JSON-defined story
- consequences of your previous actions

## how to play

clone and then:
```bash
go run .
```
also, the storyline json in main is not really functional yet. you might want to check what is hiding in the burning-branch instead 👀

## context

the idea is that this is small terminal story game with proper UI using BubbleTea and LipGloss libraries,
it comes in pink but if you don't have good taste then (maybe) there is going to be an option to run it in blue instead <br>
this project is going to be a showoff of any mechanics
and systems that will come to my mind, simply wrapped in the concept of a game. that will include both Go programming 
as well as DevOps and platform related solutions. I'm even planning to create a k8s operator for this game.
one might suggest that's unnecessary overengineering and I agree, but this is exactly the point.<br>
each system/solution/module will contain its own architectural documentation, so I can explain the made up scenario and
requirements that contributed to the design.

## roadmap 
- add weather service that will affect Game object (and UI view) so I can fight concurrency demons
- add database to store scenes instead of json (just because)
- skills and skills system to affect gameplay (somehow)
- inventory
- achievements
- blue mode
- packaging the whole thing into Dockerfile and ship to kubernetes
- operator to manage some game definitions
- refactor the whole thing into proper packages
- add proper CICD (including tests) and release in Actions
- actually develop tests and testing strategy
- proper logging and error handling 
- http/s endpoint to get current probability of Dave coming back

## questions that no one asked
<details>
<summary>why terminal and not GUI?</summary>
I wanted to be able to play this in kubernetes while waiting for lokistack to reconcile
</details>
<details>
<summary>is the roadmap real?</summary>
yes! well, its more a todo-list, but I think it still counts
</details>
<details>
<summary>are you ever planning to stop programming?</summary>
you wish
</details>

<p align="center">
  made with ❤️ and soy vanilla premium raw paleo iced matcha latte
</p>

