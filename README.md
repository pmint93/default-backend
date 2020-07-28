# HTTP Default backend

Return "default-backend - 404" by default.

But can return custom response by setting environment variable `RESPONSE_CODE` and `RESPONSE_TEXT`

```
docker run -d -p 8080:8080 -e RESPONSE_CODE=404 -e RESPONSE_TEXT=NotFound pmint93/default-backend:latest
```
