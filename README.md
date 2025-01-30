# Go CLI template

Opinionated template for building a Go CLI application.

It uses cobra for parsing commands and flags and a viper wrapper to handle configurations.

The viper wrapper is configured out-of-the-box to support for reading configuration from files, environment variables, and flags. It makes use of standard locations to store configuration and application data files.
A `config` command is provided that allows the user to list and modify configurations.

The project structure is slight variation on the typical Cobra project structure, with each command having its own package inside the `cmd` directory.
Logic that can be shared between commands should be placed in the `lib` directory to avoid circular dependencies.


## Tips on getting started with the template

- Go to the `cmd/root.go` file and change the `appName` variable to your app's name
- Also in the `cmd/root.go` file, you might want to provide a short and long description for the command.
- Optionally, go to the `go.mod` file and change the module name `app` to your app's name. This will require some imports to be updated.
- Add commands by creating a new package in the `cmd` directory. Make sure each command package has a `GetCmd` function that returns the Cobra command. Add each command to the root command by calling the `GetCmd` function of each package in the `cmd/root.go` file.

## Build

Currently, the only way to run the project is to have Go 1.23+ installed locally and building a project.

```bash
$ go build
$ ./myapp --help # or ./myapp.exe on Windows
```
