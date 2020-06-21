# AA
Coding Assignment

Api endpoints
http://127.0.0.1:8001/server1 - Post request with amazon url as body will return Title,Thumbnail,Reviews and Description
Sample Body {"url":"https://www.amazon.com/Skeleteen-Witch-Broomstick-Costume-Accessories/dp/B07PHRWH4C"}

http://127.0.0.1:8002/server2 - Post request which dumps the scrapped data into the document store.



Document store used is mongo db

You can start running by executing the following commands
docker-compose build
docker-compose up

You can then use Postman or insomnia to access the api
