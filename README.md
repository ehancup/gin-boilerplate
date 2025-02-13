# Gin-gonic Boilerplate

this Gin Boilerplate is inspirated by NestJS.

this boilerplate include package :
- Gin gonic
- Gorm
- Swaggo
- Logger
- Validator

## How to use it
1. Run
```bash
$ go mod tidy
```
2. make a `.env` file (you can copy all necessary variable in `.env.example`)
3. set the value of `.env`   
    note : if you want to use database. you can see [gorm documentation](https://gorm.io/docs/connecting_to_the_database.html) and set the `DB_DSN` value in `.env`, if you dont use it, just leave it same as in the `.env.example`.
4. If you use database, you need to uncomment code in `src/app/index.go` line 43-53
