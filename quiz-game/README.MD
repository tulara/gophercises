### A reminder how to run
```
> go build .\quizGame.go
> .\quizGame.exe -filename="location" -limit=2
```

### GOAL: want to also print totals if timeout finishes before quiz does
but also remain testable
what if we published answer to another channel each time it was completed
hmm you'd get a backlog of answers you'd then have to aggregate -> buffer issues?