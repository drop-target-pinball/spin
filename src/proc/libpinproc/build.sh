#!/bin/bash -e

LIBUSB_VERSION=1.0.27
LIBFTDI_VERSION=1.5
LIBPINPROC_VERSION=2.0

BUILDROOT=$(pwd)/root
PREFIX=/usr/local

mkdir -p sources
wget \
    --input-file wget.txt \
    --no-clobber \
    --no-directories \
    --directory-prefix sources

rm -rf build $BUILDROOT
mkdir -p build $BUILDROOT

tar xf sources/libusb-${LIBUSB_VERSION}.tar.bz2 -C build
(
    cd build/libusb-${LIBUSB_VERSION}
    ./configure --prefix=${PREFIX}
    make
    sudo make install
)

tar xf sources/libftdi1-${LIBFTDI_VERSION}.tar.bz2 -C build
(
    cd build/libftdi1-${LIBFTDI_VERSION}
    mkdir -p _build
    cd  _build
    cmake \
        -DCMAKE_INSTALL_PREFIX:PATH=${PREFIX} \
        -DFTDI_EEPROM=OFF \
        ..
    make
    sudo make install
)

(
    cd build
    git clone https://github.com/preble/libpinproc.git
    cd libpinproc
    mkdir -p _build
    cd _build
    cmake \
        -DCMAKE_INSTALL_PREFIX:PATH=${PREFIX} \
        ..
    make
    sudo make install
)