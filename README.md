# Requirements
- Windows:
    - install [chocolatey](https://chocolatey.org/) for ```make``` command

# Commands
- make templ
    > prepare components in ```bin/components``` folder 

- make clean
    > remove all components from ```bin/components``` folder

- make serv
    > serve http server (default port :8080)

# Dev workflow
- build new components in ```src/components``` folder
- prepare created components with ```make templ``` command
- serve crated component with ```make serv``` command