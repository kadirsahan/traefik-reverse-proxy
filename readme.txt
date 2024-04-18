web2:
docker build -t web2:1.0 .
docker build -t web2:2.0 . # same port

web1:
docker build -t web1:1.0 . # different port
docker build -t web1:2.0 . # same port



docker run --rm -p 8080:8080 [container-name]:latest

docker build -t web2:2.0 .
docker run --rm web2:2.0

docker build -t web1:2.0 .
docker run --rm web1:2.0

credit : https://iceburn.medium.com/reverse-proxy-in-traefik-with-subdirectories-eef4261939e

credit : https://devopscube.com/create-self-signed-certificates-openssl/



problem : https://stackoverflow.com/questions/61946945/traefik-pathprefix-is-not-working-as-expected-for-underlying-requests

