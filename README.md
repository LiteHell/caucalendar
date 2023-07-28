# calendar.puang.network
중앙대학교 학사일정을 스마트폰에서 편하게 확인하기 위해 만든 웹 어플리케이션입니다. 이 프로그램은 중앙대학교 학사일정을 주기적으로 크롤링하여 ICS 파일 형식으로 제공합니다.

이 어플리케이션은 Docker를 이용합니다.

## How to build and launch
```
docker build -t caucalendar https://github.com/LiteHell/caucalendar.git
docker run -p PORT:8080 -d caucalendar
```

위에서 PORT를 원하는 포트번호로 바꾸면 됩니다.