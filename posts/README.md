# Posts Service

This is the Posts service

Generated with

```
micro new posts
```

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run .
```

## Example
http://localhost:8080/posts/store?title="Hello%20world"&slug="hello-world"&description="some%20description"&user_id="3"

http://localhost:8080/posts/get


## cli Example
micro posts store --title="Some" --slug="Some" --description="description" --user_id="1"

micro posts get