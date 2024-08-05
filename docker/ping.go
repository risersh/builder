package docker

// Ping checks if the Docker daemon is running and is reachable.
//
// Returns:
//   - true if the Docker daemon is running and is reachable
func Ping() bool {
	client, err := GetClient(GetClientArgs{})
	if err != nil {
		return false
	}

	return client.Ping()
}
