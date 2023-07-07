# rediboard-go
A simple app using golang and redis. I made this to learn how to connect with redis using golang. 

## Routes
POST -> /points           -> send a request with JSON data that consists of { "username: "your_username", "points": 99 } 
GET  -> /points/:username -> send a request with the username to get info about the username specified
GET -> /leaderboard       -> get all the the users in the redis db. They will be sorted according to the points in ascending order. 
