# JobSphere 🌐

Connecting job seekers and employers through an efficient backend in **Go** and an engaging frontend in **React.js**.

## 🚀 Features

- **User-Friendly Interface**: Enjoy a seamless and intuitive navigation experience.
- **High Performance**: Built using Go for fast and efficient backend processing.
- **Responsive Design**: Tailored for all devices using React.js.
- **Job Filtering**: Easily search and filter jobs to find the perfect opportunity.
- **Employer Dashboard**: A dedicated space for employers to manage job listings and applications.

## 💻 Tech Stack

- **Frontend**: 
  - React.js
  - Redux (for state management)
  - CSS/SCSS (for styling)

- **Backend**:
  - Go
  - Gin (web framework)
  - PostgreSQL (database)

## 🌟 Getting Started

### Prerequisites

Make sure you have the following tools installed:

- [Node.js](https://nodejs.org/) (for frontend)
- [Go](https://golang.org/dl/) (for backend)
- [PostgreSQL](https://www.postgresql.org/download/) (for database)

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/V4T54L/JobSphere.git
   cd JobSphere
   ```

2. **Setup Frontend**:
   ```bash
   cd frontend
   npm install
   npm start
   ```

3. **Setup Backend**:
   ```bash
   cd backend
   go mod tidy
   go run main.go
   ```

### Environment Variables

Make sure to set your environment variables for the backend. Create a `.env` file in the backend directory and include:
```
DATABASE_URL=your_database_url
```

## 📄 Contributing

We welcome contributions! If you have suggestions for improvements, feel free to create an issue or submit a pull request.

1. Fork the repo
2. Create your feature branch
   ```bash
   git checkout -b feature/YourFeature
   ```
3. Commit your changes
   ```bash
   git commit -m 'Add some feature'
   ```
4. Push to the branch
   ```bash
   git push origin feature/YourFeature
   ```
5. Open a Pull Request

**Happy job searching with JobSphere!** 🌟
