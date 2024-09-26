# eros

## requirements
* Docker
* KinD
* Just

## Get started
Start the API Backend:
```shell
eros server start
```

Start the Frontend Server:
```shell
eros frontend start
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