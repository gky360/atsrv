# atsrv

```
        __
       /\ \__
   __  \ \ ,_\   ____  _ __   __  __
 /'__'\ \ \ \/  /',__\/\''__\/\ \/\ \
/\ \L\.\_\ \ \_/\__, '\ \ \/ \ \ \_/ |
\ \__/.\_\\ \__\/\____/\ \_\  \ \___/
 \/__/\/_/ \/__/\/___/  \/_/   \/__/
```

Backend for atcli ( https://github.com/gky360/atcli )


## Requirements

- chromedriver >= 2.35


## Installation

```
go get -u github.com/gky360/atsrv
```


## Usage

1. Run `atsrv`
2. Enter your password to login to AtCoder
3. Copy generated auth token
4. Set token to `atcli` with `atcli config -a xxxxxxxxxx`


## Endpoints

| method | need token | path |
|---|---|---|
| GET  | n | / |
| GET  | y | /me |
| GET  | y | /contests/arc090 |
| POST | y | /contests/arc090/join |
| GET  | y | /contests/arc090/tasks |
| GET  | y | /contests/arc090/tasks?full=true |
| GET  | y | /contests/arc090/tasks/d |
| GET  | y | /contests/arc090/submissions |
| GET  | y | /contests/arc090/submissions?task_name=d&status=AC |
| GET  | y | /contests/arc090/submissions/2167890 |
| POST | y | /contests/arc090/submissions?task_name=d |
