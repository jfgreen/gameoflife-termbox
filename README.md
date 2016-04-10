# Game Of Life - Termbox

![Game Of Life](/img/achim.gif)

Yet another Go implementation of Conway's [Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life).
Mostly an excuse to learn and practise the language but also to test out
[Termbox Go](https://github.com/nsf/termbox-go).

## Running

To start a randomly initiated game:

`go run main.go`


To load and run a particular pattern, specify a path to a valid
[Life 1.6](http://www.conwaylife.com/wiki/Life_1.06) file:

`go run main.go -file pattern.lif`

The [Game Of Life Wiki](http://www.conwaylife.com/wiki/Main_Page) has a collection of 3000+ patterns in this format.

### Command line flags:

```
  -alive string
    	Character to use to render alive cells. (default "●")
  -file string
    	Path of pattern file to initialise game with. Takes precedence over --seed.
  -fps int
    	Frames per second. (default 10)
  -log string
    	Path of logfile to write debugging messages to. (default "/dev/null")
  -seed int
    	Seed to be used in initialisation of random life.
```
### Controls:

- Spacebar to pause/unpause.
- Mouse click to toggle cell (if your terminal supports it).
- R to restart with randomised grid.
- Q or CTRL-C to exit.

## Performance

Can run a bit slow on on large terminals, your millage may vary.
