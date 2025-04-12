# Build stage
FROM node:20-alpine AS builder

WORKDIR /app
COPY . .

RUN npm install
RUN npm run build

# Production stage
FROM node:20-alpine AS runner

WORKDIR /app

# 必要なファイルだけコピー
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json

RUN npm install --production

EXPOSE 3000

CMD ["npm", "start"]
