# zcrapper

Brograma para descargar archivos de la página.

### Uso
1) Descargar el binario que corresponde a tu sistema operativo en la página de [releases](https://github.com/anongolico/zcrapper/releases) de github.
2) En Windows, dar doble click sobre el binario (si así no funciona, ver las instrucciones más abajo). En Linux, en una terminal dentro del directorio que lo contiene, correr ./zcrapper (en caso de que no esté marcado como ejecutable, mandar un chmod +x zcrapper)
3) Pega el id del rouz. Si la URL es: `https://rouzed.one/Hilo/07TOQ14MNBNI9Z1U1QHP`
entonces el id es `07TOQ14MNBNI9Z1U1QHP`
4) lito

![el op](https://github.com/anongolico/zcrapper/blob/main/img/opegolico.gif?raw=true "OP")

***

#### Aclaración para usuarios de windowns

si al hacer doble click sobre el ejecutable, no funciona, lo único que hay que hacer es correrlo desde la terminal directamente. Para hacerlo, en el menú de inicio buscar *cmd*, luego arrastrar la carpeta donde está el ejecutable (o navegar hasta ella con el comando *chdir*). Una vez ahí, escribir zcrapper.exe y lito.


La nueva red del odio (boxed.fun) permite descargar archivos incluso sin credenciales. En caso que, como el muro que ponía Mordrake, la página no permita ver nada, solamente hay que seguir las instrucciones de abajo.
##### Guía para agregar la cookie
- en tu navegador web, en la página de rouzed apretar **F12** para ingresar a las herramientas de desarrollo.
- buscar la pestaña '*Storage*' o '*Almacenamiento*' ('*Aplicación*' si es un navegador basado en Chromium, como Brave, Opera, etc).
- buscar la cookie de nombre *.AspNetCore.Identity.Application*
- en la columna '*value*' o '*valor*', dar doble click y luego Ctrl+C para pegar el valor
![](https://raw.githubusercontent.com/anongolico/zcrapper/main/img/2.png "instrucciones")
- en el archivo **config.yaml** hay un valor que dice CookieValue, dentro de esas comillas pegar el valor.
