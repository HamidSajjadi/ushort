# USHORT

UShort is another URL Shortener, developed with Go and Gin. It is a learning project for me to practice Go other things
like dependency injection, unit testing and other things.

Thus, it may seem a little over-complicated in some aspects and naive in others. Feel free to give any feedback or
suggestion to improve it.

## API Reference

#### Shorten a URL

```http
  POST /api/shorten
```

| Parameter | Type     | Description                         |
|:----------|:---------|:------------------------------------|
| `Source`  | `string` | Url to be shortened (**Required**). |

#### Redirect to the original URL

```http
  GET /:shortenedURL
```

| Parameter      | Type     | Description                          |
|:---------------|:---------|:-------------------------------------|
| `shortenedURL` | `string` | URL to be redirected (**Required**). |

## Checklist

- [x]  ~~Add configuration management~~
- [x]  ~~Add Redis as a Key-Value Database option~~
- [ ]  Create Dockerimage and docker-compose files
- [ ]  Add PSQL as another Database option
- [ ]  Deploy on AWS
