# Docker Image Example in Go

This folder contains an example web application called `webapp`. It is compiled for linux and has a directory named 
`assets` where its static content lives. 
The original source code is written in Go and is included for reference.

When the webapp runs it checks for an environment variable called `SITE`. If `SITE` is set to `1` it will use
the static assets from the `assets/site1` directory. If the environment variable is set to `2` it will use the 
static assets from the `assets/site2` directory. If it isn't set, it will default to `assets/site1`.

The Dockerfile adds the compiled binary and the assets folder to the container when running `docker build`.
There are comments in the Dockerfile explaining each line in more detail.


