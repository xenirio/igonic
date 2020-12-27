# Gin template for Openware go projects

## Introduction

Gin, Gorm micro framework, this repository is a boilerplate you can use with a generator.

## Todo

- Make a better migration system
- Create a logger
- Write tests
- Simplify Seeding
- Prepare daemon structure with commando
- Migrate to commando
- create drone file
- create Dockerfile

### Quick start

1. Download the tool barong-jwt to fake authentication and run it once to generate RSA key

```bash
wget https://github.com/openware/barong-jwt/releases/download/1.1.0/barong-jwt
chmod +x barong-jwt
./barong-jwt
```


2. Run the server

```bash
export JWT_PUBLIC_KEY=$(cat config/rsa-key.pub|base64 -w0)
go run ./cmd/edge
```

3. Test an authenticated endpoint

```bash
curl -i  -H "Authorization: Bearer $(./barong-jwt)" localhost:6009/api/v2/app/private/profile
```
