# cdown 0.1A
A countdown timer for your terminal, based on figlet.

## Written by
Björn Westerberg Nauclér (mail@bnaucler.se) 2016

## Thanks to
Luke Sampson

## Usage
`mkdir $HOME/.fonts`  
`cp univers.flf $HOME/.fonts`  
`go install`  
`cdown [args]`  

Output of `cdown -h[elp]`:  
```
-f string
      font to use (default "univers")
-m int
      minutes to count down
-msg string
      message to display when done (default "Time up!")
-s int
      seconds to count down
```

`CTRL + C` to abort.

## License
MIT (do whatever you want)
