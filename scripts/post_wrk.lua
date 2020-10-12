wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"

--wrk -t10 -c100 -d15s --script=./post_wrk.lua --latency http://127.0.0.1:6788/redis/set\?key\=name\&value\=feifei