# Processor Mock
## Description
The Processor consumes data received from the Controller and aggregates it. This service should send metrics to Prometheus and store data.

## Components
* Ingester - reads input data
* Aggregator - <TBD>
* Conveyor - passes data to prometheus/ storage

## Input
Type: JSON packet with workflow data and timestamp 
```json
{
    "workflow_Id": <uuid>,
    "timestamp": <now>,
    "metrics": {
        "motor_temperature": {
            "value":70
        },
        "motor_speed": {
            "value":5000
        },
        "motor_noise": {
            "value":90
        },
        "motor_power_consumption": {
            "value":14
        }
    }
}
```

## Output
<TBD>
