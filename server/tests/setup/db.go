package setup

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbName     = "testdb"
	dbUser     = "testuser"
	dbPassword = "testpass"
	dbPort     = "5432"
)

var (
	dockerClient *client.Client
	containerId  string
	testDB       *gorm.DB
)

// SetupTestDB initializes a test database using Docker.
// It pulls a Postgres image, creates and starts a container, and sets up the database.
// It also runs any necessary migrations for the provided models.
//
// Parameters:
//   - ctx: The context for database operations
//   - models: The database models to be migrated
//
// Returns:
//   - *gorm.DB: A pointer to the initialized database connection
//   - error: An error if any step in the setup process fails
func SetupTestDB(ctx context.Context, models ...interface{}) (*gorm.DB, error) {
	var err error
	dockerClient, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to create docker client: %v", err)
	}

	var options image.PullOptions
	_, err = dockerClient.ImagePull(ctx, "docker.io/library/postgres:13", options)
	if err != nil {
		return nil, fmt.Errorf("failed to pull postgres image: %v", err)
	}

	config := &container.Config{
		Image: "postgres:13",
		Env: []string{
			fmt.Sprintf("POSTGRES_DB=%s", dbName),
			fmt.Sprintf("POSTGRES_USER=%s", dbUser),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", dbPassword),
		},
		ExposedPorts: nat.PortSet{
			"5432/tcp": struct{}{},
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"5432/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: dbPort,
				},
			},
		},
	}
	resp, err := dockerClient.ContainerCreate(ctx, config, hostConfig, nil, nil, "")
	if err != nil {
		return nil, fmt.Errorf("failed to create container: %v", err)
	}
	containerId = resp.ID

	err = dockerClient.ContainerStart(ctx, containerId, container.StartOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %v", err)
	}

	if err := waitForDB(); err != nil {
		return nil, fmt.Errorf("database not ready: %v", err)
	}

	dsn := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbPort, dbUser, dbPassword, dbName)
	testDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := testDB.AutoMigrate(models...); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}

	return testDB, nil
}

// waitForDB attempts to connect to the database for up to 30 seconds.
// It returns an error if the database is not ready after 30 attempts.
func waitForDB() error {
	dsn := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbPort, dbUser, dbPassword, dbName)

	for i := 0; i < 30; i++ {
		db, err := sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				db.Close()
				return nil
			}
		}
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("timeout waiting for database to be ready")
}

// CleanupTestDB closes the database connection and removes the Docker container.
// It should be called after all tests are completed to clean up resources.
//
// Parameters:
//   - ctx: The context for cleanup operations
func CleanupTestDB(ctx context.Context) {
	if testDB != nil {
		db, _ := testDB.DB()
		db.Close()
	}

	if containerId != "" {
		timeout := 10
		stopOpts := container.StopOptions{
			Timeout: &timeout,
		}
		err := dockerClient.ContainerStop(ctx, containerId, stopOpts)
		if err != nil {
			log.Printf("Failed to stop container: %v", err)
		}

		removeOpts := container.RemoveOptions{
			Force: true,
		}
		err = dockerClient.ContainerRemove(ctx, containerId, removeOpts)
		if err != nil {
			log.Printf("Failed to remove container: %v", err)
		}
	}

	if dockerClient != nil {
		dockerClient.Close()
	}
}
