
set -o errexit
set -o nounset

export PARSER_CONTAINER=go-parser

if [ "$(docker ps -a | grep ${PARSER_CONTAINER})" ]
then	  
    echo "killing existing image..."
    docker kill "${PARSER_CONTAINER}"
		docker rm "${PARSER_CONTAINER}"
fi
 
 echo "Building container..." 
 docker build --tag "${PARSER_CONTAINER}" .	
 echo "building finished."
