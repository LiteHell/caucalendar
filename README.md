# caucalendar.online
중앙대학교 학사일정을 스마트폰에서 편하게 확인하기 위해 만든 웹 어플리케이션입니다. 이 프로그램은 중앙대학교 학사일정을 주기적으로 크롤링하여 ICS 파일 형식으로 제공합니다.

이 어플리케이션은 Docker를 이용합니다.

## How to build and launch
```
docker build -t caucalendar https://github.com/LiteHell/caucalendar.git
docker run -v /PATH/TO/DATA/PATH:/app/data -p PORT:PORT -d caucalendar
```

## Configuration file
You need configuration file `config.json` inside data directory.

```
{
    "port": 80,
    "database": "sqlite://debug.db"
}
```
|   name   |        definition       |
|----------|-------------------------|
|   port   | http port for listening |
| database | database connection uri |