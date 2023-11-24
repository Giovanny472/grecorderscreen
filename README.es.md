
# **gScreenshot**

| [ру](https://github.com/Giovanny472/gscreenshot) | [es](https://github.com/Giovanny472/gscreenshot/blob/main/README.es.md) |

Es un simple programa para hacer screenshot. Escrito en el lenguaje go.
Este programa esta basado en la biblioteca [screenshot](https://github.com/kbinani/screenshot). Con las bibliotecas  [glfw](https://github.com/go-gl/glfw) и [gl](https://github.com/go-gl/gl) fue hecha la interfaz de usuario, la cual permite elegir una region determinada para el screenshot. Las coordenadas de la region seleccionada se envian como parametros a la biblioteca [screenshot](https://github.com/kbinani/screenshot) para hacer la captura de la region de la pantalla. Despues de la obtener la region de la pantalla se utiliza la biblioteca [clipboard](https://github.com/golang-design/clipboard) para pasar la imagen al portapapeles. De esta forma se puede obtener la imagen del screenshot con una simple combinacion crtl+v.

![gscreenshot](/asset/img/gscreenshot.png)

## Bibliotecas utilizadas
+  [screenshot](https://github.com/kbinani/screenshot)
+  [glfw](https://github.com/go-gl/glfw)
+  [gl](https://github.com/go-gl/gl)
+  [clipboard](https://github.com/golang-design/clipboard) 

https://github.com/Giovanny472/gscreenshot/assets/43882462/cdc30927-4055-4a4d-b065-bf576c4526aa

## Instalacion del programa
Para instalar **gScreenshot** se puede utilizar [gScreenshot_setup.exe](https://github.com/Giovanny472/gscreenshot/releases/tag/v0.1). con la combinacion ctrl+alt+s se ejecuta el programa.
