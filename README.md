# Blog App
This is a simple Blog Application intended to be used just as a sample app for testing purposes.

The project was tested succesfully on Mac with M1 Chip and Linux with amd64 architectures.

For covenience we include a Makefile with commands to ease local development

To launch the containers with the integration tests container
```
make starttest
```

To launch the containers without integration tests container

```
make start
```

## TODOs

* Use Grafana Recording Rules to optimize compute used

## Deployment in AWS EC2

* Create AWS EC2 preferebly with Ubuntu Linux and x86 compatible CPU

```

sudo apt install make

sudo apt-get update
sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin

sudo docker run hello-world

sudo usermod -aG docker ubuntu

# Log out and log in again

git clone git@github.com:gerodp/perf-test-sample-app.git

cd perf-test-sample-app

make starttest
```

## Testing 

We use the following frameworks:

* https://github.com/stretchr/testify

* https://github.com/gavv/httpexpect

* https://github.com/go-testfixtures/testfixtures

### Monitoring

Grafana Dashboards to monitor DBs

https://github.com/percona/grafana-dashboards



Enabling performance schema in DB to monitor performance. Edit the file ./mariadb/config.cnf and set the PERFORMANCE_SCHEMA var to ON or OFF

Check if it's on/off:

```
SHOW VARIABLES LIKE 'perf%';
```

Then restrict for wich users to run the performance monitoring:

UPDATE performance_schema.setup_actors
       SET ENABLED = 'NO', HISTORY = 'NO'
       WHERE HOST = '%' AND USER = '%';

INSERT INTO performance_schema.setup_actors
       (HOST,USER,ROLE,ENABLED,HISTORY)
       VALUES('%','root','%','YES','YES');


Install the plugins to collect query response time

INSTALL PLUGIN QUERY_RESPONSE_TIME_AUDIT SONAME 'query_response_time.so'
INSTALL PLUGIN QUERY_RESPONSE_TIME SONAME 'query_response_time.so';
INSTALL PLUGIN QUERY_RESPONSE_TIME_READ SONAME 'query_response_time.so';
INSTALL PLUGIN QUERY_RESPONSE_TIME_WRITE SONAME 'query_response_time.so';

SET GLOBAL query_response_time_stats = ON;

SHOW PLUGINS;


### Performance Tests

We use k6 tool from Grafana Labs

https://k6.io/


k6 works with the concept of virtual users (VUs), which run your test scripts. VUs are essentially parallel while(true) loops.


### Profiling

The profiler web interface can be accessed via browser by opening

http://localhost:9494/debug/pprof

We can access the /debug/pprof/profile endpoint to activate CPU profiling. Accessing this endpoint executes CPU profiling for 30 seconds by default. For 30 seconds, our application is interrupted every 10 ms. Note that we can change these two default values: we can use the seconds parameter to pass to the endpoint how long the profiling should last (for example, /debug/pprof/profile?seconds=15), and we can change the interruption rate (even to less than 10 ms). But in most cases, 10 ms should be enough, and in decreasing this value (meaning increasing the rate), we should be careful not to harm performance. After 30 seconds, we download the results of the CPU profiler.


You can open the file with this command:

```
go tool pprof -http=:8181 ~/Downloads/profile
```

## Sample curl API calls

List all users
```
curl -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzM0NTU4ODQsImlkIjoidGVzdGludDEiLCJvcmlnX2lhdCI6MTY3MzQ1MjI4NH0.FfCowmEOuGIXTAwl9geWmbItMBEbpKJbSasECQRJwW8" http://localhost:9494/auth/user
```

List all posts
```
curl -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzM0NTU4ODQsImlkIjoidGVzdGludDEiLCJvcmlnX2lhdCI6MTY3MzQ1MjI4NH0.FfCowmEOuGIXTAwl9geWmbItMBEbpKJbSasECQRJwW8" http://localhost:9494/auth/post
```

Login
```
curl -X POST -H "Content-Type: application/json" -d '{"username":"testint1", "password":"testint1"}' http://localhost:9494/login
```


Create a new user
```
curl -X POST -H "Content-Type: application/json" -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzMzNTUzNTksImlkIjoidGVzdGludDEiLCJvcmlnX2lhdCI6MTY3MzM1MTc1OX0.KPlH4EDFpJMdgWo2ZidMr4cvuw3wnvT_fnzHdqDaHcI" -d '{"username":"pepe", "email":"pepe@blogapp.com", "password":"testint1"}' http://localhost:9494/auth/user
```

Try Create a new user with blank username
```
curl -X POST -H "Content-Type: application/json" -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzI4NjA2ODMsImlkIjoicGVwZSIsIm9yaWdfaWF0IjoxNjcyODU3MDgzfQ.CkOlTSVsnjdIYoAFdgJSP91SBa2gchwV_33jYrMHOUs" -d '{"username":"", "email":"pepe@blogapp.com", "password":"123"}' http://localhost:9494/auth/user
```