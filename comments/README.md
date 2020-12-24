# Comments Service

This is the Comments service

Generated with

```
micro new comments
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


## Web Example
http://localhost:8080/comments/store?title=%22EMad%20Hello%20world%22&description=%22some%20description%22&userId=%223%22&postId=%225%22

http://localhost:8080/comments/get

## cli Example
micro comments store --title="Some" --description="description" --user_id=1 --post_id=9

micro comments get