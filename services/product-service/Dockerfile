FROM golang:1.25-alpine

WORKDIR /app

# ✅ Step 1: Copy only go.mod & go.sum
COPY go.mod go.sum ./

# ✅ Step 2: Download dependencies (cached layer)
RUN go mod download

# ✅ Step 3: Copy rest of the code
COPY . .

# ✅ Step 4: Run app
CMD ["go", "run", "main.go"]