# atm
## stop all containers
docker stop $(docker ps -aq)
## remove all stopped
docker rm1 $(docker ps -aq)
