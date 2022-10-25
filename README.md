# problem-company-go
A RESTful API created to apply test on Problem Company

Some libs used **httprouter** (A nice http library), **gorm** (An ORM for Go), **bcrypt** (A generate and check hash passwords).

## Before run the API
```bash
# Run postgres database
docker run -e POSTGRES_PASSWORD=postgres postgres:14.5-alpine
```

## To run the API
```bash
# Build image of the API
docker build -t problem-company-api .

# Run image
docker run problem-company-api
```

Some options to running API server.
```bash
# Configure the port the application will run
docker run -e PORT=<port> problem-company-api
```

## Structure
```
├── pkg
│   ├── db
│   │   └── postgres.go     // Configs for database our application
│   ├── lib
│   │   └── password.go     // Libs for some bussiness rule for our application
│   ├── models
│   │   └── customer.go     // Models for our application
│   ├── routes
│   │   └── routes.go     // Routes for our application
└── main.go
```

## API

#### /customers
* `GET` : Get all customers
* `POST` : Create a new customer

#### /customers/:id
* `GET` : Get a customer
* `PUT` : Update a customer