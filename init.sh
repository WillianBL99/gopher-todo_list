docker build -t application-server .

docker run -it --rm -p 5050:5050 application-server