## Details

Veriif is a verification application that will verify certificates issued by a CerTiiF compatible issuer. This version of Veriif will verify certificates issued by a version `0.3.0` or older of CerTiiF.

## Getting started
Execute the application file you should then see a prompt like the following:

![Samuel Hawksby-Robinson - Veriif screenshot](https://user-images.githubusercontent.com/5702426/51871815-5b243980-234f-11e9-9f3b-9f5ddf770278.png)

For testing purposes you can enter the file names of the two example files provided in the below binaries:

**Succeeds -**
```
example/cert_0.pdf
```

**Fails -**
```
example/cert_1.pdf
```

![screenshot 2019-01-23 16 57 01](https://user-images.githubusercontent.com/5702426/51871818-5e1f2a00-234f-11e9-9834-140c81709d09.png)
![screenshot 2019-01-23 16 57 01](https://user-images.githubusercontent.com/5702426/51871825-624b4780-234f-11e9-9fd8-2fc7deed6503.png)

## Downloads

https://github.com/TiiQu-Network/Veriif/releases

## Docker

For launching in a server context use the following commands. Ensure that you've `cd`ed into the `veriif` directory on the server.

```bash
docker-compose down
docker system prune
docker-compose build
docker-compose up -d
```
