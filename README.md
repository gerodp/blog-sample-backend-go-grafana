# Blog Backend API
This is a simple Blog Backend API written in Go and intended to be used to support the Blog article about performance testing [TODO PASTE URL HERE].

The project was tested succesfully on Mac with M1 Chip and Linux with amd64 architectures.

For covenience we include a Makefile with commands to ease local development.

To launch the containers run the following command:

```
make start
```

The following services will be started:

- backend: The backend in Go that implement the Blog API
- mariadb: MariaDB Database Server
- cadvisor: Collects different container usage metrics
- prometheus: Monitoring backend that will collect the metrics from backend
- grafana: Monitoring frontend with dashboards to check different metrics

To access Grafana open a browser and navigate to http://localhost:3000

To access Prometheus open a browser and navigate to http://localhost:9090


## Deployment in AWS EC2

* Create AWS EC2 preferebly with Ubuntu Linux and x86 compatible CPU (Tested with t3.large and m5.large)

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

sudo usermod -aG docker ubuntu

sudo docker run hello-world

# Log out and log in again

git clone git@github.com:gerodp/blog-sample-backend-go-grafana.git

cd blog-sample-backend-go-grafana

make start

```


## Sample curl API calls

Login
```
curl -X POST -H "Content-Type: application/json" -d '{"username":"testint1", "password":"testint1"}' http://localhost:9494/login
```

List all users
```
curl -H "Authorization:Bearer [TOKEN OBTAINED IN LOGIN]" http://localhost:9494/auth/user
```

List all posts
```
curl -H "Authorization:Bearer [TOKEN OBTAINED IN LOGIN]" http://localhost:9494/auth/post
```

Create a new user
```
curl -X POST -H "Content-Type: application/json" -H "Authorization:Bearer [TOKEN OBTAINED IN LOGIN]" -d '{"username":"pepe", "email":"pepe@blogapp.com", "password":"testint1"}' http://localhost:9494/auth/user
```

Try Create a new user with blank username
```
curl -X POST -H "Content-Type: application/json" -H "Authorization:Bearer [TOKEN OBTAINED IN LOGIN]" -d '{"username":"", "email":"pepe@blogapp.com", "password":"123"}' http://localhost:9494/auth/user
```