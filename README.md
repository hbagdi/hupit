# hupit

Watch files and hup processes when they change.

hupit watches files and executes an arbitrary shell command whenever any
of the files being watches changes.


```shell
$ hupit --help
Usage of hupit:
  -command string
        command to execute when a file changes
  -file value
        file to watch for changes;flag can be used multiple times to specify multiple files
  -version
        output version information
```

# Dockerimage

Docker image is hosted on Dockerhub:

```shell
$ docker pull hbagdi/hupit
```
