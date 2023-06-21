# Data Persistent Reader

This service, built with Go, receives a text file as input and persists its data into a relational database (PostgreSQL). The service splits the data into columns in the database and performs data cleansing and validation.

## Prerequisites

Before running the service, make sure you have the following installed:

- Docker
- Docker Compose

## Setup and Execution

1. Clone the repository:

```bash
git clone https://github.com/ederrochax/data-persistence-reader.git
```

2. Navigate to the project directory:

```bash
cd data-persistence-data-persistence-reader
```

3. Start and build Docker:

```bash
docker-compose up --build
```