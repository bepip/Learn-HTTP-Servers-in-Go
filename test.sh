#!/bin/bash

# Define test cases as an array of JSON objects
test_cases=(
  '{"body":"This is a short valid chirp"}|{"cleaned_body":"This is a short valid chirp"}'
  '{"body":"This is a very long chirp that exceeds the maximum character limit. It goes on and on and on and on and on and on and on and on and on and on and on and on and on and on and on and on and on and should be invalid."}|{"error":"Chirp is too long"}'
  '{"body":"This is a kerfuffle opinion I need to share with the world"}|{"cleaned_body":"This is a **** opinion I need to share with the world"}'
  '{"body":"This is a FornAx opinion I need to share with the world"}|{"cleaned_body":"This is a **** opinion I need to share with the world"}'
  '{"body":"This is a Sharbert Fornax kerfuffle! opinion I need to share with the world"}|{"cleaned_body":"This is a **** **** kerfuffle! opinion I need to share with the world"}'
)

# Initialize counters for scoring
total_tests=0
passed_tests=0

# Loop through each test case
for test_case in "${test_cases[@]}"; do
  # Split test case into input (before |) and expected output (after |)
  IFS='|' read -r input expected <<< "$test_case"

  # Make the API call
  response=$(curl -s -X POST http://localhost:8080/api/validate_chirp \
    -H "Content-Type: application/json" \
    -d "$input")

  # Print test details
  echo "Input: $input"
  echo "Response: $response"
  echo "Expected: $expected"

  # Compare response with expected output
  if [[ "$response" == "$expected" ]]; then
    echo "Test Passed ✅"
    ((passed_tests++))
  else
    echo "Test Failed ❌"
  fi

  echo ""
  ((total_tests++))
done

# Print final score
echo "Total Tests: $total_tests"
echo "Passed Tests: $passed_tests"
echo "Failed Tests: $((total_tests - passed_tests))"
echo "Score: $((passed_tests * 100 / total_tests))%"
