# 2-10

- Moved env variables from `Dockerfiles` to `docker-compose.yml`
- Removed ports from backend `REQUEST_ORIGIN` & frontend `REACT_APP_BACKEND_URL` and let Nginx handle proxying
