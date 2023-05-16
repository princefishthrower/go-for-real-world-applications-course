# Lesson 3

Hourly pollen levels URL:

https://www.pollenwarndienst.at/index.php?eID=appinterface&action=getHourlyLoadData&type=zip&value=6800&country=AT&lang_id=0&pure_json=1&day=0

Example response:

```json
{
  "success": 1,
  "result": {
    "total": 8,
    "dayrisk_personalized": false,
    "hourly": [
      5,
      4,
      4,
      3,
      5,
      4,
      3,
      2,
      3,
      6,
      8,
      8,
      8,
      8,
      8,
      8,
      8,
      8,
      8,
      8,
      8,
      8,
      8
    ]
  }
}
```

Current and historical chart data URL:

https://www.pollenwarndienst.at/index.php?eID=appinterface&action=getCurrentChartData&poll_id=5&region_id=&zip=6800&season=2&lang_id=1&pure_json=1

Example response:

```json
{
  "success": 1,
  "results": [
    {
      "date": "2023-04-05",
      "current": 0.4,
      "average": 0.6,
      "season": "false",
      "datetime": 1680652800000
    },
    {
      "date": "2023-04-06",
      "current": 0.5,
      "average": 0.7,
      "season": "false",
      "datetime": 1680739200000
    },
    {
      "date": "2023-04-07",
      "current": 0.6,
      "average": 0.7,
      "season": "false",
      "datetime": 1680825600000
    },
    {
      "date": "2023-04-08",
      "current": 0.6,
      "average": 0.7,
      "season": "false",
      "datetime": 1680912000000
    },
    {
      "date": "2023-04-09",
      "current": 0.5,
      "average": 0.7,
      "season": "false",
      "datetime": 1680998400000
    },
    {
      "date": "2023-04-10",
      "current": 0.9,
      "average": 0.8,
      "season": "false",
      "datetime": 1681084800000
    }
  ]
}
```