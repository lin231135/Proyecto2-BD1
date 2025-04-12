-- Insertar usuarios
INSERT INTO usuario (nombre, email) VALUES
('Ana Pérez', 'ana.perez@example.com'),
('Luis Gómez', 'luis.gomez@example.com'),
('Carla Torres', 'carla.torres@example.com'),
('Usuario 4', 'usuario4@example.com'),
('Usuario 5', 'usuario5@example.com'),
('Usuario 6', 'usuario6@example.com'),
('Usuario 7', 'usuario7@example.com'),
('Usuario 8', 'usuario8@example.com'),
('Usuario 9', 'usuario9@example.com'),
('Usuario 10', 'usuario10@example.com'),
('Usuario 11', 'usuario11@example.com'),
('Usuario 12', 'usuario12@example.com'),
('Usuario 13', 'usuario13@example.com'),
('Usuario 14', 'usuario14@example.com'),
('Usuario 15', 'usuario15@example.com'),
('Usuario 16', 'usuario16@example.com'),
('Usuario 17', 'usuario17@example.com'),
('Usuario 18', 'usuario18@example.com'),
('Usuario 19', 'usuario19@example.com'),
('Usuario 20', 'usuario20@example.com'),
('Usuario 21', 'usuario21@example.com'),
('Usuario 22', 'usuario22@example.com'),
('Usuario 23', 'usuario23@example.com'),
('Usuario 24', 'usuario24@example.com'),
('Usuario 25', 'usuario25@example.com'),
('Usuario 26', 'usuario26@example.com'),
('Usuario 27', 'usuario27@example.com'),
('Usuario 28', 'usuario28@example.com'),
('Usuario 29', 'usuario29@example.com'),
('Usuario 30', 'usuario30@example.com');


-- Insertar un evento
INSERT INTO evento (nombre, fecha, lugar) VALUES
('Concierto de Primavera', '2025-05-15', 'Auditorio Nacional');

-- Insertar asientos para el evento (id_evento = 1)
INSERT INTO asiento (id_evento, numero_asiento, zona) VALUES
(1, 'A1', 'VIP'),
(1, 'A2', 'VIP'),
(1, 'A3', 'VIP'),
(1, 'B1', 'General'),
(1, 'B2', 'General'),
(1, 'B3', 'General'),
(1, 'C1', 'Económica'),
(1, 'C2', 'Económica'),
(1, 'C3', 'Económica'),
(1, 'C4', 'Económica');