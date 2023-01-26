cd /mnt/cache/SSD/overLinks/overLinks
docker build -t overlinks-be -f Dockerfile.backend .
docker build -t overlinks-fe -f Dockerfile.frontend .
docker build -t overlinks-tg -f Dockerfile.telegram .