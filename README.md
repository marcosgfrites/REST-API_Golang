# REST-API_Golang

### Example CURLs

- CURL for check if API is running
```cmd
curl --location --request GET 'http://localhost:8080/'
```

- CURL for test GetPosts function (GET)
```cmd
curl --location --request GET 'http://localhost:8080/posts'
```

- CURL for test AddPost function (POST)
```cmd
curl --location --request POST 'http://localhost:8080/posts' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "random title for test AddPost",
    "text": "random text for test AddPost"
}'
```