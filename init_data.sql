-- Tabla: activities
INSERT INTO activities (id, titulo, descripcion, categoria, instructor, duracion, imagen, estado) VALUES
(1, 'Yoga Matutino', 'Clase de yoga para principiantes', 'Yoga', 'María García', 60, 'yoga.jpg', 'activo'),
(2, 'Spinning Intenso', 'Clase de spinning de alta intensidad', 'Spinning', 'Juan Pérez', 45, 'spinning.jpg', 'activo'),
(3, 'Entrenamiento Funcional', 'Rutina completa para tonificar', 'Funcional', 'Laura Ramírez', 50, 'funcional.jpg', 'activo');

-- Tabla: schedules
INSERT INTO schedules (id, actividad_id, dia_semana, hora_inicio, hora_fin, cupo) VALUES
(1, 1, 'Lunes', '09:00', '10:00', 15),
(2, 1, 'Miércoles', '09:00', '10:00', 15),
(3, 2, 'Martes', '18:00', '18:45', 10),
(4, 2, 'Jueves', '18:00', '18:45', 10),
(5, 3, 'Viernes', '17:00', '17:50', 12),
(6, 3, 'Sábado', '10:00', '10:50', 12);

-- Tabla: users
INSERT INTO users (id, password_hash, email, rol, username) VALUES
(1, 'adc4d83d58802ca312d1585f24d8efc7fb1a993376fb6307a06cda6832df1eec', 'franco@gmail.com', 'socio', 'Franco'),
(2, '0c8a54523c19212fc67a27cf8e3e46bb618bc860b695d0439c7490d0fa7d8d8e', 'lucia@gmail.com', 'socio', 'Lucia'),
(3, '3fb6181ca88d5a26211063f724ee79b01636805b298926548713a17989ab554a', 'mariano@gmail.com', 'socio', 'Mariano'),
(4, 'ad8d83ffd82b5a8ed429e8592b5cb3e6e83a033770868a1a00c6fd1e7fae242c', 'carla@gmail.com', 'socio', 'Carla'),
(5, 'f536fb10ec50ead5cb957530f32cc06a9e2d517b9a31bfe45a3c8e3117c8012c', 'diego@gmail.com', 'socio', 'Diego'),
(6, 'c09213cf20b32fe2436b92ded92caeed5d938b454c0cd94fea66658823c812b3', 'valentina@gmail.com', 'socio', 'Valentina'),
(7, 'af11205043f66fa4770fbf9b03dd861a90c6a2ae049d7b8b759fa5e78d20e3ff', 'matias@gmail.com', 'socio', 'Matias'),
(8, 'bb866952a8675f537773e9854cc9c5b3be33b22bee2e11dbb71fd07f3317c335', 'sofia@gmail.com', 'socio', 'Sofia'),
(9, '90a5d5a8573d7f52021a85d65d98f44c58a2006b9532204039206705a4d31f80', 'julian@gmail.com', 'socio', 'Julian'),
(10, '80d161545383dd77c8703ec50d566ae19f9f6bb5b68b42b26482f1cdf21bf8f2', 'camila@gmail.com', 'socio', 'Camila'),
(11, 'ea14b35aa297b61e95e0692c4ba04b71d38ed6dcfb49de4e37fb4ae39ffa2bbf', 'felipe@gmail.com', 'Administrador', 'Felipe'),
(12, '9ff18ebe7449349f358e3af0b57cf7a032c1c6b2272cb2656ff85eb112232f16', 'maria@gmail.com', 'Administrador', 'Maria'); 