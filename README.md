# atm  

back:  
http://localhost:11080/ANYTHING_HERE  
front:  
http://localhost:2015  

## build  
docker-compose build --parallel  

## run  
docker-compose up  
docker-compose up -d  

## stop all containers  
docker stop $(docker ps -aq)  

## remove all stopped  
docker rm $(docker ps -aq)  

## install  
cd tests  
npm install cypress --save-dev  
sudo chmod 777 -R cypress  
              -- or node_modules & so on if needed  

## cypress tests 
docker-compose up --exit-code-from atm_tests --force-recreate  


#### create and start containers 
docker-compose up  
#### start services with detached mode  
docker-compose -d up  
#### start specific service  
docker-compose up <service-name>  
#### list images  
docker-compose images  
#### list containers  
docker-compose ps  
#### start service  
docker-compose start  
#### stop services  
docker-compose stop  
#### display running containers  
docker-compose top  
#### kill services  
docker-compose kill  
#### remove stopped containers  
docker-compose rm  
#### stop all containers and remove images, volumes  
docker-compose down   

# Info:  
fast making dirs example:  
mkdir -p ./parent_dir/{dir1,dir2,dir3}  

# Gql todos  

mutation {
    createTodo(
        input:{text:"hello", userId:"world"}
        ){
            id
            text
            done
            user{
                id
                name
        }
    }
}

query{
    todos{
    id
    }
}
