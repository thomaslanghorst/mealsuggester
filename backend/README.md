# mealsuggester backend
The mealsuggester is a small little helper I wrote to help me plan my meals for the week. This is the backend part written in golang. It is a simple REST service backed by a SQLite database. I am running it on my raspberry pi.

### build binary for pi
`GOOS=linux GOARCH=arm go build -o ./bin/pi/mealsuggester`

### copy binary to pi
`scp ./bin/pi/mealsuggester <pi_user>@<pi_ip>:/path/on/your/pi`

### REST calls
* list meals 
    `curl -X GET <pi_ip>:8081/meals?category=fast%20food`
* create a meal
    `curl -X POST <pi_ip>:8081/meals -d '{"name":"Burger","category":"fast food"}'`
* edit a meal
    `curl -X PUT <pi_ip>:8081/meals -d '{"id":"1", "name":"Yummy Burger","category":"fast food"}'`
* delete a meal
    `curl -X DELETE <pi_ip>:8081/meals/1`
* suggest meals
    `curl -X GET <pi_ip>:8081/suggestions?category=fast%20food&count=2`