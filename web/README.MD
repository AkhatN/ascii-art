# ASCII-ART-WEB

## Description

Ascii-art-web is a web server, in which it will be possible to use a web GUI (graphical user interface) version of ascii-art. This web application can be run in a docker container, which you need to run after creating a docker image with a dockerfile.

In the webpage you can use different banners. The web application can also export the output to a file.

### HTTP endpoints

 1. GET /: Sends HTML response, the main page.

1. POST /ascii-art: that sends data to Go server (text and a banner)

The main page has:
* text input
* select object to switch between banners
* button, which sends a POST request to '/ascii-art' and outputs the result on the page.



## How to run
Run the following commands:
1. For running the server: **go run server.go**
1. For building an image: **docker image build -t ascii-art-web .**
2. For showing images: **docker images**
3. For running a container: **docker container run -d --name=ascii-art -p 8070:8070 ascii-art-web**
4. For showing files inside the running container: **docker exec -it ascii-art ls -l**

Then:
1. Open **localhost:8070** in any browser.
2. Choose any banner you want.
3. Type some text and click "Show" for printing the text in ASCII-art or click "Download" for saving the art as a file in .txt format 


