# Task Manager API

## Descripción
Esta es una API para un gestor de tareas, desarrollado con Go, Gorm y Gin.

## Probar
Para términos de prueba puede ejecutar un 
```docker-compose up```
con la información de prueba que se encuentra en el archivo docker-compose.yml
y con el gestor de base de datos que prefiera generar la conexión a la base de datos.
El archivo requests.http contiene algunas de las peticiones que se pueden realizar para probar la API que de encuentran en el archivo instructions.md

## Endpoints Disponibles
Todas las operaciones CRUD para users, tasks y projects

### Proyectos
- **GET /projects**: Obtener todos los proyectos.
  - **Request**:
    ```http
    GET /projects
    ```
  - **Response**:
    ```json
    {
      "message": "Projects found",
      "projects": [
        {
          "id": 6,
          "name": "Desarrollo de App de Fitness Tracker"
        },
        {
          "id": 7,
          "name": ""
        },
        {
          "id": 10,
          "name": "Desarrollo de App de Gestión de Tareas"
        }
      ]
    }
    ```
- **POST /projects**: Crear un nuevo proyecto without engineers.
  - **Request**:
    ```http
    POST /projects
    {
      "Name": "Desarrollo de App de Gestión de Tareas",
      "DeliveryDate": "2023-12-01",
      "Tasks": [
        {
          "Name": "Integración con Dispositivos Wearables",
          "Description": "Desarrollar funcionalidades para sincronización con smartwatches.",
          "EstimateHours": 150
        },
        {
          "Name": "Creación de Rutinas Personalizadas",
          "Description": "Diseñar y programar la creación de rutinas personalizadas.",
          "EstimateHours": 90
        }
      ],
      "Engineers": [
        {
          "Name": "Ana",
          "Surname": "Ruiz"
        }
      ]
    }
    ```
  - **Response**:
    ```json
    {
      "message": "Project created successfully",
      "project": {
        "id": 10,
        "name": "Desarrollo de App de Gestión de Tareas",
        "delivery": "2023-12-01",
        "engineers": [
          {
            "id": 1,
            "name": "Ana",
            "surname": "Ruiz",
            "projectId": 10
          }
        ],
        "tasks": [
          {
            "id": 1,
            "name": "Integración con Dispositivos Wearables",
            "description": "Desarrollar funcionalidades para sincronización con smartwatches.",
            "estimate": 150,
            "projectId": 10
          },
          {
            "id": 2,
            "name": "Creación de Rutinas Personalizadas",
            "description": "Diseñar y programar la creación de rutinas personalizadas.",
            "estimate": 90,
            "projectId": 10
          }
        ]
      }
    }
    ```
- **GET /projects/:id**: Obtener un proyecto específico.
  - **Request**:
    ```http
    GET /projects/10
    ```
  - **Response**:
    ```json
    {
      "message": "Project found",
      "project": {
        "id": 10,
        "name": "Desarrollo de App de Gestión de Tareas",
        "delivery": "2023-12-01",
        "engineers": [
          {
            "id": 1,
            "name": "Ana",
            "surname": "Ruiz",
            "projectId": 10
          }
        ],
        "tasks": [
          {
            "id": 1,
            "name": "Integración con Dispositivos Wearables",
            "description": "Desarrollar funcionalidades para sincronización con smartwatches.",
            "estimate": 150,
            "projectId": 10
          },
          {
            "id": 2,
            "name": "Creación de Rutinas Personalizadas",
            "description": "Diseñar y programar la creación de rutinas personalizadas.",
            "estimate": 90,
            "projectId": 10
          }
        ]
      }
    }
    ```
- **DELETE /projects/:id**: Eliminar un proyecto específico.
  - **Request**:
    ```http
    DELETE /projects/10
    ```
  - **Response**:
    ```json
    {
      "message": "Project deleted"
    }
    ```

### Usuarios
- **GET /users**: Obtener todos los usuarios.
  - **Request**:
    ```http
    GET /users
    ```
  - **Response**:
    ```json
    
      {
      "message": "Users found",
      "users": [
        {
          "id": 1,
          "name": "Ana"
        },
        {
          "id": 2,
          "name": "Juan"
        },
      ]
    }
- **POST /projects/:id/users**: Crear un nuevo usuario con un proyecto específico.
  - **Request**:
    ```http
    POST /projects/10/users
    {
      "Name": "Ana",
      "Surname": "Ruiz"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "User created successfully",
      "engineer": {
        "ID": 13,
        "CreatedAt": "2024-10-20T14:16:11.1512816-05:00",
        "UpdatedAt": "2024-10-20T14:16:11.1512816-05:00",
        "DeletedAt": null,
        "Name": "Gabriela",
        "Surname": "Rojas",
        "ProjectID": 10,
      }
    }
- **GET /users/:id**: Obtener un usuario específico.
  - **Request**:
    ```http
    GET /users/1
    ```
  - **Response**:
    ```json
    {
      "message": "User found",
      "user": {
        "ID": 1,
        "CreatedAt": "2023-12-01T14:16:11.1512816-05:00",
        "UpdatedAt": "2023-12-01T14:16:11.1512816-05:00",
        "DeletedAt": null,
        "Name": "Ana",
        "Surname": "Ruiz",
        "ProjectID": 10,
      }
    }
    ```
- **PUT /projects/:idProject/users/:idUser**: Reasignar un proyecto a un usuario específico.
  - **Request**:
    ```http
    PUT /projects/10/users/1  
    ```
  - **Response**:
    ```json
    {
      "message": "Engineer reassigned to project",
      "engineer": {
        "id": 1,
        "name": "Ana",
        "surname": "Ruiz",
        "project": {
          "id": 10,
          "name": "Desarrollo de App de Gestión de Tareas",
          "delivery": "2023-12-01",
        }
      }
    } 
    ```

### Tareas
- **POST /projects/:id/tasks**: Crear una nueva tarea asociada a un proyecto específico. 

  - **Request**:
    ```http
    POST /projects/10/tasks
    {
      "Name": "Integración con Dispositivos Wearables",
      "Description": "Desarrollar funcionalidades para sincronización con smartwatches.",
      "EstimateHours": 150
    }
    ```
  - **Response**:
    ```json
    {
      "id": 1,
      "name": "Integración con Dispositivos Wearables",
      "description": "Desarrollar funcionalidades para sincronización con smartwatches.",
      "estimate": 150,
      "projectId": 10
    }
    ```


