### Univerisad de San Carlos de Guatemala
### Laboratorio de Sistemas Operativos 1
### María Isabel Masaya Córdova
### 201800565

#### Introducción
El proyecto consiste en la implementación de un sistema distribuido de votaciones para un concurso de bandas de música guatemalteca. El objetivo principal es enviar tráfico por medio de archivos con votaciones hacia distintos servicios desplegados en Kubernetes. Estos servicios se encargarán de encolar los datos, almacenarlos en bases de datos y visualizarlos en tiempo real a través de dashboards. Se utilizarán tecnologías como gRPC, Web Assembly (Wasm), Kafka, Redis, MongoDB, Grafana y Cloud Run para lograr este objetivo.

#### Objetivos
1. Implementar un sistema distribuido con microservicios en Kubernetes.
2. Utilizar sistemas de mensajería para encolar y distribuir datos entre servicios.
3. Utilizar Grafana como interfaz gráfica de dashboards para visualizar datos en tiempo real.
4. Desplegar una API en Node.js y una webapp con Vue.js en Cloud Run para consultar registros de MongoDB.

#### Descripción de Tecnologías Utilizadas
1. **Kubernetes**: Kubernetes es una plataforma de código abierto diseñada para automatizar el despliegue, escalado y manejo de aplicaciones en contenedores. En este proyecto, Kubernetes se utiliza como plataforma de orquestación para gestionar y desplegar los diferentes servicios del sistema distribuido, garantizando su disponibilidad, escalabilidad y confiabilidad.
2. **Kafka**: Apache Kafka es una plataforma de streaming distribuido que se utiliza para la ingestión, almacenamiento y procesamiento de flujos de datos en tiempo real. En este proyecto, Kafka actúa como sistema de mensajería para encolar y distribuir datos entre los servicios productores (gRPC y Wasm) y el consumidor. Permite una comunicación asíncrona y tolerante a fallos entre los distintos componentes del sistema.
3. **Redis**: Redis es una base de datos en memoria de código abierto que se utiliza para almacenar datos de forma rápida y eficiente. En este proyecto, Redis se emplea para almacenar los contadores de votaciones que los consumidores envían en tiempo real. Proporciona un almacenamiento de clave-valor de alto rendimiento y permite la consulta y actualización rápida de datos.
4. **MongoDB**: MongoDB es una base de datos NoSQL de documentos que se utiliza para almacenar datos de forma flexible y escalable. En este proyecto, MongoDB se utiliza para almacenar los logs generados por el sistema. Proporciona una forma eficiente de almacenar y consultar grandes volúmenes de datos no estructurados, facilitando el análisis y la visualización de los registros de actividad.
5. **Grafana**: Grafana es una plataforma de análisis y visualización de datos de código abierto que se utiliza para crear dashboards y gráficos interactivos. En este proyecto, Grafana se emplea como interfaz gráfica de dashboards para visualizar los contadores de votaciones en tiempo real almacenados en Redis. Permite crear visualizaciones personalizadas y proporciona herramientas avanzadas para el monitoreo y análisis de datos.
6. **Cloud Run**: Plataforma de Google Cloud para desplegar servicios en contenedores de forma escalable y gestionada.

#### Descripción de Deployment y Service de Kubernetes
- **Deployment de gRPC y Wasm**: Este Deployment se encarga de desplegar los servicios productores de gRPC y Web Assembly (Wasm). Ambos servicios están empaquetados en contenedores Docker y ejecutados en pods gestionados por Kubernetes. Estos pods pueden escalar horizontalmente según la carga de trabajo para manejar picos de tráfico.
- **Deployment de Consumidor**: Este Deployment despliega el daemon consumidor, que es responsable de procesar los datos recibidos de los servicios productores. Estos datos se almacenan en bases de datos como Redis y MongoDB para su posterior análisis y consulta. El número de réplicas del consumidor puede ajustarse automáticamente mediante el autoescalado de Kubernetes para garantizar un rendimiento óptimo del sistema.
- **Service de Kafka**: Este Service define una interfaz de red para acceder al servidor de Kafka desde otros componentes del sistema distribuido. Permite que los servicios productores y consumidores se comuniquen de manera eficiente con el servidor de Kafka para enviar y recibir mensajes en la cola.
- **Service de Redis y MongoDB**: Estos Services proporcionan una capa de abstracción para acceder a las bases de datos Redis y MongoDB desde cualquier parte del sistema distribuido. Permiten que los servicios productores, consumidores y otros componentes interactúen con las bases de datos de manera transparente, independientemente de su ubicación en el clúster de Kubernetes.

#### Ejemplo de Funcionamiento
[Ver Ejemplo de Funcionamiento](https://drive.google.com/drive/folders/1p6Za_-bzwUVVi37riiQbuLadXKGuGnm0?usp=drive_link)

En este ejemplo, se muestra el dashboard de Grafana con dos gráficas que visualizan los contadores de votaciones en tiempo real almacenados en Redis. Los datos son actualizados automáticamente y proporcionan una visualización dinámica del flujo de votaciones durante el concurso de bandas.

#### Conclusiones
El proyecto ha logrado implementar con éxito un sistema distribuido de votaciones utilizando tecnologías modernas como Kubernetes, Kafka, Redis, MongoDB y Grafana. Se ha demostrado la viabilidad y eficacia de este enfoque para gestionar y procesar grandes volúmenes de datos en tiempo real. La arquitectura modular y escalable permite adaptarse a diferentes escenarios y requerimientos, proporcionando una base sólida para futuros desarrollos y mejoras en el sistema.


