# Direct-Message_backend
```
curl -X POST -H "Content-Type: application/json" -d '{"username":"user1","password":"password1"}' http://localhost:8090/register
curl -X POST -H "Content-Type: application/json" -d '{"username":"user1","password":"password1"}' http://localhost:8090/login
curl  -H "Content-Type: application/json" -d '{"token":"fe6440c86adff1a3789cb5803dd7d638c80aa9da09282b9071d683dd3ee22433", "searchName":"", "PageIdx":1}' http://localhost:8090/get-info
```