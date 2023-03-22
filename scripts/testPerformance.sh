#!/bin/bash

# Set the number of parallel requests to 50
parallel_requests=50

# Set the number of requests per parallel request to 100
requests_per_parallel=100

function measure_response_time() {
    local url=$1

    # Send the requests using ab and output the result to a temporary file
    ab -n $((parallel_requests*requests_per_parallel)) -c $parallel_requests $url > timing.txt

    wait

    # Extract the response times from the ab output
    response_times=$(grep "Time per request" timing.txt | awk '{print $4}')

    # Calculate the average, minimum, and maximum response times
    average_time=$(echo "$response_times" | awk '{ total += $1 } END { print total/NR }')
    minimum_time=$(echo "$response_times" | sort -n | head -n1)
    maximum_time=$(echo "$response_times" | sort -n | tail -n1)

    # Print the results
    echo ""
    echo "Average response time: $average_time miliseconds"
    echo "Minimum response time: $minimum_time miliseconds"
    echo "Maximum response time: $maximum_time miliseconds"

}

echo "This script will compare the response times of a website when using cache and when not using cache."
echo ""
echo "First sending 5000 requests to the website without using cache."

# THE URL TO TEST which is not using cache
url="http://localhost:8080/posts/1"
measure_response_time $url 

echo ""
echo "Now sending 5000 requests to the website using cache."

# THE URL TO TEST which is using cache
url="http://localhost:8080/posts/1"
measure_response_time $url

