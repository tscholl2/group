# group
A kind-of generic group

##TODO

- ~~Write something~~
- Write group operations
- Write encryption mechanism (AES)
- Write API
- Write example DH key exchange alg

## API

###/scale/:k

returns kg, i.e. the group generator scaled by k

###/scale/:h/:k

returns kh, i.e. h scaled by k

##Notes

generate ssl certs with

    openssl ecparam -genkey -name prime256v1 -out key.pem
    openssl req -new -key key.pem -out csr.pem
    openssl req -x509 -days 365 -key key.pem -in csr.pem -out certificate.pem
