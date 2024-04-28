1. To Build the image run:
```
docker build -t go_server .
```

2. To run the image run:
```
docker run --rm go_server
```

3. To build pkl config as json:
```
# Write each resource as its own file to a path within `.output/`
$ pkl eval -m .output/ **/*.pkl -f json
```
