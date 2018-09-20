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

- [chromedriver](http://chromedriver.chromium.org/) >= 2.35


## Installation

```
go get -u github.com/gky360/atsrv
```


## Usage

1. Generate auth token (like `cat /dev/urandom | base64 | fold -w 32 | head -n 1` )
   and set it to environment variable ATSRV_AUTH_TOKEN.
2. Set user id of AtCoder to environment variable ATSRV_USER_ID.
3. Run `atsrv` .
4. Enter your password to login to AtCoder.
5. `atsrv` will start running.


## Endpoints

| method | need token | exapmle path |
|---|---|---|
| GET  | n | / |
| GET  | y | /me |
| GET  | y | /contests/arc090 |
| GET  | y | /contests/arc090?with_testcases_url=true |
| POST | y | /contests/arc090/join |
| GET  | y | /contests/arc090/tasks |
| GET  | y | /contests/arc090/tasks?full=true |
| GET  | y | /contests/arc090/tasks/d |
| GET  | y | /contests/arc090/submissions |
| GET  | y | /contests/arc090/submissions?task_name=d&status=AC |
| GET  | y | /contests/arc090/submissions/2167890 |
| POST | y | /contests/arc090/submissions?task_name=d |
