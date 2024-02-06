version: "3"
services:
  {{.moduleProjectName}}:
    container_name: {{.moduleProjectName}}
    build: .
    ports:
      - 8000:8000
    volumes:
      - ./config:/config
      - ./runtime:/runtime
    restart: always
