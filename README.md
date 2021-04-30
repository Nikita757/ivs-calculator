# IVS Calculator
Graphical calculator in Golang. BUT FIT project to practice Git and teamwork

## Environment

Ubuntu 64bit

## Building

First install the dependencies: 

```
sudo apt-get install golang-go go-bindata libgtk-3-dev libcairo2-dev libglib2.0-dev libwebkit2gtk-4.0-dev pandoc
# for building PDF manual:
sudo apt-get install texlive-latex-base texlive-fonts-recommended texlive-fonts-extra texlive-latex-extra texlive-xetex
```

To build the app and profiler:

```
cd src
make
```

To build .deb package:

```
make debian
```

To build the user manual:

```
make manual
```

To install manually after building:

```
sudo cp cmd/gui/gui /usr/bin/ivs-calculator
sudo cp assets/ivs-calculator.desktop /usr/share/applications/ivs-calculator.desktop
sudo cp assets/ivs_calculator.svg -p /usr/share/icons/hicolor/scalable/apps/ivs-calculator.svg
```

To uninstall manually:

```
sudo rm -f /usr/bin/ivs-calculator
sudo rm -f /usr/share/applications/ivs-calculator.desktop
sudo rm -f /usr/share/icons/hicolor/scalable/apps/ivs-calculator.svg
```

To build the source code docs, refer to [src/docs/README.md](src/docs/README.md).

## Authors

- Tomáš Martykán (xmarty07)

- Filip Štolfa (xstolf00)

- Andrei Meleca (xmelec02)

- Nichita Gutu (xgutun00)

## License

The program is provided under the [GNU General Public License v3.0](LICENSE)

© Copyright 2021 Tomáš Martykán, Filip Štolfa, Andrei Meleca, Nichita Gutu
