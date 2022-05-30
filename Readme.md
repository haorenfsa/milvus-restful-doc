# Milvus RESTful doc

This repository is for generate RESTful API docs.

Here is how it works:

First, use `swaggo` to generate `docs/swagger.yaml` swagger v2.0 specifications.

```
swag init
```

Then use convert tool to convert `docs/swagger.yaml` to markdown or html

```
redoc-cli build ./docs/swagger.yaml
```

If you want md: use
`swagger-markdown -i ./docs/swagger.yaml -o api-doc.md`

# tools:
- https://github.com/swaggo/swag