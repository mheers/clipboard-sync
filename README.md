# Clipboard sync

Syncs the clipboard over multiple machines and operating systems.

## Create seeds for the docker-compose file

```bash
make create-seeds
```

Add the resulting lines to the `.env`

## Create a credentials file

```bash
make create-credentials
```

## Usage in corporate networks

In many corporate networks, the clipboard sync will not work out of the box. This is because the nats server is not reachable.

As secure way - without using a local VPN - is to use a ssh tunnel in the browser.

The approach will use

- tailscale on the local machine
    - on this run the docker-compose from github.com/mheers/docker-ssh-server
    - exec into the container as `client` and run `screen -S shared`
- https://webvm.io on the corporate machine
    - in this connect to tailscale
    - in this connect to the ssh server using client:client as user:password and 2244 as port
    - in this run `screen -x shared`


# TODO
- [x] directly read the .creds file
- [x] add a way to specify the nats server
- [ ] sync keyboard directly
- [ ] use headless chrome to sync the clipboard
