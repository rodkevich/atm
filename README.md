# atm

## run  
docker-compose up  
docker-compose up -d  

## stop all containers  
docker stop $(docker ps -aq)  

## remove all stopped  
docker rm $(docker ps -aq)  
