
# **URL Shortener Service**

This is a simple URL Shortener implementation written in **Go (Golang)**. The application allows users to shorten long URLs and redirect them using short URLs. It includes features like custom aliasing and expiration of links.

---

## **Contents**

1. [Features](#features)
2. [Getting Started](#getting-started)
3. [Testing Using `curl`](#testing-using-curl)
4. [Building the Application](#building-the-application)
5. [Running Unit Tests](#running-unit-tests)

---

## **Features**

- Generate short URLs for long URLs.
- URL expiration policies.
- In-memory data store.

---

## **Getting Started**

### **Prerequisites**
- **Go** (1.17 or later) installed on your system.
- Basic familiarity with using a terminal or command line.

### **Run Locally**

1. Clone the repository to your local system:
   ```bash
   git clone https://github.com/your-repo/url-shortener.git
   cd url-shortener
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. By default, the application will be running on **http://localhost:8080**.

---

## **Testing Using `curl`**

You can use the following `curl` commands to test the endpoints of the application.

### **1. Shorten a URL**

#### **Anonymous User**
To shorten a URL as an anonymous user:
```bash
curl -X POST http://localhost:8080/shorten \
-H "Content-Type: application/json" \
-d '{
  "long_url": "https://example.com"
}'
```

**Expected Response:**
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

#### **Registered User (With a Custom Alias)**
To shorten a URL with a custom alias (for registered users):
```bash
curl -X POST http://localhost:8080/shorten \
-H "Content-Type: application/json" \
-H "Authorization: registered_user_token" \
-d '{
  "long_url": "https://example.com",
  "custom_alias": "customAlias123"
}'
```

**Expected Response (if alias is available):**
```json
{
  "short_url": "http://localhost:8080/customAlias123"
}
```

**Error Response (if alias is already taken):**
```json
{
  "error": "Alias already taken. Please try another."
}
```

---

### **2. Access a Shortened URL**

Suppose you received a short URL like `http://localhost:8080/abc123` in the previous step. You can test the redirect functionality with:

```bash
curl -v http://localhost:8080/abc123
```

**Expected Output (Redirect):**
