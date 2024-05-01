# Example service for [DI](https://github.com/sabahtalateh/di)

This example includes various usage of `di` functions

See 
- [internal/config/config.go](https://github.com/sabahtalateh/diexample/blob/main/internal/config/config.go) for usage of `di.InitE` and `di.Setup`
- [internal/repos/greetings.go](https://github.com/sabahtalateh/diexample/blob/main/internal/repos/greetings.go) for usage of `di.Name`, `di.Init` and `di.Get` without `di.Name`
- [internal/services/greeting.go](https://github.com/sabahtalateh/diexample/blob/main/internal/services/greeting.go) for usage of `di.Get` with `di.Name`
- [internal/system/db.go](https://github.com/sabahtalateh/diexample/blob/main/internal/system/db.go) for usage of `di.Stage`
- [internal/setup/stages/stages.go](https://github.com/sabahtalateh/diexample/blob/main/internal/setup/stages/stages.go) contains list of stages
- [internal/setup/container.go](https://github.com/sabahtalateh/diexample/blob/main/internal/setup/container.go) for example of container creation with list of setup functions


Start Postgres in docker container
```sh
make env.start
```

Run app
```sh
make run
```

Stop Postgres in docker container
```sh
make env.stop
```


