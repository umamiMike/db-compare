# Db-Compare WIP
I am currently refactoring and restructuring the app, so I wouldnt consider it
a simple matter to get up and running.

A quick and dirty tool which allows you to query multiple databases in the same view.


## Why?

I am making this mostly as a code sample.  Much of my work has been for software that is not "Open" or public on github, gitlab, or the like.  I have mostly worked on bigger more enterprise level things.


I want to demonstrate my ability to create an application, both server
and client, from the ground up using skills I have developed  over the last
few years.

the DB compare is a tool to let you compare and query different datastores in the same interface.  In a job I had, I was always having to run sql  commands between seperate instances of our db to see what might be different, or what data might have been corrupted by state issues.


I started this app years ago as a fun little side thing to learn vue.js.  I thought it would be fun to:  

a. refactor something I did a long time ago
b. learn from my past assumptions
c. evaluate how I make do things now with how I did things then, when I had much less experience.

## The Umbrella App

I use [nix](https://nixos.org/) for my development package management as it can
be made immutable, and is based on a function domain specific language (nix)
which I think is neat.

It is very much a work in progress at the moment

1. Install nix ` curl -L https://nixos.org/nix/install | sh `
2. cd into the project directory
3. Run `nix-shell`

## Clients

## React Client

### Vue-Client
A vue.js based js app which is poorly styled but functions.  This was my
original ui.  I am refactoring and formalising this in React.

## Server
Written in [Go](https://golang.org/), uses [Chi](https://github.com/go-chi/chi) for routing

