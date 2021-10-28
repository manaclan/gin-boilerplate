# gin-boilerplate

## Getting started
Create your .env file in the same location with main.go  
Then add your mongo uri. For example:
```shell
MONGODB_URI=mongodb+srv://username:password@abc.mongodb.net/test?authSource=admin&replicaSet=atlas-iywr24-shard-0&w=majority&readPreference=primary&appname=MongoDB%20Compass&retryWrites=true&ssl=true
```
## Run development
First install air for hot reload: https://github.com/cosmtrek/air  
Then, run:
```shell
air
```

## Run deployment
```shell
docker-compose -f run.yml up --build -d
```