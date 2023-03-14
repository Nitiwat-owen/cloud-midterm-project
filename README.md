# cloud-midterm-project-2110524

contributors: 
1. Nitiwat Jongruktrakoon 6331320421
2. Theerachot Dejsuwannakij 6331315321

Getting Start 
1. clone the repository
2. download dependencies by `go mod download`
3. create `.env` for connect to the database
    - `DB_USERNAME`
    - `DB_PASSWORD`
    - `DB_NAME`
    - `DB_HOST`
 4. run server by `go run main.go`

Database Schema:
1. messages table
    - uuid: CHAR(36), UNIQUE
	- author :  VARCHAR(64)
    - message : VARCHAR(1024)
    - likes : UNSIGNED INT(10)
    - is_deleted: timestamp
    - last_update_at: timestamp
    - last_image_update: timestamp
2. users table
    - username: VARCHAR(191), UNIQUE
	- last_online_at: datetime
