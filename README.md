# zcrapper

Brograma para descargar archivos de la página.

### Uso
1) Descargar el binario que corresponde a tu sistema operativo en la página de [releases](https://github.com/anongolico/zcrapper/releases) de github.
2) La primera vez que el programa corra, pedirá ingresar la cookie de acceso. Esto pasa porque Mordekai tiene el muro activado y si no está esa cookie, la página devuelve el screen de Chocamo.\

![el op](https://github.com/anongolico/zcrapper/blob/main/img/opegolico.gif?raw=true "OP")
***
##### Guía para agregar la cookie
- en tu navegador web, en la página de rouzed apretar **F12** para ingresar a las herramientas de desarrollo.
- buscar la pestaña '*Storage*' o '*Almacenamiento*'
- buscar la cookie de nombre *.AspNetCore.Identity.Application*
- en la columna '*value*' o '*valor*', dar doble click y luego Ctrl+C para pegar el valor
![](https://raw.githubusercontent.com/anongolico/zcrapper/main/img/2.png "instrucciones")
***
3) Pega el id del rouz. Si la URL es: `https://rouzed.one/Hilo/07TOQ14MNBNI9Z1U1QHP`
entonces el id es `07TOQ14MNBNI9Z1U1QHP`
4) lito

#### Aclaración para usuarios de windowns

si al hacer doble click sobre el ejecutable, no funciona, lo único que hay que hacer es correrlo desde la terminal directamente. Para hacerlo, en el menú de inicio buscar *cmd*, luego arrastrar la carpeta donde está el ejecutable (o navegar hasta ella con el comando *chdir*). Una vez ahí, escribir zcrapper.exe y lito.

El programa usa este otro [repo](https://github.com/anongolico/base). Si deseas modificar algo, es mejor clonar ambos. También con ese es posible modificar algunos parámetros para hacer un scrapper para otros clones.