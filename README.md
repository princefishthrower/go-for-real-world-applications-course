# Lesson 15

Restarting the Container with Minimal Downtime

Command is:

```shell
docker-compose build --no-cache && docker-compose up -d --force-recreate
```

*Note this is "minimal" downtime and not truly zero downtime.