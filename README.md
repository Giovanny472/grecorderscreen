
# **gScreenshot**

Простая программа для скриншотов. Написано на языке go.  
Эта программа основана на библиотеке  [screenshot](https://github.com/kbinani/screenshot). C помощью библиотек: [glfw](https://github.com/go-gl/glfw) и [gl](https://github.com/go-gl/glfw) был реализован интерфейс пользователя, который позволяет выбрать область(координаты) на экране.  Эти координаты передаются в качестве параметров в библиотеку [screenshot](https://github.com/kbinani/screenshot), чтобы сделать снимок экрана. После получения изображения выбранного региона экрана, используется библиотека [clipboard](https://github.com/golang-design/clipboard) для передачи изображения в буфер обмена. Таким образом можно получит изображение простым нажатием ctrl+v.

![gscreenshot](/asset/img/gscreenshot.png)



https://github.com/Giovanny472/gscreenshot/assets/43882462/cdc30927-4055-4a4d-b065-bf576c4526aa

