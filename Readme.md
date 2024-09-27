
# GPS BUS TRACKER SIMULATION

This app is a web socket that have two end point, one to register coordenates of a client and the other to get all the locations of the buses

## Tech Stack

**Golang, PostgreSQL, Docker, Docker-Compose, DevContainers**

## Environment Variables

To run this project locally, you will need to add the following environment variables to your .env file

#### api

- `DB_HOST="postgres"`
- `DB_PORT="5432"`
- `DB_USER="postgres"`
- `DB_PASSWORD="example"`
- `DB_NAME="postgres"`
- `API_SOCKETURL="gps:81"`

#### DB

`POSTGRES_PASSWORD="example"`

## Run Locally

Clone the project

```bash
  git clone git@github.com:davidsgv/Parqueadero-back.git
```

Go to the project directory

```bash
  cd Parqueadero-back
```

Run compose

``` bash
    cd build
    docker compose up
```

**api:** runs on port 80

**gps:** runs on port 81

## API Reference

#### Get all municipios

```http
  GET /api/municipios
```

#### Get parqueaderos

```http
  GET /api/parqueaderos
```

### Post parqueadero

```http
  POST /api/parqueaderos
```

Request Body:

```json
{
    "nombre": "nombre",
    "capacidad": 10,
    "longitud": 15.5522,
    "latitud": 41.1115,
    "municipio": {
        "id": 1
    }
}
```

### Get all buses

```http
  GET /api/buses
```

### Post bus

```http
  POST /api/buses
```

Request Body:

```json
{
    "placa": "abc123",
    "capacidad": 15
}
```

### Get all programas

```http
  GET /api/programas
```

### Post programa

```http
  POST /api/programas
```

Request Body:

```json
{
    "llegada": 1697475384,
    "salida": 1697475384,
    "estadia": 60,
    "bus": 1,
    "parqueaderoSalida": 1,
    "parqueaderoLlegada": 2
}
```

### Simulate movement

```http
  POST /api/simulate
```

simulate multiple buses to start tracking (check the front project)

## Socket Reference

### Send GPS data

```http
  ws /gps
```

send bus location

Request Body:

```json
{
    "latitude": 4.677584,
    "longitude": -74.1478,
    "plate": "xyz123"
}
```

### Get GPS data

```http
  ws /getgps
```

get bus live locations

## Doc

[Excalidraw Diagram](https://excalidraw.com/#json=czxgP2yhu8F2bkc6MkZQE,4a6E030q_r0bydWHYPXO3g)


## Appendix

[Front Project](https://github.com/davidsgv/Parqueadero-front).
