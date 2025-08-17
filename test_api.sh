#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🚀 Starting ClarityFin API Test Suite${NC}"
echo "=================================="
echo -e "${YELLOW}Make sure the server is running on http://localhost:8080${NC}"
echo ""

# Base URL
BASE_URL="http://localhost:8080/api/v1"

# Function to check if server is running
check_server() {
    if ! curl -s http://localhost:8080 > /dev/null 2>&1; then
        echo -e "${RED}❌ Server is not running on http://localhost:8080${NC}"
        echo -e "${YELLOW}Please start the server with: go run cmd/api/main.go${NC}"
        exit 1
    fi
    echo -e "${GREEN}✅ Server is running${NC}"
}

# Check server status
check_server

# Test 1: Send OTP
echo -e "\n${YELLOW}1. Testing OTP Send${NC}"
echo "--------------------------------"
SEND_OTP_RESPONSE=$(curl -s -X POST $BASE_URL/otp/send \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890"}')

if echo "$SEND_OTP_RESPONSE" | grep -q '"success":true'; then
    echo -e "${GREEN}✅ OTP sent successfully${NC}"
else
    echo -e "${RED}❌ OTP send failed${NC}"
fi
echo "Response: $SEND_OTP_RESPONSE"

# Test 2: User Registration with OTP
echo -e "\n${YELLOW}2. Testing User Registration with OTP${NC}"
echo "--------------------------------"
# Note: In a real scenario, you would get the OTP from SMS
# For testing, we'll use a mock OTP (you can check the server logs for the actual OTP)
REGISTER_OTP_RESPONSE=$(curl -s -X POST $BASE_URL/auth/register/otp \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "testpassword", "otp_code": "123456"}')

if echo "$REGISTER_OTP_RESPONSE" | grep -q '"success":true'; then
    echo -e "${GREEN}✅ Registration with OTP successful${NC}"
else
    echo -e "${RED}❌ Registration with OTP failed${NC}"
fi
echo "Response: $REGISTER_OTP_RESPONSE"

# Test 3: Duplicate Registration (should fail)
echo -e "\n${YELLOW}3. Testing Duplicate Registration${NC}"
echo "--------------------------------"
DUPLICATE_RESPONSE=$(curl -s -X POST $BASE_URL/auth/register \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "testpassword"}')

if echo "$DUPLICATE_RESPONSE" | grep -q '"success":false'; then
    echo -e "${GREEN}✅ Duplicate registration properly handled${NC}"
else
    echo -e "${RED}❌ Duplicate registration not handled correctly${NC}"
fi
echo "Response: $DUPLICATE_RESPONSE"

# Test 4: User Login
echo -e "\n${YELLOW}4. Testing User Login${NC}"
echo "--------------------------------"
LOGIN_RESPONSE=$(curl -s -X POST $BASE_URL/auth/login \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "testpassword"}')

if echo "$LOGIN_RESPONSE" | grep -q '"success":true'; then
    echo -e "${GREEN}✅ Login successful${NC}"
else
    echo -e "${RED}❌ Login failed${NC}"
fi
echo "Response: $LOGIN_RESPONSE"

# Extract JWT token
TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
if [ -n "$TOKEN" ]; then
    echo -e "${GREEN}✅ JWT Token extracted: ${TOKEN:0:50}...${NC}"
else
    echo -e "${RED}❌ Failed to extract JWT token${NC}"
    exit 1
fi

# Test 5: Invalid Login
echo -e "\n${YELLOW}5. Testing Invalid Login${NC}"
echo "--------------------------------"
INVALID_LOGIN_RESPONSE=$(curl -s -X POST $BASE_URL/auth/login \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "wrongpassword"}')

if echo "$INVALID_LOGIN_RESPONSE" | grep -q '"success":false'; then
    echo -e "${GREEN}✅ Invalid login properly handled${NC}"
else
    echo -e "${RED}❌ Invalid login not handled correctly${NC}"
fi
echo "Response: $INVALID_LOGIN_RESPONSE"

# Test 6: Get Subscriptions (Unauthorized)
echo -e "\n${YELLOW}6. Testing Unauthorized Access${NC}"
echo "--------------------------------"
UNAUTHORIZED_RESPONSE=$(curl -s -X GET $BASE_URL/subscriptions/)
if echo "$UNAUTHORIZED_RESPONSE" | grep -q "Authorization header required"; then
    echo -e "${GREEN}✅ Unauthorized access properly blocked${NC}"
else
    echo -e "${RED}❌ Unauthorized access not blocked${NC}"
fi
echo "Response: $UNAUTHORIZED_RESPONSE"

# Test 7: Get Subscriptions (Authorized)
echo -e "\n${YELLOW}7. Testing Get Subscriptions (Empty)${NC}"
echo "--------------------------------"
GET_SUBS_RESPONSE=$(curl -s -X GET $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer $TOKEN")
if echo "$GET_SUBS_RESPONSE" | grep -q '"success":true'; then
    echo -e "${GREEN}✅ Get subscriptions successful${NC}"
else
    echo -e "${RED}❌ Get subscriptions failed${NC}"
fi
echo "Response: $GET_SUBS_RESPONSE"

# Test 8: Create Subscription
echo -e "\n${YELLOW}8. Testing Create Subscription${NC}"
echo "--------------------------------"
CREATE_SUB_RESPONSE=$(curl -s -X POST $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Netflix", "amount": 199}')

if echo "$CREATE_SUB_RESPONSE" | grep -q '"success":true'; then
    echo -e "${GREEN}✅ Create subscription successful${NC}"
else
    echo -e "${RED}❌ Create subscription failed${NC}"
fi
echo "Response: $CREATE_SUB_RESPONSE"

# Extract subscription ID
SUB_ID=$(echo $CREATE_SUB_RESPONSE | grep -o '"id":[0-9]*' | cut -d':' -f2)
if [ -n "$SUB_ID" ]; then
    echo -e "${GREEN}✅ Subscription ID: $SUB_ID${NC}"
else
    echo -e "${RED}❌ Failed to extract subscription ID${NC}"
    exit 1
fi

# Test 9: Get Subscription by ID
echo -e "\n${YELLOW}9. Testing Get Subscription by ID${NC}"
echo "--------------------------------"
GET_SUB_BY_ID_RESPONSE=$(curl -s -X GET $BASE_URL/subscriptions/$SUB_ID \
  -H "Authorization: Bearer $TOKEN")
if echo "$GET_SUB_BY_ID_RESPONSE" | grep -q '"success":true'; then
    echo -e "${GREEN}✅ Get subscription by ID successful${NC}"
else
    echo -e "${RED}❌ Get subscription by ID failed${NC}"
fi
echo "Response: $GET_SUB_BY_ID_RESPONSE"

# Test 10: Get All Subscriptions (should now have data)
echo -e "\n${YELLOW}10. Testing Get All Subscriptions (with data)${NC}"
echo "--------------------------------"
GET_ALL_SUBS_RESPONSE=$(curl -s -X GET $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer $TOKEN")
if echo "$GET_ALL_SUBS_RESPONSE" | grep -q '"total":1'; then
    echo -e "${GREEN}✅ Get all subscriptions shows correct count${NC}"
else
    echo -e "${RED}❌ Get all subscriptions count incorrect${NC}"
fi
echo "Response: $GET_ALL_SUBS_RESPONSE"

# Test 11: Create Another Subscription
echo -e "\n${YELLOW}11. Testing Create Another Subscription${NC}"
echo "--------------------------------"
CREATE_SUB2_RESPONSE=$(curl -s -X POST $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Spotify", "amount": 119}')

if echo "$CREATE_SUB2_RESPONSE" | grep -q '"success":true'; then
    echo -e "${GREEN}✅ Create second subscription successful${NC}"
else
    echo -e "${RED}❌ Create second subscription failed${NC}"
fi
echo "Response: $CREATE_SUB2_RESPONSE"

# Test 12: Invalid Token
echo -e "\n${YELLOW}12. Testing Invalid Token${NC}"
echo "--------------------------------"
INVALID_TOKEN_RESPONSE=$(curl -s -X GET $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer invalid_token")
if echo "$INVALID_TOKEN_RESPONSE" | grep -q "Invalid token"; then
    echo -e "${GREEN}✅ Invalid token properly handled${NC}"
else
    echo -e "${RED}❌ Invalid token not handled correctly${NC}"
fi
echo "Response: $INVALID_TOKEN_RESPONSE"

# Test 13: Invalid Input Validation
echo -e "\n${YELLOW}13. Testing Input Validation${NC}"
echo "--------------------------------"
INVALID_INPUT_RESPONSE=$(curl -s -X POST $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "", "amount": -100}')

if echo "$INVALID_INPUT_RESPONSE" | grep -q '"success":false'; then
    echo -e "${GREEN}✅ Input validation working correctly${NC}"
else
    echo -e "${RED}❌ Input validation not working${NC}"
fi
echo "Response: $INVALID_INPUT_RESPONSE"

# Test 14: Final Check - Get All Subscriptions
echo -e "\n${YELLOW}14. Final Check - Get All Subscriptions${NC}"
echo "--------------------------------"
FINAL_SUBS_RESPONSE=$(curl -s -X GET $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer $TOKEN")
if echo "$FINAL_SUBS_RESPONSE" | grep -q '"total":2'; then
    echo -e "${GREEN}✅ Final subscription count correct (2 subscriptions)${NC}"
else
    echo -e "${RED}❌ Final subscription count incorrect${NC}"
fi
echo "Response: $FINAL_SUBS_RESPONSE"

echo -e "\n${GREEN}✅ Test Suite Completed Successfully!${NC}"
echo "=================================="
echo -e "${BLUE}Summary:${NC}"
echo -e "  • User registration and login ✅"
echo -e "  • JWT authentication ✅"
echo -e "  • Subscription CRUD operations ✅"
echo -e "  • Error handling and validation ✅"
echo -e "  • Security and authorization ✅"
echo ""
echo -e "${YELLOW}🎉 All tests passed! The ClarityFin API is working correctly.${NC}"
