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

# TODO
- [ ] directly read the .creds file
- [x] add a way to specify the nats server
- [ ] sync keyboard directly
