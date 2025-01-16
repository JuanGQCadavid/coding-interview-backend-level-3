# Bienvenido al coding-interview-backend-level-3

PROD: https://interview.pululapp.com/items

Postman: See email for invite link

## Service

The Golang service was coded using Hexagonal architecture.

## Infraestructure

All the infra is IaaC, the only component that was not as code is the RDS, as it is one already created for my personal projects.

The service is a docker image deployed on an amd64 lambda using ECR; it is exposed using an API gateway and route53 for the DNS, and Certificate manager for the SSL certificate.

![interview-Infra (1)](https://github.com/user-attachments/assets/18b21f9d-b5b0-42ea-a1c5-31ecf648d7fc)

## CI/CD

The CI/CD is triggered automatically by a change on the main branch; it will detect whether it is an infra or service change and deploy its corresponding workflow.

![interview-CI_cd](https://github.com/user-attachments/assets/2e0379ed-b480-45f8-bcfd-ebba61e258b5)

## Descripción
Este proyecto es una API REST que permite realizar operaciones CRUD sobre una entidad de tipo `Item`.

La entidad tiene 3 campos: `id`, `name` y `price`.

Tu tarea es completar la implementación de toda la funcionalidad de forma tal de que los tests e2e pasen exitosamente.

### Que puedes hacer: 
- ✅ Modificar el código fuente y agregar nuevas clases, métodos, campos, etc.
- ✅ Cambiar dependencias, agregar nuevas, etc.
- ✅ Modificar la estructura del proyecto (/src/** es todo tuyo)
- ✅ Elegir una base de datos
- ✅ Elegir un framework web
- ✅ Cambiar la definición del .devContainer


### Que **no** puedes hacer:
- ❌ No puedes modificar el archivo original /e2e/index.test.ts (pero puedes crear otros e2e test si lo deseas)
- ❌ El proyecto debe usar Typescript 
- ❌ Estresarte 🤗


## Pasos para comenzar
1. Haz un fork usando este repositorio como template
2. Clona el repositorio en tu máquina
3. Realiza los cambios necesarios para que los tests pasen
4. Sube tus cambios a tu repositorio
5. Avísanos que has terminado
6. ???
7. PROFIT

### Cualquier duda contactarme a https://www.linkedin.com/in/andreujuan/
