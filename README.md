### CineMind — AI-Powered Movie Streaming Platform

A full-stack movie streaming application built with Go (Gin-Gonic) for the backend, React and Bootsrap for the frontend, and MongoDB for data storage.  
The platform includes an AI-powered movie recommendation system powered by LangChain-Go and OpenAI model, which analyzes user behavior and preferences to recommend personalized movies.

---

#### Core Platform
- User Authentication – Signup, login, JWT-based session management  
- Movie Streaming – Play, pause, resume, and track watched movies  
- Categorization – Filter movies by genre, popularity, release date, or rating  
- Search System – Fast and dynamic movie search  

#### AI Integration
- LangChain-Go AI Recommender
  - Analyzes user watch history, preferences, and ratings  
  - Generates intelligent recommendations using embeddings & similarity search  
  - Can integrate with OpenAI or other LLM APIs via LangChain-Go  
---

#### Tech Stack

| Layer | Technology | Description |
|-------|-------------|-------------|
| **Frontend** | React, TailwindCSS | UI for browsing, watching, and managing movies |
| **Backend** | Go (Gin-Gonic) | REST API for authentication, movie CRUD, and streaming logic |
| **AI Layer** | LangChain-Go | Intelligent movie recommendation engine |
| **Database** | MongoDB | Stores user data, movie metadata, and embeddings |
| **Auth** | JWT + bcrypt | Secure token-based authentication |
| **Streaming** | Static file serving + presigned URLs | Handles video streaming from local storage or cloud |

---
