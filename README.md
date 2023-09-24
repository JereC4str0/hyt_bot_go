
# Simple Go Bot

Este es un bot de Telegram diseñado para monitorear la temperatura y humedad de una incubadora.


## Acerca del Proyecto

Este bot utiliza la API de Telegram para proporcionar información en tiempo real sobre la temperatura y humedad de una incubadora. Utiliza una API web para obtener datos actualizados y luego responde a comandos específicos para mostrar esta información.

## Funcionalidades

- **/temperatura**: Devuelve la temperatura actual en la incubadora en grados Celsius.
- **/humedad**: Devuelve la humedad actual en la incubadora en porcentaje.

## Instalación

1. Clona este repositorio: `git clone https://github.com/JereC4str0/hyt_bot_go.git`.
2. Accede al directorio del proyecto: `cd cmd/bot`.
3. Crea un archivo `config.json` con la siguiente estructura y agrega tu token de bot de Telegram:
    ```json
    {
        "telegram_bot_token": "TU_TOKEN_AQUI"
    }
    ```
4. Instala las dependencias: `go mod tidy`.
5. Ejecuta el bot: `go run main.go`.

## Uso

Una vez que el bot esté en funcionamiento, simplemente interactúa con él en Telegram usando los comandos `/temperatura` y `/humedad` para obtener la información de monitoreo.



## Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo [LICENSE](LICENSE) para obtener más detalles.
