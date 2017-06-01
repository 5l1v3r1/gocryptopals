Hello CTF Builder!

This is a very basic outline on how to get your CTF challenge up and running on docker. 

This assumes a few things, and because you've been approached to make a challenge, we are sure you can handle it.

- That you know what docker is.
- That you have a (locally) working challenge (or at least the framework)
- That you can install docker good

PROD VERSION OF DOCKER:
docker-ce | 17.03.1~ce-0~ubuntu-xenial | https://download.docker.com/linux/ubuntu xenial/stable amd64 Packages
docker-compose version 1.13.0
 - VERSION 3.0 syntax

 Get these versions of docker/docker-compose to ensure that your challenge will work as expected in the game environment. Other versions of docker *may* work, but don't be that guy.
 _if your challenge doesn't work in QA, and we find out it's because of docker versioning, we will be very unhappy_

OK! Now that we have that out of the way... In this directory, there is a few files. I will draw your attention to docker-compose.yml:

This is the docker-compose file. It defines how the docker containers will be run