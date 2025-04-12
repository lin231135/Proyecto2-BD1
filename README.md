# 🎟️ Simulación de Reservas Concurrentes

Este proyecto fue desarrollado para el curso **Bases de Datos 1 (CC3088)** de la Universidad del Valle de Guatemala. El objetivo es simular múltiples usuarios intentando reservar el mismo asiento en un evento, aplicando conceptos de **transacciones**, **bloqueos** y **niveles de aislamiento** en PostgreSQL usando concurrencia con Go.

---

## 🧰 Requisitos

Asegúrate de tener lo siguiente instalado antes de ejecutar el programa:

- [Go](https://golang.org/dl/) 1.20 o superior
- [PostgreSQL](https://www.postgresql.org/download/) instalado y corriendo localmente

---

## ⚙️ Configuración

En el archivo `main.go` se encuentran las siguientes constantes para conectarse a la base de datos:

```go
const (
    defaultDSN  = "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable"
    proyectoDSN = "postgres://postgres:12345@localhost:5432/Proyecto_2?sslmode=disable" // CAMBIAR EL USUARIO Y CONTRASEÑA SEGÚN SU CONFIGURACIÓN
)
```

🔧 **IMPORTANTE:** Cambia `postgres:12345` por tu usuario y contraseña reales de PostgreSQL si son distintos.

---

## 🚀 Ejecución

Una vez clonado el repositorio y configurado el DSN correctamente, simplemente ejecuta el siguiente comando en la raíz del proyecto:

```bash
go run .
```

🔄 El programa preguntará si deseas reiniciar la base de datos. Si respondes `s`, ejecutará automáticamente los scripts `ddl.sql` y `data.sql` para crear las tablas y poblar los datos de prueba.

📦 También se encargará de crear la base de datos `Proyecto_2` si no existe, por lo que no necesitas crearla manualmente.

---

## 🗂️ Estructura del Proyecto

```
├── db/
│   ├── ddl.sql          # Script SQL para crear la estructura de la base de datos
│   └── data.sql         # Script SQL para insertar datos de prueba
├── main.go              # Código fuente principal del programa
├── go.mod               # Módulo de dependencias de Go
└── README.md            # Manual de uso
```

---

## 🔍 ¿Qué hace el programa?

- Ejecuta una simulación de múltiples usuarios intentando reservar asientos de manera concurrente.
- Utiliza hilos (`goroutines`) y transacciones en PostgreSQL.
- Prueba los siguientes niveles de aislamiento:
  - READ COMMITTED
  - REPEATABLE READ
  - SERIALIZABLE
- Evalúa el desempeño con distintos niveles de concurrencia (5, 10, 20 y 30 usuarios).
- Registra los resultados para su posterior análisis.

---