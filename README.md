# Contact Store [![CD](https://github.com/r-hydra/contact-store/actions/workflows/deploy-do.yml/badge.svg)](https://github.com/r-hydra/contact-store/actions/workflows/deploy-do.yml)

* > I could hardly make time for this project and didn't have enough time to do TDD
* I tried to follow a standard project structure that I found on the internet
* Didn't use anything special, I prefer using as less package as possible
* Can't do design, absolutely no CSS
* Used typescript for, you know, the types
* I know webpack very well but wanted to skip the manual configuration with laravel-mix
* The CI/CD is configured to deploy the project at http://128.199.202.48:8080
* Used supervisor to run the program in background
* Firewall not enabled, happy attacking


### Build process

* First go to `/web/` directory
* Run `yarn install` to install node modules
* Then run `yarn prod` to generate static production files
* Go back to root directory
* Run `go build cmd/web/*.go`
* Run the binary
* ..............
* ..............
* ..............
* Have fun
    