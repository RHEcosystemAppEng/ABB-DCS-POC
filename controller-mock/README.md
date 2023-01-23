# Controller Mock
## Description
A Controller monitors workflows in production cycles and collects statuses from various metrics.  

## Metrics
* Motor Temperature 
* Motor Speed 
* Motor Noise
* Motor Power Consumption

## Functionality
Reports to the Processor every <time interval> with new metric data. 
Following a deterministic approach, starting from the lowest range limit, with every <time interval>, the metric data will be incremented by one unit, up until reaching the top range limit. Whereas then it will be decremented by one unit until the lowest range limit has been reached. From this point the process will be repeated indefinitely.

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
* Workflow - holds data on the monitored workflow and data on all related metrics
* Api - sents workflow data to Processor-mock over a tcp network

## Input
None

## Output
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