# routes_ms
# Microservicio de Rutas de Vuelo

Este microservicio ayuda a encontrar las mejores escalas posibles para viajar de un punto A a un punto B. Utiliza un algoritmo topológico (toposort) para calcular las rutas óptimas.

## Modelos

El microservicio utiliza un principal modelo:
### Routes (Toposort)

El modelo `Routes` representa un grafo de rutas y almacena el resultado del algoritmo topológico (toposort) para determinar el orden de las escalas en una ruta óptima.

- `NumNodes`: Número de nodos en el grafo.
- `Ordering`: Ordenamiento topológico de los nodos del grafo.

## Cómo funciona

El microservicio funciona de la siguiente manera:

1. Proporciona endpoints de API para crear, obtener, actualizar y eliminar gráficos de rutas (`Routes`).

2. Utiliza un toposort para calcular la mejor secuencia de escalas para viajar de un punto A a un punto B.

3. Almacena los datos en una base de datos de PostgreSQL.

## Requisitos previos

- Docker
- Asegúrate de haber descargado las imágenes de docker anteriormente
  ```cmd
   docker pull felipeperezleal/routes_db
  ```

  ```cmd
   docker pull felipeperezleal/routes_ms
  ```

## Pasos para ejecutar el programa

1. Clona este repositorio:
   ```cmd
    git clone https://github.com/felipeperezleal/routes_ms.git
   ```
2. Ejecutar el contenedor de Docker de la base de datos:
   ```cmd
     docker run -p 5432:5432 felipeperezleal/routes_db 
   ```
3. Ejecutar el contenedor de Docker del microservicio:
   ```cmd
    docker run -p 8081:8081 felipeperezleal/routes_ms
   ```
   
## Endpoints del API

- `GET /routes`: Obtiene la lista de todas las rutas.
- `GET /routes/{id}`: Obtiene un la ruta identificada por su ID.
- `POST /routes`: Crea una nueva ruta.
- `PUT /routes/{id}`: Actualiza una ruta existente por su ID.
- `DELETE /routes/{id}`: Elimina una ruta por su ID.






