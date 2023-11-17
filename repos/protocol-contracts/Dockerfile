# Set the base image to ethereum/client-go
FROM ethereum/client-go:alltools-v1.11.5

# Update and install dependencies
RUN apk update && \
    apk add jq unzip bash

# Copy the scripts directory into the docker container
COPY . .

# Make the script executable
RUN chmod +x /scripts/generate_go.sh

# Set the working directory
WORKDIR /

# Print the contents of /scripts/pkg after running the script
CMD ["ls", "-la", "pkg"]

# Set the entry point to your script
ENTRYPOINT ["/bin/bash", "/scripts/generate_go.sh"]
