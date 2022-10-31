# Install current master nats server

```bash
cd /tmp
git clone git@github.com:nats-io/nats-server.git
cd nats-server
go build
sudo cp nats-server /usr/local/bin
```

## Create a localhost cert:

```bash
sudo apt install libnss3-tools
go get github.com/FiloSottile/mkcert
mkcert -install
mkcert localhost
```

or just start `google-chrome --ignore-certificate-errors`

```bash
nats-server -c ./server.conf
```

# Jet Stream

```bash
# Install nats-seeder
go install github.com/mheers/nats-seeder@latest

nats-seeder \
 user-nkey \
 --operator-seed SOAON2QVZ5L7CMOO5W3PV4F7OCDU7L6AXIO5VA2YWIBTTSLUN64UNOU63M \
 --account-seed SAADBIEN2MTECGRQDK3Y6XHK7PADDSXR6SOCQOM5GFORHLBAX6V6C65OOE \
 -u test \
 -p "\$JS.API.>" -s "\$JS.API.>" -p "_INBOX.>" -s "_INBOX.>" \
 -p "\$KV.CFG.*" -p "\$JS.ACK.KV_CFG.>"  \
  > test.creds
```

```bash
nats-seeder account-jwt \
--operator-seed SOAON2QVZ5L7CMOO5W3PV4F7OCDU7L6AXIO5VA2YWIBTTSLUN64UNOU63M \
--account-seed SAADBIEN2MTECGRQDK3Y6XHK7PADDSXR6SOCQOM5GFORHLBAX6V6C65OOE
```

This creates the content for a `.creds` file that can be used with the nats-cli:

```bash
# Install nats-cli
go get github.com/nats-io/natscli...

# Get account info
nats account info --creds=test.creds --trace
```
