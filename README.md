# routes_ms
# Microservicio de Rutas de Vuelo

Este microservicio ayuda a encontrar las mejores escalas posibles para viajar de un punto A a un punto B. Utiliza un algoritmo topológico (toposort) para calcular las rutas óptimas.

## Modelos

El microservicio utiliza dos modelos principales:

### Flight

El modelo `Flight` representa información sobre vuelos individuales. Contiene los siguientes atributos:

- `Origin`: Ciudad de origen del vuelo.
- `Destination`: Ciudad de destino del vuelo.
- `Duration`: Duración del vuelo en minutos.
- `Price`: Precio del vuelo en la moneda especificada.

### Routes (Toposort)

El modelo `Routes` representa un grafo de rutas y almacena el resultado del algoritmo topológico (toposort) para determinar el orden de las escalas en una ruta óptima.

- `NumNodes`: Número de nodos en el grafo.
- `Ordering`: Ordenamiento topológico de los nodos del grafo.

## Cómo funciona

El microservicio funciona de la siguiente manera:

1. Proporciona endpoints de API para crear, obtener, actualizar y eliminar vuelos (`Flight`) y gráficos de rutas (`Routes`).

2. Utiliza un toposort para calcular la mejor secuencia de escalas para viajar de un punto A a un punto B.

3. Almacena los datos en una base de datos de PostgreSQL.

## Requisitos previos

Asegúrate de tener instalado Docker, Go y PostgreSQL en tu sistema.

## Pasos para ejecutar el programa

1. Clona este repositorio:
   ```cmd
    git clone https://github.com/felipeperezleal/routes_ms.git
   ```
2. Accede al directorio routes_ms/db: 
   ```cmd
    cd routes_ms/db
   ```
3. Construye la imagen Docker de la base de datos:
   ```cmd
    docker build -t routes-db .
   ```
4. Ejecutar el contenedor de Docker de la base de datos:
   ```cmd
    docker run --name routes-db -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 routes-db  
   ```
5. Vuelve al directorio routes_ms: 
   ```cmd
    cd ..
   ```
6. Construye la imagen Docker del microservicio:
   ```cmd
    docker build -t routes-ms .
   ```
7. Ejecutar el contenedor de Docker del microservicio:
   ```cmd
    docker run --name routes-ms -p 8080:8080 routes-ms
   ```
   
## Endpoints del API

- `GET /flights`: Obtiene la lista de todos los vuelos.
- `GET /flights/{id}`: Obtiene un vuelo por su ID.
- `POST /flights`: Crea un nuevo vuelo.
- `PUT /flights/{id}`: Actualiza un vuelo existente por su ID.
- `DELETE /flights/{id}`: Elimina un vuelo por su ID.

- `GET /routes`: Obtiene la lista de todos los gráficos de rutas.
- `GET /routes/{id}`: Obtiene un gráfico de ruta por su ID.
- `POST /routes`: Crea un nuevo gráfico de ruta.
- `PUT /routes/{id}`: Actualiza un gráfico de ruta existente por su ID.
- `DELETE /routes/{id}`: Elimina un gráfico de ruta por su ID.






