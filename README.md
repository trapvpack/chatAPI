### Test chatAPI app on Go, with Goose and GORM

#### Currently implemented:
- HTTP API on `net/http`
- PostgreSQL support via GORM
- Database migrations using Goose
- Docker/docker-compose for the application and database
- Basic architecture with layering (handler/usecase/repository)

#### Requirements:
- Go 1.24
- Docker, docker-compose
- PostgreSQL via Docker

#### How to start:
- `git remote add origin git@github.com:trapvpack/chatAPI.git`
- `cd /chatAPI`
- `source env/goose.env` (`goose.env` pushed in repo to simplify building)
- `docker-compose up --build`
- `docker-compose logs app`

