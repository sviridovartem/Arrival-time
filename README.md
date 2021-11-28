# Lunar Arrival-time
## Background
___
Colonizing the Moon is going to be step one to colonize Mars. Phase 3 of the colonization would
require having some sort of economy.
Trading, traveling and other types of business will be the start of creating a sustainable economy.

## Project
Vestiaire Collective would be the one to start with the lunar fashion project. The purpose of the project is to take a fashion stand in space wear.

And as part of the project, we need to estimate the time it takes our shipments to arrive at Lunar Colony

## Task:
>Create a microservice, with an API that takes the earth time in UTC when the shipment left the warehouse, and calculate the time of arrival to Lunar Colony, the time return should be in LTS."

## Facts
+ It takes 1 day from our warehouse to the earth space station(ESP).
+ The trip from the Earth to the Lunar Colony is about 3 days, You could consider it 3 days
exactly.
___
## Software Requirements
+ Go 1.17.3 or above. 
+ Docker & Docker-compose lts

## Installation
Run the following command to start the service:

    make run
## Testing
Run the following command to start the service:

    make test
___

## HTTP/HTTPS Request example:
`http://localhost:8080/v1/GetLunarArrivalTime`

body example:
+ Default body (default arrival time 4 days: (1d) warehouse -> ESP), (3d) ESP -> moon :
```
{"LocalTime": "2020-01-01 10:10:30 +0300"}
```

        
+ Body with another arrival time:
```
{
  "LocalTime": "2020-01-01 10:10:30 +0300",
  "DeliverDays": 5
}
```



## HTTP/HTTPS Response example:
```
{"LunarArrivalTime":"53-01-04 V 03:18:54"}
```


 
