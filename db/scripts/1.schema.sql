CREATE TABLE municipio(
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    nombre VARCHAR NOT NULL UNIQUE
);

CREATE TABLE parqueadero(
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL UNIQUE,
    capacidad SMALLINT NOT NULL CHECK (capacidad > 0),
    latitud FLOAT(15) NOT NULL,
    longitud FLOAT(15) NOT NULL,
    municipio_id INT NOT NULL,
    CONSTRAINT fk_parqueadero_municipio
      FOREIGN KEY(municipio_id)
	  REFERENCES municipio(id)
);

CREATE TABLE bus(
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    placa VARCHAR(10) NOT NULL UNIQUE,
    capacidad SMALLINT NOT NULL CHECK (capacidad > 0)
);

CREATE TABLE programacion(
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    llegada timestamp NOT NULL,
    salida timestamp NOT NULL,
    estadia INT NOT NULL CHECK (estadia > 0),
    bus_id INT NOT NULL,
    parqueadero_id INT NOT NULL,
    CONSTRAINT fk_programacion_bus
      FOREIGN KEY(bus_id)
	  REFERENCES bus(id),
    CONSTRAINT fk_programacion_parqueadero
      FOREIGN KEY(parqueadero_id)
	  REFERENCES parqueadero(id)
);