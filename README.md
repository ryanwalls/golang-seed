# golang-seed *Description required
Seed project for starting a dockerized, golang project.  Features include:
* Command parsing with https://github.com/spf13/cobra
* Logging setup with https://github.com/inconshreveable/log15
* Configuration management with https://github.com/spf13/viper
* Dependency management and vendoring with https://github.com/Masterminds/glide

Feel free to fork this repo to start your own project.  The rest of the README below
assumes you are using this as a template and will fill in extra details.  

If you want to duplicate without forking see: https://help.github.com/articles/duplicating-a-repository

## Technical Specifications *required

### Platforms Supported
For all projects, it will be MacOS, Windows, and Linux or N/A

### Inputs / Outputs
Briefly describe inputs and expected outputs if needed

### Resource Requirements
Memory / CPU if applicable

## Structure

* `commands` - add commands to execute/start your code in this package. You should not have to modify `main.go`.
Read more about the library being used here: https://github.com/spf13/cobra.
* `config` - package for managing configuration.  Currently setup for dev, qa, prod environments.  Defaults
to dev.  To set the environment, make sure an environment variable `ENV` is set on your target system.  
Configuration put in `config-common.yml` is accessible in all environments. To
get a value from configuration do something like this after importing `config` package:
  ```
  myConfigurationValue := config.Viper.GetString("myConfigurationKey")
  ```
  See https://github.com/spf13/viper for more usage information.  
* `logger` - package to use for logging.  Import this package and setup your own logger
or use the root logger.  e.g.  

  ```
  // Root logger
  logger.Log.Info("This is using the root logger")

  // Child logger
  myPackageLogger := logger.Log.New("package", "myPackage", "someContextKey", "someContextValue")
  myPackageLogger.Info("This uses myPackageLogger which is a child of the root logger and shares root logger configuration.")
  ```
  See https://github.com/inconshreveable/log15 for more usage information.

## Setup Instructions

### Install Go (Windows)
Install chcolatey if you have not already done so (https://chocolatey.org/)
```
choco install golang
```

### Install Go (Mac)
https://golang.org/doc/install

### IDE
* Install an IDE.  Recommend Visual Studio Code
* Create a directory for your go projects, ex: `c:\source\go`

#### Setting up Intellisense in Visual Studio Code (Mac/Windows)
Install `go` extension in visual studio code.
* Open extensions window (cmd + shift + x)
* Type `go`
* Install the one with author of lukehoban.

### Setting up Code directories
Assuming your workspace folder is `c:\source\go`, setup  environment variable `GOPATH = c:\source\go`

Create bin, pkg, src sub-folders in `$GOPATH`.

Create `$GOPATH\src\github.com\<your github username/organization>` folder.

In the `$GOPATH\src\github.com\<your github username/organization>` folder, clone the repo. When finished you should have the folder: `$GOPATH\src\github.com\<your github username/organization>\golang-seed`

### Packages
We use https://github.com/Masterminds/glide to manage package dependencies.

1.  Install glide and add it to your path.  See https://github.com/Masterminds/glide#install
2.  Inside `golang-seed` directory run  `glide install --strip-vendor`.  This downloads all package dependencies
to a `vendor/` folder.
3.  If you need to add more dependencies to the project at any point run
`glide get <package name>`, which will add the package to the `glide.yaml`.  If
you want to update dependencies to the latest version run `glide up`, this will update
the `glide.lock` file with the latest dependencies.  

### Running tests
Golang has testing built in.  

To run tests for a single package run:
```
go test ./<package name>
```

To run all tests run:
```
// Run all tests but skip vendor directory
go test $(glide novendor)
```

## Automated deployments via Jenkins/Travis CI/etc

TODO fill in your deployment pipeline info here

## To run locally
* From root directory of project run
```
go install
```
* To see what commands you can run:
```
golang-seed
```

## Coding tips

### Testing
* Generate mocks using counterfeiter
```
go get github.com/maxbrunsfeld/counterfeiter
```

From project folder or inside package folder
```
go generate
```

## Contributing code
Read this article and follow the steps they outline: http://scottchacon.com/2011/08/31/github-flow.html

All PRs should be signed off by a member of the team before merging.

## Team *Contributors Required
* List
* Your
* Team
* Members

## Original release
February 2016
