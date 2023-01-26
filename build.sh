cd /mnt/cache/SSD/overLinks/overLinks
docker build -t overlinks-be Dockerfile.backend .
docker build -t overlinks-fe Dockerfile.frontend .
docker build -t overlinks-tg Dockerfile.telegram .