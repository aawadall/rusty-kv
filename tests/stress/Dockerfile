# Build test harness based on ddosify
FROM ddosify/ddosify:latest

# Workdir
WORKDIR /run

# Copy test script with executable permissions
COPY . .

# Set permission to +x 
RUN chmod +x ./tests.sh

# inspect file 
RUN ls -l

