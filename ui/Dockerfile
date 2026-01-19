FROM node:25-alpine AS builder
WORKDIR /app

COPY package*.json ./
RUN npm i
COPY . .
RUN VITE_BACKEND_URL="__BACKEND_URL__" \
    npm run build

FROM nginx:stable-alpine AS production

COPY --from=builder /app/dist /usr/share/nginx/html
COPY entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
ENTRYPOINT ["/docker-entrypoint.sh"]