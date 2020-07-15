# Gold Notification Documents

## Demo

GET: [https://goldnoti.herokuapp.com/api/health](https://goldnoti.herokuapp.com/api/health)
POST: [https://goldnoti.herokuapp.com/api/today](https://goldnoti.herokuapp.com/api/today)

## Postman Collection

[Goldnoti.postman_collection.json](./Goldnoti.postman_collection.json)

## Health Check

HTTP Request

Method: `*`
Path: `http://{{URL}}:{{PORT}}/api/health`

Request: (No Content)

Response:
  
- Http Status: 200 OK

  ```json
  {
    "project_name": "goldnoti",
    "status": "I'm OK.",
    "version": "0.1.0",
    "env":"dev",
    "request_timestamp": "2020-07-15T11:18:02Z"
  }
  ```

## Get Today Price

HTTP Request

Method: `POST`
Path: `http://{{URL}}:{{PORT}}/api/today`

Request: (No Content)

Response:
  
- Http Status: 200 OK

  ```json
  {
    "response_data": {
      "bar_buy": 26850,
      "bar_sell": 26950,
      "ornament_buy": 26363.24,
      "ornament_sell": 27450,
      "status_change": "ทองขึ้น",
      "today_change": 50,
      "updated_date": "15 กรกฎาคม 2563",
      "updated_time": "เวลา 09:20 น."
    },
    "response_message": "",
    "response_timestamp": "2020-07-15T11:05:07Z"
  }
  ```
