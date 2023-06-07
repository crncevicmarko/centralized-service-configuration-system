# airs-projekat

Alati i Razvoj Softvera Projekat 2023

# Docker Komande

Latest(Non Multistage):
docker build --tag airs-projekat:latest .
docker run -p 8000:8000 airs-projekat:latest

Multistage:
docker build -t airs-projekat:multistage -f Dockerfile.multistage .
docker run -p 8000:8000 airs-projekat:multistage
