
# what the fuck is docker???

docker is a open source platafrom for shippig... blabalbla

Docker is a tool for packaging and running applications in isolated environments called containers.

so we can use container to isolete and reproduce enrivoment, in this way is more "safer" and well wedont have the problem "works on my machine" 

okay one question what is the container anatomy?
as far a undenst all container still runs o the host OS so they shere all resoureces with the host
what make the cool is two things:
1. namespaces: 
    Namespaces give a process its own view of certain system resources.
2. cgroups:
    cgroups answer how much can this process use? so it can control resourece like
    - cpu
    - momory
    - disk io
    - process limits


By default, a container is relatively well isolated from other containers and its host machine.


## what the fuck are namespaces?
is kernel feature that creates isolated environments for process. 

why this is aswome?
well because of that feature that we can run multple docker containers cuz the are isolated from each other.

Linux provides seven types of namespaces, each isolating different system resources:

1. PID Namespace - Isolates process IDs
2. Mount Namespace - Isolates filesystem mount points
3. Network Namespace - Isolates network interfaces and routing tables
4. UTS Namespace - Isolates hostname and domain name
5. IPC Namespace - Isolates inter-process communication resources
6. User Namespace - Isolates user and group IDs
7. Cgroup Namespace - Isolates cgroup hierarchies

okay i'm not going deeper then that cuz latter i can make some linux investrigations

okay but lest make some buzz words clear:

## container:
so a container is a isolated process for each component of a syustema like lets say that we have a backend im C, a front end in valina web and postgreSQL we can run one container for each "part" of this websystem so we can ensure that every one get the same version, and if needed reproduce this enviroment becomes easy as fuck.

## Image
so a container image is a standardized package that incluides all of the files, binaries, libs and configuratrions to run a container

## dockerfile
great another syntax to learn so dockerfile is where thigs get interesthging, like we can put some instrution for building your source code

exemple:
```dockerfile
# syntax=docker/dockerfile:1
FROM ubuntu:22.04

# install app dependencies
RUN apt-get update && apt-get install -y python3 python3-pip
RUN pip install flask==3.0.*

# install app
COPY hello.py /

# final configuration
ENV FLASK_APP=hello
EXPOSE 8000
CMD ["flask", "run", "--host", "0.0.0.0", "--port", "8000"]

```
we also have docker compose whitch is a why to orchestrate multiples containers
