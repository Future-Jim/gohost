# gohost

Gohost is a toy monitoring solution that aggregates host data and makes it accessible via a local webserver. 

There are three components required to run gohost

1. gohost-mon binary
   * Gohost-mon binary needs to run on the host that will be monitored. 
   * Aggregates host metrics into the corresponding db
2. gohost-api binary
   * User facing webserver that serves monitoring data via API endpoints
3. postgresql db
   * Stores monitoring data and user credentials

# Endpoints

## Metrics

### <u>*_Get metric by ID_*</u>
### *_Request_*
```
http://localhost:port/metric/{id}
```
### *_Response_*
```
{
    "id": int,
    "al_1": float64,
    "al_5": float64,
    "al_15": float64,
    "hutd": uint64,
    "huth": uint64,
    "hutm": uint64,
    "pmu": int,
    "createdAt": time.Time
}
```
### <u>*_Get metric by date range_*</u>
### *_Request_*
```
http://localhost:port/metrics
```
### __Paramters__:

#### *_Header_*:

```
x-jwt-token: JWT-token
```

#### *_Body_*: 
```
{
    "startDateTime": "YYYY-MM-DDTHH:MM:SSZ",
    "endDateTime": "YYYY-MM-DDTHH:MM:SSZ"
}
```

Note: the character T and Z are required and should not be changed. 
### *_Response_*
```
{
    "id": int,
    "al_1": float64,
    "al_5": float64,
    "al_15": float64,
    "hutd": uint64,
    "huth": uint64,
    "hutm": uint64,
    "pmu": int,
    "createdAt": time.Time
}
```
