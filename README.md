# golang-todo-api
A simple todo list api with go and mongodb


# USING THIS


```bash
docker-compose up -d
```

live reload using modd ;) 


# EXAMPLES

```bash
curl -XPOST -d '{"task": "Kill Stormtroopers", "completed": false}' localhost:9090

curl localhost:9090

# other rest operations work as well, but i'm too lazy to create example ;) 
```