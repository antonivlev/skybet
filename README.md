# A simple roulette API

To run:

```
$ go get  
$ go run main.go
```

Then go to http://localhost:8080/betSingle?money=12.45&number=3 and refresh a few times. Can also try: http://localhost:8080/betColour?money=32.50&colour=red


NOTEs and TODOs in the code show thoughts on code structure and future development. 

---
Monitoring wise, it would be good to see the frequency of requests, and response delay. Could write a request log to a text file and analyse.