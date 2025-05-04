# Build stage
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

# Production stage - just the built files
FROM node:20-alpine
WORKDIR /app
COPY --from=builder /app/build ./build
RUN npm install -g serve


RUN echo '{"rewrites": [{"source": "**", "destination": "/index.html"}]}' > serve.json

CMD ["serve", "build", "-p", "3000", "-c", "../serve.json"]
