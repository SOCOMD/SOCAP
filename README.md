# SOCAP
OCAP rewrite for our A3 milsim group [SOCOMD](http://www.socomd.com/)

# Terrain Export
## Prerequisites
* Download https://community.bistudio.com/wiki/BI_Tools_2.5 (Required: Visitor 3)
* Download https://trac.osgeo.org/osgeo4w/ (Required: gdal/gdal2 + dependencies)
* Patch OSGeo4W64. (./Resources/OSGeo4W64)

## .emf Export
1. Open Arma 3
2. Open Editor Map Selection and highlight a terrain
3. Open 2D Editor [Ctrl - O]
4. Press [Left Shift + Numpad Minus]
5. Blindly type exportnogrid
6. Terrain .emf should be exported to C:/

## .emd > .png
```
BI_Tools_2.5/Visitor 3/EmfToPng.exe [<.emf file>] [scale]
```
1. Target image size is 16384 x 16384
2. An export with the scale of 1 will give the native map size, if Width/Height isn't equal take the smaller one. (Size used in map.json)
3. Calculate scale required [image size / 16384]
4. Export with calculated scale. (Scale used in map.json)
5. Crop the image if dimentions are not equal

## OSGeo4W64 .png > tiles
1. Launch OSGeo4W64.bat
2. gdal2tiles_legacy_no-tms -p raster -z 0-6 [map.png]
3. Exports to a folder with the name of the terrain

## map.json
Add map.json to exported terrain tile folder
```
{
	"name": "Isla Duala 3", //Display Name
	"worldName": "isladuala3", //Same as mission extension
	"worldSize": 10240, //Map size from 1x scale export
	"imageSize": 16384,
	"multiplier": 1.6 //Scale used for 16384 x 16384 export
}
```

## Credits
* [armago](https://github.com/code34/armago_x64)
* [OCAP](https://github.com/ocapmod/OCAP)
* [Zealot111 OCAP](https://github.com/Zealot111/OCAP)