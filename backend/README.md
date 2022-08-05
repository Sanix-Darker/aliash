# ALIASHUB

The documentation should be generated, the whole tech side should not be in this readme file !

### Interactive runner
Will show you each line of your code and execute for you !
```bash
curl http://127.0.0.1:5000/i | bash -s we-ap9dea4
```

### None interactive runner 
Will directly execute your saved aliases bash script
```bash
curl http://127.0.0.1:5000/we-ap9dea4 | bash
```


### To get the whole list

Maybe a paggination should be implemented here !
```bash
curl -s http://127.0.0.1:5000/all | jq
```
