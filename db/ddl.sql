-- Proyecto 2 - Simulación de Reservas Concurrentes
-- Script de creación de tablas y relaciones (DDL)

-- Eliminar tablas respetando orden de claves foráneas
DROP TABLE IF EXISTS reserva;
DROP TABLE IF EXISTS asiento;
DROP TABLE IF EXISTS evento;
DROP TABLE IF EXISTS usuario;

-- Tabla: usuario
CREATE TABLE usuario (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL CHECK (char_length(nombre) >= 3),
    email VARCHAR(150) NOT NULL UNIQUE CHECK (position('@' in email) > 1)
);

-- Tabla: evento
CREATE TABLE evento (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL CHECK (char_length(nombre) >= 5),
    fecha DATE NOT NULL CHECK (fecha >= CURRENT_DATE),
    lugar VARCHAR(100) NOT NULL
);

-- Tabla: asiento
CREATE TABLE asiento (
    id SERIAL PRIMARY KEY,
    id_evento INTEGER NOT NULL,
    numero_asiento VARCHAR(10) NOT NULL CHECK (char_length(numero_asiento) >= 1),
    zona VARCHAR(50) NOT NULL CHECK (zona IN ('VIP', 'General', 'Económica')),
    FOREIGN KEY (id_evento) REFERENCES evento(id) ON DELETE CASCADE,
    UNIQUE (id_evento, numero_asiento)  -- Evita duplicados de asiento por evento
);

-- Tabla: reserva
CREATE TABLE reserva (
    id SERIAL PRIMARY KEY,
    id_usuario INTEGER NOT NULL,
    id_asiento INTEGER NOT NULL UNIQUE,  -- Solo un usuario puede reservar un asiento
    fecha_reserva TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_usuario) REFERENCES usuario(id) ON DELETE CASCADE,
    FOREIGN KEY (id_asiento) REFERENCES asiento(id) ON DELETE CASCADE
);