#Create your local SSL certificates

```
openssl req -new -newkey rsa:4096 -x509 -sha256 -days 365 -nodes -out certs/location-tracker-dev.crt -keyout certs/location-tracker-dev.key

export public certificate for client

openssl x509 -outform pem -in ./certs/location-tracker-dev.crt -out ./certs/cacert.pem

or, when app is running,

openssl s_client -showcerts -servername  127.0.0.1 -connect 127.0.0.1:8080 > certs/cacert.pem

export human-readable data

openssl x509 -inform PEM -in certs/cacert.pem -text -out certs/certdata.txt

```