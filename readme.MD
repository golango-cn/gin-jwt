
### 请求登录

```
curl http://localhost:8000/login \
 -X POST \
 -H "Content-type:application/json" \
 -d '{"username": "admin", "password": "123" }'

{"code":200,"expire":"2020-03-20T23:21:52+08:00","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODQ3MTc3MTIsIm1vYmlsZSI6IjE1MjExMTEwMDAwIiwibmFtZSI6ImFkbWluIn0.TTQz9urgb34Gt8KWQ_3fqbXRn_JlXH4PGFGV4XvYUYI"}
```

### 获取用户信息
```json

curl http://localhost:8000/api/user \
 -X GET \
 -H "Content-type:application/json" \
 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODQ3MTc3MTIsIm1vYmlsZSI6IjE1MjExMTEwMDAwIiwibmFtZSI6ImFkbWluIn0.TTQz9urgb34Gt8KWQ_3fqbXRn_JlXH4PGFGV4XvYUYI"

{"message":"ok","user":{"exp":1584717712,"mobile":"15211110000","name":"admin"}}

```

