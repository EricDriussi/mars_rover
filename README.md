# Mars Rover

## Install

Get latest Golang version for your OS [here](https://go.dev/doc/install).

Get lastest Node.js version for your OS [here](https://nodejs.org/en/download).

```sh
git clone git@github.com:EricDriussi/mars_rover.git && cd mars_rover && npm i
```

## Run

### Servers

Run API and web servers (might take a while to init the DB the first time).

```sh
go run .
```

Visit [http://localhost:6969](http://localhost:6969) with your browser or curl the API at [http://localhost:4242/api/](http://localhost:4242/api/) (no docs sry).

### Tests

Run backend tests

```sh
go test ./...
```

Run frontend tests

```sh
npm run test
```

## UI

The "frontend" is served statically and will create a new random game when you first visit [http://localhost:6969/](http://localhost:6969/), as well as each time you press the "New Game" button at the top.

The Rover ID will be printed below so that you can load it later.

You can input a series of Commands (or just one command) into the Movement Commands input field.

Try both to see how it behaves!

Alternatively, you can use the vim or the arrow keys to move around the planet.

I went a bit offroad with this behavior: the Kata asks for the Rover to stop moving once it collides.

This seemed unnatural to me, so after implementing that I added another implementation in which the Rover reports the collision but just skips that command and keeps moving.

This can be easily changed by swapping which Move Action is passed to the relevant Controller.

Orange warnings will appear on collision, while red error messages will show up in case of network and/or server issues (try loading a game with a made up Rover ID).

Black obstacles are "big" while colored ones are "small", notice how a black obstacle will never be generated with a size of one cell.

This makes no difference in practice and is done just to showcase code flexibility (more info below).

## Code

You'll find a bunch of `TODO.LM` and `TODO` within the code: the former are used to explain why things are done in a certain way while the latter are just for me to not lose track of pending tasks.

The "frontend" is under `static/` while the rest of the code is under `src/`

The Kata per se is implemented within `domain/` (less care was taken with the rest of the code), you can omit the other dirs.

Just keep in mind that my Rover does not accept Commands as requested by the Kata.
Rather, Commands are handled at the Application layer and translated to Rover function calls.

The design is (hopefully) easy to follow:

- The Planet has a Size and a set of Obstacles.
- A Rover can only exist within a Planet.
- The Rover creates a Map of the Planet on landing.
- The WrappingCollidingRover has a GPS module that takes care of moving (including wrapping) around the Planet.
- Collision detection is handled by the Rover itself.
- Direction, Coordinate(s), Obstacles and Size are little more that wrapper objects to enclose related logic and/or data.

You'll find a lot of places where there are two versions of the same concept: two types of Planet, two types of Rover two types of Obstacles, etc.

This was done in an effort to test how flexible my code was in case different functionality had to be implemented.

Some of this is hidden behind factories while some is just left there as a viable alternative (this is the case with the Creation and Movement Actions).

Since errors are commonly returned as second values in Go, you will find a bunch of `assert.Nil(t, err)` in seemingly unexpected places in tests.

Keep in mind that this is not always intended to work as an assertion in the classical sense, but rather as a way to ensure that tests **fail** instead of **panicking**.

## Caveats

I ran out of time :(

I thought this was due by end of month, so there are a couple of things that where done in a hurry (I wanted to but didn't add Domain errors, for example).

Go's import/module system kinda sucks.

There are a bunch of instances where I had to deal with namespace collisions (especially while iterating over Obstacles and Coordinates), so naming could be much better.

I'm not good with pointers: I basically let the compiler tell me what to do, they might be used incorrectly here and there.

This was done over a lot of iterations and refactors, so you might encounter oddly named or even unused code left behind from earlier verions.

I spent **a lot** more time on this than I should have.
This was not done on learning hours alone.
