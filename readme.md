# LightHouseScanner
LightHouseScanner is small program which scans given url with the LightHouse 
Tool build into Chrome. It also contains (a work in progress) analysis tooling. 

### quickstart
1. clone dir `$ - git clone https://github.com/SoenkeD/lighthouse-scanner.git`
2. set rights for the output folder `chmod -R 0777 outputs`
3. set configs `config/config.docker.json` from `config/config.docker.sample.json`
4. build docker container `docker build -o lighthouse-scanner .`
5. run load `docker run --security-opt seccomp:unconfined -v $PWD/outputs:/app/outputs:rw -v $PWD/config/config.docker.json:/config/config.json:ro lighthouse-scanner cli load start`
6. run analysis `docker run --security-opt seccomp:unconfined -v $PWD/outputs:/app/outputs:rw -v $PWD/config/config.docker.json:/config/config.json:ro lighthouse-scanner cli analyze start`


### install
1. install go dependencies `$ - go get` 
2. go to client dir `$ - cd lighthouse-client` 
3. install npm dependencies `$ - npm ci`
4. create config file at `config/config.json` based on `config/config.sample.json`


### cli
1. `cli load start` - run the lighthouse analysis on the urls in the configs
2. `cli analyze start` - starts the analysis on the data in the output directory