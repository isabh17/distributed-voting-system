### Documentación del Proyecto

#### Introducción
El proyecto consiste en la implementación de un sistema distribuido de votaciones para un concurso de bandas de música guatemalteca. El objetivo principal es enviar tráfico por medio de archivos con votaciones hacia distintos servicios desplegados en Kubernetes. Estos servicios se encargarán de encolar los datos, almacenarlos en bases de datos y visualizarlos en tiempo real a través de dashboards. Se utilizarán tecnologías como gRPC, Web Assembly (Wasm), Kafka, Redis, MongoDB, Grafana y Cloud Run para lograr este objetivo.

#### Objetivos
1. Implementar un sistema distribuido con microservicios en Kubernetes.
2. Utilizar sistemas de mensajería para encolar y distribuir datos entre servicios.
3. Utilizar Grafana como interfaz gráfica de dashboards para visualizar datos en tiempo real.
4. Desplegar una API en Node.js y una webapp con Vue.js en Cloud Run para consultar registros de MongoDB.

#### Descripción de Tecnologías Utilizadas
1. **Kubernetes**: Plataforma de orquestación para gestionar y desplegar microservicios.
2. **Kafka**: Sistema de mensajería para encolar y distribuir datos entre servicios.
3. **Redis**: Base de datos en memoria para almacenar contadores de votaciones en tiempo real.
4. **MongoDB**: Base de datos NoSQL para almacenar logs generados por el sistema.
5. **Grafana**: Plataforma de análisis y visualización de datos para crear dashboards interactivos.
6. **Cloud Run**: Plataforma de Google Cloud para desplegar servicios en contenedores de forma escalable y gestionada.

#### Descripción de Deployment y Service de Kubernetes
- **Deployment de gRPC y Wasm**: Despliega los servicios productores de gRPC y Web Assembly (Wasm) en pods gestionados por Kubernetes.
- **Deployment de Consumidor**: Despliega el daemon consumidor en pods con autoescalado para gestionar la recepción y almacenamiento de datos en bases de datos.
- **Service de Kafka**: Define un servicio para acceder al servidor de Kafka desde otros componentes del sistema distribuido.
- **Service de Redis y MongoDB**: Define servicios para acceder a las bases de datos Redis y MongoDB desde los distintos componentes del sistema.

#### Ejemplo de Funcionamiento
[Ver Ejemplo de Funcionamiento](https://drive.google.com/drive/folders/1p6Za_-bzwUVVi37riiQbuLadXKGuGnm0?usp=drive_link)

En este ejemplo, se muestra el dashboard de Grafana con dos gráficas que visualizan los contadores de votaciones en tiempo real almacenados en Redis. Los datos son actualizados automáticamente y proporcionan una visualización dinámica del flujo de votaciones durante el concurso de bandas.

#### Conclusiones
El proyecto ha logrado implementar con éxito un sistema distribuido de votaciones utilizando tecnologías modernas como Kubernetes, Kafka, Redis, MongoDB y Grafana. Se ha demostrado la viabilidad y eficacia de este enfoque para gestionar y procesar grandes volúmenes de datos en tiempo real. La arquitectura modular y escalable permite adaptarse a diferentes escenarios y requerimientos, proporcionando una base sólida para futuros desarrollos y mejoras en el sistema.


