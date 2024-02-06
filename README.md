# Freestyle Timer

This project aims to develop a timer for a highline freestyle competition between two athletes.

## Required Features

- [x] Each session/battle involves two timers, one for each athlete.
- [x] A timer can only run when no other timers are active.
- [x] Timers count down from the specified duration (2 minutes) of the battle.
- [x] A new session can be created via a network request, requiring the session name and athlete names.
- [x] Sessions run in memory until they are terminated.
- [ ] Control events include session destruction and actions for each athlete: start, pause, reset, and adjust time.
  - [x] start
  - [x] pause
  - [ ] reset
  - [ ] add second
  - [ ] remove second
- [ ] Other events include creation, athlete completion, session completion, and error/warning notifications.
  - [x] session creation
  - [ ] athlete completion
  - [ ] error warning
- [ ] All session events are recorded and saved in a database under the session name, with timestamps and event names.
- [ ] Each athlete's timer can be streamed over the network.
- [ ] Athlete timers can be viewed in a browser/app.
- [ ] Athlete timers can be overlaid on a video stream.
- [ ] Session events can be streamed over the network.
- [ ] Session events can be monitored live in a browser/app.
- [ ] Previous sessions can be searched, exported as CSV, or deleted via browser/app interface.
