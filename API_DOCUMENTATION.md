# API Documentation

Base URL: `http://localhost:8080/api`

All requests require header: `X-API-Key: <your key>`

Missing/invalid key → `401`

```json
{ "error": "missing or invalid API key" }
```

## POST /customers

```json
{
  "nationality_id": 1,
  "cst_name": "Vincent Xu",
  "cst_dob": "1995-05-20T00:00:00Z",
  "cst_phoneNum": "081234567890",
  "cst_email": "vincent@example.com",
  "family": [
    { "fl_relation": "Father", "fl_name": "John Xu", "fl_dob": "1965-01-01" }
  ]
}
```

→ `201`

```json
{ "cst_id": 1 }
```

→ `400`

```json
{ "error": "nationality_id, cst_name, and cst_email are required" }
```

---

## GET /customers

→ `200`

```json
[
  {
    "cst_id": 1,
    "nationality_id": 1,
    "cst_name": "Vincent Xu",
    "cst_dob": "1995-05-20T00:00:00Z",
    "cst_phoneNum": "081234567890",
    "cst_email": "vincent@example.com",
    "nationality": {
      "nationality_id": 1,
      "nationality_name": "Indonesia",
      "nationality_code": "ID"
    },
    "family": [
      {
        "fl_id": 1,
        "cst_id": 1,
        "fl_relation": "Father",
        "fl_name": "John Xu",
        "fl_dob": "1965-01-01"
      }
    ]
  }
]
```

---

## GET /customers/{id}

→ `200`

```json
{
  "cst_id": 1,
  "nationality_id": 1,
  "cst_name": "Vincent Xu",
  "cst_dob": "1995-05-20T00:00:00Z",
  "cst_phoneNum": "081234567890",
  "cst_email": "vincent@example.com",
  "nationality": {
    "nationality_id": 1,
    "nationality_name": "Indonesia",
    "nationality_code": "ID"
  },
  "family": [
    {
      "fl_id": 1,
      "cst_id": 1,
      "fl_relation": "Father",
      "fl_name": "John Xu",
      "fl_dob": "1965-01-01"
    }
  ]
}
```

→ `404`

```json
{ "error": "customer not found" }
```

---

## PUT /customers/{id}

Replaces the entire family list with what's sent.

```json
{
  "nationality_id": 1,
  "cst_name": "Vincent Xu Updated",
  "cst_dob": "1995-05-20T00:00:00Z",
  "cst_phoneNum": "081298765432",
  "cst_email": "vincent.updated@example.com",
  "family": [
    { "fl_relation": "Father", "fl_name": "John Xu", "fl_dob": "1965-01-01" }
  ]
}
```

→ `200`

```json
{ "message": "updated successfully" }
```

→ `400`

```json
{
  "error": "failed to update customer — check that nationality_id exists and all fields are valid"
}
```

---

## DELETE /customers/{id}

Family members cascade-delete automatically.

→ `200`

```json
{ "message": "deleted successfully" }
```

---

## GET /nationalities

→ `200`

```json
[
  {
    "nationality_id": 1,
    "nationality_name": "Indonesia",
    "nationality_code": "ID"
  },
  {
    "nationality_id": 2,
    "nationality_name": "Malaysia",
    "nationality_code": "MY"
  }
]
```

---

## Notes

- `cst_dob`: ISO 8601 datetime (`YYYY-MM-DDT00:00:00Z`)
- `fl_dob`: plain date string (`YYYY-MM-DD`)
- `nationality_id` must reference an existing row
- Create/Update run in a single DB transaction — customer + family succeed or fail together
- Rate limit: 100 requests/minute per IP → `429` if exceeded
- CORS: only requests from `ALLOWED_ORIGIN` (env var) are accepted
