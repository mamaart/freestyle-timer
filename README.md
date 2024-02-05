# Freestyle Timer

This project aims to develop a timer for a highline freestyle competition between two athletes.

## Required Features

- [ ] Each session/battle involves two timers, one for each athlete.
- [ ] A timer can only run when no other timers are active.
- [ ] Timers count down from the specified duration of the battle.
- [ ] A new session can be created via a network request, requiring the session name and athlete names.
- [ ] Sessions run in memory until they are terminated.
- [ ] All session events are recorded and saved in a database under the session name, with timestamps and event names.
- [ ] Control events include session destruction and actions for each athlete: start, pause, reset, and adjust time.
- [ ] Other events include creation, athlete completion, session completion, and error/warning notifications.
- [ ] Each athlete's timer can be streamed over the network.
- [ ] Athlete timers can be viewed in a browser/app.
- [ ] Athlete timers can be overlaid on a video stream.
- [ ] Session events can be streamed over the network.
- [ ] Session events can be monitored live in a browser/app.
- [ ] Previous sessions can be searched, exported as CSV, or deleted via browser/app interface.
