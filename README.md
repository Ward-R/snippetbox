### Snippetbox

This is a web application for creating and sharing code snippets, built as a guided project to learn core concepts in Go. The project is based on the book *Let's Go Professional* by Alex Edwards. This experience has significantly improved my understanding of topics like file handling, database models, dependency injection, and testing in a more structured way than with dynamic languages like Ruby.

https://lets-go.alexedwards.net/

<img width="600" height="640" alt="SnippetHPPic" src="https://github.com/user-attachments/assets/88cf31dc-95cd-4115-9482-1e9c27dca618" />
<img width="600" height="640" alt="SnippetView" src="https://github.com/user-attachments/assets/51a35f95-0b3a-4087-86a0-2206f15e91b7" />




### Key Features

* **Mocking Dependencies:** The application uses mock database models to facilitate testing without needing a live database connection.
* **Database Integration:** It connects to a MySQL database to handle persistent storage of snippets and user information.
* **User Authentication:** The application includes a user registration and login system.
* **Session Management:** It uses a session manager to handle user sessions, ensuring security and continuity.
* **Routing and Handlers:** The program uses Go's standard `net/http` package to handle routing and process requests.
* **Form Handling:** It includes functionality to handle user-submitted forms, including validation and error handling.
* **Template Rendering:** The application uses Go's `html/template` package to render dynamic HTML pages.

### Built With

* **Go** - Backend Programming Language
* **MySQL** - Database
* **alexedwards/scs** - Session Management
* **go-playground/form** - Form Handling

### Use

To run this application:

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/YOUR_GITHUB_USERNAME/snippetbox.git](https://github.com/YOUR_GITHUB_USERNAME/snippetbox.git)
    ```
2.  **Navigate to the project directory:**
    ```bash
    cd snippetbox
    ```
3.  **Install dependencies and run the application:**
    ```bash
    go run ./cmd/web
    ```
    (Note: You will need a configured MySQL database to run the application in production mode, as well as set the correct command-line flags. The provided code for mocking allows you to run tests without a database.)

### Acknowledgements

This project is based on the book *Let's Go Professional* by Alex Edwards. I coded along with the material to deepen my knowledge of Go. My understanding of the language, particularly with interfaces and building a structured application, has improved greatly. While I may not 100% understand everything in the book yet, the process of building this project has been a valuable and enjoyable learning experience.
