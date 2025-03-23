# Dungeons & Dragons Character and Quest Management API

## Objective

This project involves the development of a RESTful API for managing characters and quests in a Dungeons & Dragons (D&D) application using the Go programming language and suitable frameworks. The API will serve as the backbone for managing D&D characters and quests, providing features for both public and registered users.

## Features

### 1. Public and Registered User Access
- **Public Users**: All visitors (both registered and non-registered) can fetch public characters and quests.
- **Registered Users**: Registered users can access all characters and quests, including private ones.

### 2. Character and Quest Creation
- **Character Creation**: Registered users can create a character with the following data:
    - Title
    - Description (max 5000 characters)
    - Class (Predefined options)
    - Race (Predefined options)
    - Images (up to 10)
    - Privacy setting (Public/Private)
- **Quest Creation**: Registered users can create a quest with the following data:
    - Title
    - Description (max 5000 characters)
    - Difficulty Level (Predefined options)
    - Images (up to 10)
    - Privacy setting (Public/Private)

### 3. Character and Quest Management
- Registered users can **edit** or **delete** their own characters and quests.

### 4. Admin Control
- Admin users can manage predefined options for character Class, Race, and quest Difficulty Level:
    - Create, Edit, or Delete predefined options.
    - When deleting a predefined option, related characters or quests are archived (not permanently deleted).

### 5. Testing
- Unit tests for business logic are written to ensure correct functionality.

---

## Setup and Running the API

### Prerequisites
- Ensure you have [Docker](https://www.docker.com/) installed.

### Steps

1. **Clone the Project**  
   Start by cloning the repository to your local machine:
   ```bash
        git clone <repository_url>
        cd <project_directory>
   ```
2. **Start the Application with Docker Compose
    The project includes a docker-compose.yml file to simplify the setup. Use the following command to bring up the containers:
    ```
        docker-compose up
    ```
3.	**Access the API
    Once the containers are up and running, you can access the API at:
	•	API: http://localhost:8080/
	•	Swagger Documentation: http://localhost:8080/swagger
4. **User Accounts and Roles
    Two user accounts with different roles have been predefined:
	Admin User:
	•	Username: johndoe
	•	Password: P@ssword1234
	Regular User:
	•	Username: janedoe
	•	Password: P@ssword1234
    These credentials can be used to test the different user roles in the system.

### API Documentation

- The API is fully documented and can be accessed via the Swagger UI. It provides a detailed overview of all available endpoints, their parameters, and response formats. You can find the documentation at the following URL once the server is running on port 8080