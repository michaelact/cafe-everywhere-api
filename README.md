# Cafe Everywhere

Welcome to Cafe Everywhere, API with purpose to connecting Baristas and Customers. This documentation will guide you through setting up and building Cafe Everywhere.

> Cafe Everywhere is a toy project and was only made in less than 2 days, for a college assignment.

## Quick Start

Follow these simple steps to get Cafe Everywhere up and running on your local environment.

### Prerequisites

Make sure you have the following installed on your system:

- [Golang v1.20.4](https://go.dev/dl/go1.20.4.linux-amd64.tar.gz)
- [Rancher Desktop](https://rancherdesktop.io/) or [Docker](https://docs.docker.com/engine/install/)

### Setup

1. **Clone the Repository:**
   ```shell
   git clone https://github.com/michaelact/cafe-everywhere-api
   cd cafe-everywhere-api
   ```

2. **Start the Database:**
   ```shell
   docker compose -f docker-compose.dev.yml up -d
   ```

3. **Perform Database Migration:**
   - Access [Adminer](http://localhost:8080/)
   - Connect to PostgreSQL with the following details:
     - **System:** PostgreSQL
     - **Database:** database
     - **Username:** dev
     - **Password:** HaloPassword138
   - Create a new database named `cafe-everywhere`
   - Execute queries from `database/migrations/`

4. **Install Dependencies:**
   ```shell
   go install
   ```

5. **Set Environment Variables:**
   ```shell
   source .env.example
   ```

6. **Run the API:**
   ```shell
   go run main.go
   ```

Access Cafe Everywhere API at `http://localhost:9999`.

## Build Cafe Everywhere

If you want to build Cafe Everywhere into a single binary file, follow these steps:

1. **Clone the Repository:**
   ```shell
   git clone https://github.com/michaelact/cafe-everywhere-api
   cd cafe-everywhere-api
   ```

2. **Install Dependencies:**
   ```shell
   go install
   ```

3. **Build the Application:**
   ```shell
   go build
   ```

4. **Check the Generated Binary File:**
   - The binary file named `cafe-everywhere` will be in the root directory.

5. **Make it Executable:**
   ```shell
   chmod +x ./cafe-everywhere
   ```

6. **Move it to bin directory:**
   ```shell
   sudo mv ./cafe-everywhere /usr/local/bin/
   ```

Now you can run Cafe Everywhere from anywhere in your terminal by simply typing `cafe-everywhere`.

## License

Cafe Everywhere is licensed under the [MIT License](./LICENSE).

## Author

Cafe Everywhere is created by Michael Act.
