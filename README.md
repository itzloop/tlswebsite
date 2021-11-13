## Single Page Website Secured With TLS

- [x] Add the Page
- [x] Add the Server
- [x] Add TLS
- [x] Dockerize
- [x] Nginx
- [ ] Docker Compose

## What I've Done, TL;DR
- First we need a webserver which is located at `main.go` serving the contetns of `static` directory which contains `index.html`. This is signle page website and nothing more. The websever's code is also super simple.
- After that we had a website running, it was time to Dockerize the project with image name being `sinashk/tlswebsite`.
- I also tried to use `docker-compose` and move `nginx` and `website` on there but had some issues and failed in doing so.
- The used a domain `shantech.ir` on virtual machine from arvan cloud and got a certificate from let's encrypt and configured nginx.

## Steps in Securing With SSL
Keep in mind that the packages used are for ubuntu server you might have the same packages or not, check the name of these packages in you own linux distribution.
- first updated my DN recored with arvans dns.
- then sshed into the server and did some updates and installed `nginx` and `certbot` and `python3-certbot-nginx`
- then created `www.shantech.ir.conf` inside `/etc/nginx/conf.d/` and inside it was this content:
```
# file /etc/nginx/conf.d/www.shantech.ir.conf
server {
    listen 80 default_server;
    listen [::]:80 default_server;
    root /var/www/html;
    server_name shantech.ir www.shantech.ir;
}
```
- then we reload `nginx`. I had some problems with this operation saying configuration already exists. The reason was that my config was listening to `default_server`, so as 2 other config scratered around. I simply deleted them you can change the `default_server` to something else both work.
```
$ sudo nginx -t && nginx -s reload
```
- then ran the following command with `certbot` to generag certificates and also update `www.shantech.ir.con.conf`
```
sudo certbot --nginx -d shantech.ir -d www.shantech.ir
```
- After this operation certbot asked to update the nginx config so that it would redirect all http traffic to https and i said yes. the final config is included in the root directory of the project with the name `nginx.conf`.
- Now just update the config and your locations. I added `/` to go to `localhost:8080` which im going to run the webserver on that port later. this is the the part i added:
```
# file /etc/nginx/conf.d/www.shantech.ir.conf
erver {
    root /var/www/html;
    server_name shantech.ir www.shantech.ir;

    listen [::]:443 ssl ipv6only=on; # managed by Certbot
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/shantech.ir/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/shantech.ir/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

################ I ADDED ##################
    location ~ ^/ {
       proxy_pass  http://localhost:8080;
    }
###########################################
}
server {
    if ($host = www.shantech.ir) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


    if ($host = shantech.ir) {
        return 301 https://$host$request_uri;
    } # managed by Certbot


    listen 80 default_server;
    listen [::]:80 default_server;
    server_name shantech.ir www.shantech.ir;
    return 404; # managed by Certbot
}

```
- Then i pulled the image with `docker pull sinashk/tlswebsite`. You might need to login to your account for this one with `docker login` and you might also need a vpn ;)
- After the image was pulled, we can simply run it with:
```
$ docker run -p 8080:8080 sinashk/tlswebsite
```
- Well there is one more thing we need block other traffics otherwise some can ignore tls by going straigh to port 8080 so we need to activate our firewall. I will use `ufw`.
```
# IMPRTANT otherwise you will lose your session and probably won't be able to connect to your sever any more
$ sudo ufw allow ssh 
$ sudo ufw allow https
$ sudo ufw allow http
# now enable the firewall
$ sudo ufw enable
# check you firewall status
$ sudo ufw status
```
![proof](drive.google.com/uc?id=1KkMS5MPUz4rueYwyfg3Dzc9rCOozqJbo)
- Enjoy your secure website :). It's still sad that i couldn't make docker compose work :(.
## Notes:
Question 1 to 3 are located in `questions` directory.


