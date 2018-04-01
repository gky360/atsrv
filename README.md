# atsrv

Backend for atcli ( https://github.com/gky360/atcli )

## Requirements

- chromedriver >= 2.35

## Endpoints

| method | need token | need logged in | path |
|---|---|---|---|
| GET  | n | n | / |
| GET  | y | y | /me |
| GET  | y | y | /contests/arc090 |
| POST | y | y | /contests/arc090/join |
| GET  | y | y | /contests/arc090/tasks |
| GET  | y | y | /contests/arc090/tasks?full=true |
| GET  | y | y | /contests/arc090/tasks/d |
| GET  | y | y | /contests/arc090/submissions |
| GET  | y | y | /contests/arc090/submissions?task_name=d&status=AC |
| GET  | y | y | /contests/arc090/submissions/2167890 |
| POST | y | y | /contests/arc090/submissions?task_name=d |
