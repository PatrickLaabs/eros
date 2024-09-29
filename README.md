# eros

## requirements
* Docker
* KinD
* Just

## Get started

### Build from Source

Using just to build the binaries:
```shell
just build
```

### Running the Apps
Start the API Backend:
```shell
go run ./cmd/api-server/api-server.go

or:
./api-server
```

Start the Frontend Server:
```shell
go run ./cmd/frontend-server/frontend-server.go

or:
./frontend-server
```
You can now either open up the frontend application via `http://localhost:8080`
or do direct queries against the api, like ` curl localhost:3000/kubernetes/local/create`.

## Documentation
### Backend API
I am using the Swagger Tooling to generate the needed API documentations.
You can access the documentation with:
```shell
http://localhost:3000/swagger/index.html
```

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=PatrickLaabs/eros&type=Date)](https://star-history.com/#PatrickLaabs/eros&Date)