# lol-champ-groups
El algoritmo utlizado es k-means clustering.

El dataset fue recopilado a través de webscraping, con la biblioteca Selenium en Rust, de la página https://lolalytics.com/, en la cual se analizan estadísticas del final de millones de partidas del juego League of Legends.

El dataset contiene 148 personajes donde cada uno tiene 5 roles a los que pueden ser asignados en cada partida. Dentro de estos 5 roles hay 15 columnas que representan el promedio de estadísticas al final de cada partida en la versión 10.10 del juego con respecto al rendimiento del personaje.

En la interfaz el usuario puede:
- Seleccionar la cantidad de clusters a generar
- Seleccionar el método de cálculo de la distancia entre datapoints
- El rol a analizar de todos los personajes
- Las estadísticas (columnas) a analizar

Pasos para ejecutar
1. cd rest_api/
2. go run main.go
3. cd website/
4. npm install
5. npm start
