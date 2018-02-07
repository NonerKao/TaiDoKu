#!/bin/bash

if [ $# != 2 ]; then
	echo "usage: ./gif.sh input output"
	echo "you have to install ttyrec and ttygif first"
	exit
fi

echo "#!/bin/bash" > $2
echo "clear" >> $2
cat $1 | sed -e "s/╔/sleep 0.3 \&\& clear \&\& echo \'╔/" -e "s/╝/╝\'/" >> $2

chmod 775 $2

# ttyrec -e $2
# ttygif ttyrecord
