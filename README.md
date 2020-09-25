# HTTP proxy server golang

## HTTP сделано
## HTTPS пытался неоднократно, не получилось
## Сохранение запросов сделано
## Сканнер уязвимости сделано

# How to use on the example of http://info.cern.ch/
* 1) HTTP
- curl -x http://localhost:8081 http://info.cern.ch/

* 3) Сохранение запросов:
 - запрос сохранился автоматически в файле /requests/last_request_info.cern.ch.txt, когда вы делали запрос через curl, поэтому:
 - cd requests
 - ./repeat_request last_request_info.cern.ch.txt
 
* 4) Сканнер уязвизвимости:
- После повторение какого-либо запроса сервер напишет вам о том, что была найдена/не найдена уязвимость
