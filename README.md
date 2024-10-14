<div align="center">
<br>

![Hello_world-web_app.png](README-image/hello_world-web_app.png)

</div>


<p align="center">
<img src="https://img.shields.io/badge/-GO-darkblue">
<img src="https://img.shields.io/badge/-Linux-lightgrey">
<img src="https://img.shields.io/badge/-WSL-brown">
<img src="https://img.shields.io/badge/-Ubuntu%2020.04.4%20LTS-orange">
<img src="https://img.shields.io/badge/-JetBrains-blue">
<img src="https://img.shields.io/badge/License-not%20specified-brightgreen">
</p>


<h1 align="center"> Hello World! - Web App </h1>


<h3 align="center">
<a href="https://github.com/RazikaBengana/Hello_World-webapp#eye-about">About</a> •
<a href="https://github.com/RazikaBengana/Hello_World-webapp#memo-learning-objectives">Learning Objectives</a> •
<a href="https://github.com/RazikaBengana/Hello_World-webapp#computer-requirements">Requirements</a> •
<a href="https://github.com/RazikaBengana/Hello_World-webapp#package-installation">Installation</a> •
<a href="https://github.com/RazikaBengana/Hello_World-webapp#keyboard-basic-usage">Basic Usage</a> •
<a href="https://github.com/RazikaBengana/Hello_World-webapp#pushpin-more info">More Info</a> •
<a href="https://github.com/RazikaBengana/Hello_World-webapp#mag_right-resources">Resources</a> •
<a href="https://github.com/RazikaBengana/Hello_World-webapp#bust_in_silhouette-authors">Authors</a> •
<a href="https://github.com/RazikaBengana/Hello_World-webapp#octocat-license">License</a>
</h3>

---

<!-- ------------------------------------------------------------------------------------------------- -->

<br>
<br>

## :eye: About

<br>

<div align="center">

**`Hello World! - web app`** is a very simple web application built using `Go`.
<br>
<br>
It serves only as a demonstration of how to create a web server, handle `HTTP` requests, manage user sessions, and render dynamic content using `Go` templates.
<br>
The application features a `home page` and an `about page`, showcasing basic routing and middleware functionalities.

</div>

<br>
<br>

<!-- ------------------------------------------------------------------------------------------------- -->

## :memo: Learning Objectives

<br>

```diff

+ Understand the structure of a Go web application

+ Learn how to manage HTTP requests and responses

+ Explore session management and CSRF protection

+ Gain experience with Go's templating system for rendering dynamic content

```

<br>
<br>

<!-- ------------------------------------------------------------------------------------------------- -->

## :computer: Requirements

<br>

```diff

+ Go version 1.12 or higher

+ A web browser for accessing the application

```

<br>
<br>

<!-- ------------------------------------------------------------------------------------------------- -->

## :package: Installation

<br>

- Clone the repository:

<br>

```yaml
git clone https://github.com/RazikaBengana/Hello_World-webapp.git
```

<br>

- Install dependencies:

<br>

```yaml
go mod tidy
```

<br>
<br>

<!-- ------------------------------------------------------------------------------------------------- -->

## :keyboard: Basic Usage

<br>

- Run the application:

<br>

```yaml
go run cmd/web/main.go
```

<br>

- For the `home` page : 

  - Access it by navigating to [http://localhost:8080](http://localhost:8080)

<br>

- For the `about` page: 

  - Access it by navigating to [http://localhost:8080/about](http://localhost:8080/about)

<br>
<br>

<!-- ------------------------------------------------------------------------------------------------- -->

## :pushpin: More Info

<br>

### Project Architecture:

<br>

The application is designed using the **`MVC` (Model-View-Controller)** architectural pattern. <br>
This pattern separates the application into **three interconnected components**, allowing for organized and maintainable code.

<br>

- **Model**:
  - **Location**: `/pkg/models`
  - The folder contains data structures like `TemplateData`, which are used to pass data from handlers to templates. <br>
    This struct includes maps for different data types and properties for handling user feedback and security tokens.

- **View**:
  - **Location**: `/templates`
  - The folder contains HTML templates for rendering the user interface. <br>
    Templates like `home.page.tmpl` and `about.page.tmpl` define the structure and content for the `home` and `about` pages, respectively, extending a base layout defined in `base.layout.tmpl`.

- **Controller**:
  - **Location**: `/pkg/handlers`
  - It uses a repository pattern to manage application configuration and session data. <br>
    It defines HTTP handlers for processing requests to the `home` and `about` pages.

<br>
<br>

### Folder Structure:

<br>

```yaml
/hello-world-web-app
├── /cmd
│   └── /web
│       ├── main.go                # Entry point for the web application, initializes settings, manages sessions, sets up handlers, and starts the server.
│       ├── middleware.go          # Implements middleware functions for CSRF protection and session management.
│       └── routes.go              # Sets up routing using the Chi router, applying middleware and defining routes.
├── /pkg
│   ├── /config
│   │   └── appconfig.go           # Defines the AppConfig struct for managing application configuration settings.
│   ├── /handlers
│   │   └── handlers.go            # Implements HTTP handlers using a repository pattern to manage configuration and session data.
│   ├── /models
│   │   └── templatedata.go        # Defines data structures like TemplateData for passing data to templates.
│   └── /render
│       └── render.go              # Manages and renders HTML templates, including setting up template caching and adding default data.
└── /templates
    ├── base.layout.tmpl           # Defines a base HTML layout structure for web pages.
    ├── home.page.tmpl             # Defines the structure and content for the home page.
    └── about.page.tmpl            # Defines the structure and content for the about page.
```

<br>
<br>

This separation of concerns ensures that the application is modular, making it easier to develop, test, and maintain.

<br>
<br>

### Services:

<br>

- No external services are required for this application.

<br>
<br>

<!-- ------------------------------------------------------------------------------------------------- -->

## :mag_right: Resources

<br>

**_Do you need some help?_**

<br>

**Read or watch:**

* [Go - Documentation](https://go.dev/doc/)

* [SCS - Session Management](https://github.com/alexedwards/scs)

* [Go - Templates](https://pkg.go.dev/text/template)

<br>
<br>

<!-- ------------------------------------------------------------------------------------------------- -->

## :bust_in_silhouette: Authors

<br>

**${\color{blue}Razika \space Bengana}$**

<br>
<br>

<!-- ------------------------------------------------------------------------------------------------- -->

## :octocat: License

<br>

```Hello World! - web app``` _project has no license specified._

<br>
<br>

---

<p align="center"><br>2024</p>
