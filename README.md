# Gestión de Biblioteca - Endpoints

### INTEGRANTES
MARCELO 

ERICK ANDRADE 

JOHAN QUINATOA

A continuación se detallan los endpoints disponibles para la aplicación de gestión de biblioteca desarrollada en Go:

- **Página Principal**
  - **URL:** `http://localhost:8080/`
  - **Descripción:** Página principal que muestra enlaces a las secciones de autores y libros.

- **Listado de Autores**
  - **URL:** `http://localhost:8080/autores`
  - **Descripción:** Muestra un listado de todos los autores registrados en la biblioteca.

- **Formulario para Agregar un Nuevo Autor**
  - **URL:** `http://localhost:8080/autores/nuevo`
  - **Descripción:** Formulario donde se puede ingresar información para agregar un nuevo autor a la biblioteca.

- **Listado de Libros**
  - **URL:** `http://localhost:8080/libros`
  - **Descripción:** Muestra un listado de todos los libros registrados en la biblioteca.

- **Formulario para Agregar un Nuevo Libro**
  - **URL:** `http://localhost:8080/libros/nuevo`
  - **Descripción:** Formulario donde se puede ingresar información para agregar un nuevo libro a la biblioteca.

## Instrucciones de Uso

1. **Inicio del Servidor**
   - Ejecuta el servidor utilizando el siguiente comando en la raíz del proyecto:
     ```
     go run main.go
     ```
   - El servidor estará disponible en `http://localhost:8080`.

2. **Acceso a las Funcionalidades**
   - Abre un navegador web y visita las siguientes URLs para acceder a las diferentes funcionalidades:
     - `http://localhost:8080/` - Página principal con enlaces a autores y libros.
     - `http://localhost:8080/autores` - Listado de autores.
     - `http://localhost:8080/autores/nuevo` - Formulario para agregar un nuevo autor.
     - `http://localhost:8080/libros` - Listado de libros.
     - `http://localhost:8080/libros/nuevo` - Formulario para agregar un nuevo libro.

3. **Interacción con la Aplicación**
   - En las páginas de listado (`/autores` y `/libros`), se pueden visualizar los registros actuales.
   - En las páginas de formulario (`/autores/nuevo` y `/libros/nuevo`), se pueden ingresar nuevos datos y guardarlos en la base de datos.

## Tecnologías Utilizadas

- **Backend:**
  - Go (Golang)
  - GORM (biblioteca de ORM para Go)
  - PostgreSQL (base de datos relacional)

- **Frontend:**
  - HTML
  - CSS (estilos básicos)

## Notas Adicionales

- Asegúrate de tener PostgreSQL instalado y configurado con una base de datos llamada `Biblioteca`.
- Los archivos de plantillas HTML se encuentran en la carpeta `templates/`.
- Los estilos CSS se encuentran en la carpeta `static/`.
