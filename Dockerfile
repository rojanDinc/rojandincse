FROM nginx:1.27.1-alpine
COPY index.html /usr/share/nginx/html
COPY blog.html /usr/share/nginx/html

