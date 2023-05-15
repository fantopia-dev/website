# website
the website of fantopia



# deploy

```
git clone https://github.com/fantopia-dev/website.git

cd website

docker-compose up -d

# log
docker-compose logs -f  website
```


test api:

```
curl  -s -X POST -H 'Content-Type: application/json'  -d '{"email":"youngqqcn@163.com"}' http://127.0.0.1:8888/api/v1/joinwaitlist  | jq

```

response
```json
{
  "code": 0,
  "msg": "ok",
  "data": "success"
}
```


----



# environment

install `docker` and `docker-compose`

```
sudo yum update -y

sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

sudo yum install docker-ce -y
sudo systemctl start docker
sudo systemctl enable docker


sudo wget https://github.com/docker/compose/releases/download/v2.17.3/docker-compose-linux-x86_64 -O /usr/bin/docker-compose

sudo chmod +x /usr/bin/docker-compose

```
