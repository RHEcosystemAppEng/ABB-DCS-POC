# Processor Mock
## Description
The Processor consumes data received from the Controller and aggregates it. This service should send metrics to Prometheus and store data.

## Components
* Ingester - reads input data
* Aggregator - <TBD>
* Conveyor - passes data to prometheus/ storage

## Input
Type: JSON packet with metric data and timestamp 
```json
{
    "motor_temperature_c": {
        "value":70
    },
    "motor_rpm": {
        "value":5000
    },
    "motor_db": {
        "value":90
    },
    "timestamp": <now>
}
```

## Output
<TBD>
