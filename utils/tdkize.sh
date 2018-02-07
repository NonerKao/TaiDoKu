#!/bin/bash

if [ $# != 1 ]; then
	echo "usage: ./tdkize.sh file"
	exit
fi

sed -e "s/０/0 /g" -i $1
sed -e "s/１/1 /g" -i $1
sed -e "s/２/2 /g" -i $1
sed -e "s/３/3 /g" -i $1
sed -e "s/４/4 /g" -i $1
sed -e "s/５/5 /g" -i $1
sed -e "s/６/6 /g" -i $1
sed -e "s/７/7 /g" -i $1
sed -e "s/８/8 /g" -i $1
sed -e "s/９/9 /g" -i $1
sed -e "s/Ａ/A /g" -i $1
sed -e "s/Ｂ/B /g" -i $1
sed -e "s/Ｃ/C /g" -i $1
sed -e "s/Ｄ/D /g" -i $1
sed -e "s/Ｅ/E /g" -i $1
sed -e "s/Ｆ/F /g" -i $1
sed -e "s/║//g" -i  $1 
sed -e "s/┆//g" -i  $1 
sed -e "s/╌//g" -i  $1 
sed -e "s/┼//g" -i  $1 
sed -e "s/╢//g" -i  $1 
sed -e "s/╟//g" -i  $1
sed -e "s/╚//g" -i $1
sed -e "s/═//g" -i $1
sed -e "s/╧//g" -i $1
sed -e "s/╝//g" -i $1
sed -e "s/╔//g" -i $1 
sed -e "s/╤//g" -i $1
sed -e "s/╗//g" -i $1
sed -e "/^$/d" -i $1
