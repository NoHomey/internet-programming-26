# Exercise
Make the `Retry` function from the `command` package Production Ready!

In order to do that you:
1. need to add an argument that can be used to check if a given error should be retried or not;
2. need to add an argument that determines what the next pause should be. The type of this function should be `func (pause int, previousPause time.Duration) time.Duration` where `pause` should be the current pause number.

Example pause functions:
1. Constant: `func (_pause int, _previousPause time.Duration) time.Duration { return 2 * time.Second }`
2. Linearly increasing: `func (_pause int, previousPause time.Duration) time.Duration {return previousPause + 2 * time.Second }`
3. Exponential Back-off: `func (_pause int, previousPause time.Duration) time.Duration {return (1 + 0.2) * previousPause + time.Second }`

Good luck!