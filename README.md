# Images service

Simple images service for uploading and retrieving iamges

## Start service

```bash
docker-compose up
```

## Endpoints
* Upload image: POST /image
* List all images: GET /image
* Delete image: DELETE /image/{image-name}
* Download image: GET /image/{image-name}
