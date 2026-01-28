### Test chatAPI app on Go, with Goose and GORM

#### Currently implemented:
- Connecting to PostgreSQL via GORM
- Database migration via Goose
- Docker for containerizing databases and applications

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