# Proyecto2025-Demo-Baudino-Ruival-Carpineti

Sistema de gestión de gimnasio desarrollado con arquitectura de microservicios, incluyendo frontend en React y backend en Go con base de datos MySQL.

## 🚀 Tecnologías Utilizadas

- **Frontend**: React.js
- **Backend**: Go (Gin Framework)
- **Base de Datos**: MySQL
- **Contenedores**: Docker & Docker Compose

## 📋 Requisitos Previos

- Docker Desktop instalado y ejecutándose
- Docker Compose (incluido con Docker Desktop)

## 🛠️ Instalación y Ejecución

### 1. Clonar el Repositorio
```bash
git clone <https://github.com/UCC-ArqSoft1/Proyecto2025-Demo-Baudino-Ruival-Carpineti.git>
cd Proyecto2025-Demo-Baudino-Ruival-Carpineti
```

### 2. Levantar la Aplicación
```bash
docker-compose up --build
```

### 3. Esperar a que los Servicios se Inicialicen
Verás logs similares a:
```
[+] Running 3/3
 ✔ Container proyecto2025-demo-baudino-ruival-carpineti-db-1        Started
 ✔ Container proyecto2025-demo-baudino-ruival-carpineti-backend-1   Started
 ✔ Container proyecto2025-demo-baudino-ruival-carpineti-frontend-1  Started
```

## 🌐 Acceso a la Aplicación

Una vez que todos los servicios estén corriendo:

- **Frontend (Interfaz de Usuario)**: http://localhost:3000
- **Backend (API)**: http://localhost:8080
- **Base de Datos**: localhost:3307 (si necesitas conectarte directamente)

## 📊 Datos de Prueba

El sistema incluye datos de prueba pre-cargados:
- **Usuarios**: 12 usuarios (10 socios, 2 administradores)
- **Actividades**: 3 actividades (Yoga, Spinning, Entrenamiento Funcional)
- **Horarios**: 6 horarios distribuidos en diferentes días

### Credenciales de Prueba
- **Socio**: franco@gmail.com / password123
- **Administrador**: felipe@gmail.com / password123

## 🛑 Detener la Aplicación

Para detener todos los servicios:
```bash
docker-compose down
```

Para detener y eliminar volúmenes (datos de la base de datos):
```bash
docker-compose down -v
```

## 🔧 Comandos Útiles

### Ver logs en tiempo real
```bash
docker-compose logs -f
```

### Ver logs de un servicio específico
```bash
docker-compose logs frontend
docker-compose logs backend
docker-compose logs db
```

### Reconstruir un servicio específico
```bash
docker-compose up --build frontend
docker-compose up --build backend
```

### Ver estado de los contenedores
```bash
docker-compose ps
```

## 📁 Estructura del Proyecto

```
Proyecto2025-Demo-Baudino-Ruival-Carpineti/
├── api/
│   ├── backend/          # Servidor Go
│   │   ├── controllers/  # Controladores de la API
│   │   ├── services/     # Lógica de negocio
│   │   ├── dao/         # Acceso a datos
│   │   ├── domain/      # Modelos de dominio
│   │   ├── db/          # Configuración de base de datos
│   │   └── Dockerfile
│   └── frontend/        # Aplicación React
│       ├── src/         # Código fuente
│       ├── public/      # Archivos públicos
│       └── Dockerfile
├── docker-compose.yml   # Configuración de servicios
├── init_data.sql       # Datos de inicialización
└── README.md
```

## 🔍 Endpoints de la API

- `GET /activities` - Obtener todas las actividades
- `GET /activities/:id` - Obtener actividad por ID
- `GET /activities/search` - Buscar actividades
- `POST /login` - Autenticación de usuarios
- `GET /users/:userID/activities` - Actividades de un usuario
- `POST /users/:userID/enrollments` - Inscribir usuario en actividad
- `POST /admin/activities` - Crear actividad (admin)
- `PUT /admin/activities/:id` - Actualizar actividad (admin)
- `DELETE /admin/activities/:id` - Eliminar actividad (admin)

## 🐛 Solución de Problemas

### Puerto 3306 en uso
Si obtienes error de puerto ocupado, el sistema automáticamente usa el puerto 3307 para MySQL.

### Contenedores no se levantan
1. Verifica que Docker Desktop esté ejecutándose
2. Ejecuta `docker-compose down` y luego `docker-compose up --build`
3. Revisa los logs con `docker-compose logs`

### Base de datos sin datos
Si la base de datos está vacía, ejecuta:
```bash
docker exec -i proyecto2025-demo-baudino-ruival-carpineti-db-1 mysql -u root -proot gimnasio < init_data.sql
```

## 📝 Notas Importantes

- **No es necesario tener Go, Node.js o MySQL instalados localmente**
- **Todo funciona dentro de contenedores Docker**
- **Los datos persisten entre reinicios gracias al volumen Docker**
- **La aplicación está lista para desarrollo y pruebas**

## 👥 Autores

- Demo
- Baudino
- Ruival  
- Carpineti

---

**¡Disfruta usando el sistema de gestión de gimnasio! 🏋️‍♂️**