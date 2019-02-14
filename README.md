# Rest API

This is a very simple API Rest to test the powerful of Go

## Prerequisites

You should to have instaled Go in yout computer (install from https://golang.org/doc/install)

```
Give examples
```

## Getting Started

Execute the following command to run the application:

```
go build && go rest-api
```

## Enjoy

### Endpoints

- GET       /cars        
- GET       /cars/{id}    
- POST      /cars/{id}
- DELETE    /cars/{id}

Sample body in post to create a car:
```
{
  "Brand": {
    "Name": "Seat",
    "Country": "Spain"
  },
  "Model": "124 Sport",
  "Power": 110
}
```

## Built With

* [Go](https://golang.org/) - The Go language

## Author

* **Oscar Soria** - [LabWeb](http://labweb.es)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Thanks a lot to [Francis Sunday](https://dev.to/codehakase) for his great programming ideas
* Thanks to the great car designers that did so [great beautiful automobiles](https://www.hagerty.com/articles-videos/articles/2018/08/21/7-rarest-hemi-muscle-cars)
