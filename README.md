## Golang REST API Example

### API Documentation

API documentation available at [Postman Public Directory](https://documenter.getpostman.com/view/5001481/TVeiDWGp)

### Docker
If you want to build this service:
```
docker-compose up --build
```

If you want to run this service in the background:
```
docker-compose up -d
```

If you want to stop this service:
```
docker-compose stop
```

If you want to real time build when you are coding:
```
brew install watchexec
docker-compose up -d
watchexec --restart --exts "go" --watch . "docker-compose restart app"
```