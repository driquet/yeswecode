#!/bin/bash

convert waldo.png waldo.jpeg
exiftool -GPSLatitude=33.12902467508919 -GPSLatitudeRef=N -GPSLongitude=-107.250729874426 -GPSLongitudeRef=W waldo.jpeg
