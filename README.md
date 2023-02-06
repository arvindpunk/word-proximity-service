# Word Proximity Service

Backend Go (gin) based web server for Word Proximity.

## APIs

### Internal

```
curl -X POST localhost:5001/internal/refresh-word-cache
```

### Public

```
curl localhost:5001/v1/get-target-word
```

