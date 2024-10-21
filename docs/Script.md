### **Slide 1: Title Slide**

**Script**:  
"Hello everyone, today I’ll be presenting a project on developing a web forum using **Go**, **SQLite**, **Docker**, and **HTML Templating**. Let's explore the technical details of how these technologies come together."

---

### **Slide 2: Introduction**

**Script**:  
"This project aims to create a web forum where users can register, log in, create posts, comment, and interact with likes and dislikes. We’re using **Go** for the backend, **SQLite** for database management, **Go templates** for HTML rendering, and **Docker** for deployment."

---

### **Slide 3: Project Architecture**

**Script**:  
"The architecture consists of:

- **Go** for handling HTTP requests and logic.
- **SQLite** for storing users, posts, and comments.
- **Go templates** for dynamic HTML.
- **Docker** to ensure consistency across environments."

---

### **Slide 4: HTML Templating in Go**

**Script**:  
"Go's templating engine separates logic from presentation. We reuse components like headers and footers across multiple pages. This reduces redundancy and makes the code easier to manage."

---

### **Slide 5: Root Handler & RenderTemplate**

**Script**:  
"The **root handler** manages sessions, fetches data like posts, and renders the template. The **renderTemplate** function centralizes this rendering, which makes it scalable and simplifies error handling."

---

### **Slide 6: Authentication and Session Management**

**Script**:  
"We use **bcrypt** to securely hash passwords, and sessions are handled with **UUIDs** stored in cookies. These cookies are **HttpOnly** and **Secure**, ensuring session safety."

---

### **Slide 7: Database Setup and Queries**

**Script**:  
"SQLite is used for managing data. Tables are created for users, posts, comments, and likes. SQL queries are run using Go’s `database/sql` package, making interactions with the database simple."

---

### **Slide 8: Docker Integration**

**Script**:  
"Docker is used to containerize the app, ensuring consistency across environments. It simplifies deployment by packaging the app and its dependencies together."

---

### **Slide 9: Questions & Answers**

**Script**:

1. **How does Go templating help?** It separates logic from presentation, making it easier to maintain.
2. **Why use Docker?** It ensures the app runs the same across environments.
3. **How are sessions secured?** UUIDs are used, and cookies are **HttpOnly** and **Secure**.
4. **Why SQLite?** It's lightweight and easy to set up.
5. **How does Go templating compare to React?** Server-side rendering is faster for small apps, with better SEO benefits.

---

### **Slide 10: Conclusion & Takeaways**

**Script**:  
"In conclusion, we’ve built a scalable, secure forum using **Go**, **SQLite**, and **Docker**. Future improvements could include adding real-time features like notifications or a search function."
