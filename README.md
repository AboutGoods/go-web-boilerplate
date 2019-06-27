# Go web boilerplate

This boilerplate is for applications using go with mongodb

lib used:

- viper (config lib)
- logrus (logger lib)
- mongo db official driver
- gin-gonic (router and web)


## Getting started

first clone the project using the following command : 

```
git clone https://github.com/AboutGoods/go-web-boilerplate.git MyProject
```

Then, if you want to rename your module, you can edit this first line of your go.mod.  

⚠ If you change it, some files will show new issues because they depend on this package. For example in ```http/controllers/DefaultController.go```
```
import (
    "App/database"
```


Then download all the dependencies using
```bash
go mod download
```

After this, you can run your app using 

```bash
go run App

```

or build it using 
```bash
go build
```

This will generate an executable file named `App`.

## Docker

The docker image is a multi-stage build image, it builds the project using the go image and move it to a scratch image.
SSL support has been added.

## Tests

In the test folder, you can find som .http files that allows you to test the project using IntelliJ IDE Goland and make api calls for your.
If you need a custom local config for your tests, duplicate the http-client.env.json and name it http-client.private.env.json edit it has you want.


## Contribute

Feel free to contribute, this project will be greater for anyone.