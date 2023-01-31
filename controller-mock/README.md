# Controller Mock
## Description
A Controller monitors workflows in production cycles and collects statuses from various metrics.  

## Metrics
* Motor Temperature 
* Motor Speed 
* Motor Noise
* Motor Power Consumption

## Functionality
Reports to the Processor every \<time interval\> over a Kafka bridge using HTTP with new metric data. 
Following a deterministic approach, starting from the lowest range limit, with every \<time interval\>, the metric data will be incremented by one unit, up until reaching the top range limit. Whereas then it will be decremented by one unit until the lowest range limit has been reached. From this point the process will be repeated indefinitely.

* Motor temperature
    * range: 70 degrees - 85 degrees
    * fluctuation unit: 1
* Motor Speed 
    * range: 5000 rpm - 6500 rpm
    * fluctuation unit: 100
* Motor noise 
    * range: 90 db - 97.5 db
    * fluctuation unit: 0.5
* Motor power consumption 
    * range: 14 kW - 21.5 kW
    * fluctuation unit: 0.5



## Components
* Controller - Generates metric data and promotes it by predefined configuration standards every single time interval
* Kafka - The Kafka Producer allocates each message containing metric data to a corresponding topic partition and sends it over HTTP to a Kafka cluster to be later consumed by the Processor-mock

## Input
Initial metrics configuration [JSON file](pkg/controller/initial_metrics_config.json)

## Output
Type: JSON packet with controller data and timestamp 
```json
{
    "value": 
    {
        "controller_id": <id>,
        "controller_name": <name>,
        "timestamp": <now>,
        "metric":
        {
            "name":"motor_temperature",
            "value":70,
            "timestamp": <now>
        }
    }
}
```
``` json
{
    "value": 
    {
        "controller_id": <id>,
        "controller_name": <name>,
        "timestamp": <now>,
        "metric":
        {
            "name":"motor_speed",
            "value":5000,
            "timestamp": <now>
        }
    }
}
```
```json
{
    "value": 
    {
        "controller_id": <id>,
        "controller_name": <name>,
        "timestamp": <now>,
        "metric":
        {
            "name":"motor_noise",
            "value":90,
            "timestamp": <now>
        }
    }
}
```
```json
{
    "value": 
    {
        "controller_id": <id>,
        "controller_name": <name>,
        "timestamp": <now>,
        "metric":
        {
            "name":"motor_power_consumption",
            "value":14,
            "timestamp": <now>
        }
    }
}
```

## Run Program Locally
From current directory run:
```bash
export HTTP_KAFKA_URL=[URL]; go run main.go
```

## Build Docker Image

From current directory run:
* Using Podman
```bash
podman build -t [NAME:TAG] -f ./docker .
```
* Using Docker
```bash
docker build -t [NAME:TAG] -f ./docker/Dockerfile .
```

## Run Program On Multiple Containers

Using [Docker Compose](https://docs.docker.com/get-started/08_using_compose/) we can run multiple instances of the controller application on docker containers.
To start docker containers, run:
```bash
KAFKA_URL=[URL] IMAGE=[NAME:TAG] docker-compose -f docker/docker-compose.yaml up --scale controller=[NUMBER_OF_REPLICAS]
```