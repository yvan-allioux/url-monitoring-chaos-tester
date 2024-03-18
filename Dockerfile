FROM php:7.4-apache

COPY ./src/ /var/www/html/

WORKDIR /var/www/html

EXPOSE 8080

CMD ["apache2-foreground"] 

#docker build -t url-monitoring-chaos-tester .
#docker save -o url-monitoring-chaos-tester.tar url-monitoring-chaos-tester
#scp url-monitoring-chaos-tester.tar debian@141.94.206.18:/home/debian

#docker load -i url-monitoring-chaos-tester.tar
#sudo rm url-monitoring-chaos-tester.tar

#docker run -p 8080:80
#docker run --restart=always -p 8080:80

#docker run -p 8080:80 url-monitoring-chaos-tester
